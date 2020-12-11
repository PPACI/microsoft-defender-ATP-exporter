package exporter

import (
	"context"
	"github.com/PPACI/microsoft-defender-ATP-exporter/pkg/api/vulnerabilities"
	"github.com/PPACI/microsoft-defender-ATP-exporter/pkg/azureauth"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"time"
)

var (
	vulnerabilityGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "defender_atp",
		Name:      "vulnerabilities",
		Help:      "Number of vulnerability found on machines",
	}, []string{"machineId", "severity"})
	machineExposureGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "defender_atp",
		Name:      "machine_exposure",
		Help:      "Number of machine exposed",
	}, []string{"product_name", "severity"})
	exposureScoreGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "defender_atp",
		Name:      "exposure_score",
		Help:      "Exposure score",
	})
)

func init() {
	prometheus.MustRegister(vulnerabilityGauge)
	prometheus.MustRegister(machineExposureGauge)
	prometheus.MustRegister(exposureScoreGauge)
}

func refreshData(authClient *azureauth.AuthClient) {
	refreshVulnerabilities(authClient)
	refreshExposureScore(authClient)
	log.Println("Refreshed data from API")
}

func refreshVulnerabilities(authClient *azureauth.AuthClient) {
	machineVulnerabilitiesData, err := vulnerabilities.GetMachineVulnerabilities(authClient)
	if err != nil {
		log.Println("Error while fetching machine vulnerabilities")
		log.Println(err)
	}
	machineVulnerabilities := make(map[string]map[string]int)
	for _, vuln := range machineVulnerabilitiesData {
		if machineVulnerabilities[vuln.MachineId] == nil {
			machineVulnerabilities[vuln.MachineId] = make(map[string]int)
		}
		machineVulnerabilities[vuln.MachineId][vuln.Severity]++
	}
	for machineId, v := range machineVulnerabilities {
		for severity, count := range v {
			gauge, err := vulnerabilityGauge.GetMetricWithLabelValues(machineId, severity)
			if err != nil {
				log.Fatal(err)
			}
			gauge.Set(float64(count))
		}
	}

	// Software > severity > machineId
	// the machineID map is then used to count the number of machine
	// Go doesn't have a set objet
	exposure := make(map[string]map[string]map[string]bool)
	for _, vuln := range machineVulnerabilitiesData {
		if exposure[vuln.ProductName] == nil {
			exposure[vuln.ProductName] = map[string]map[string]bool{vuln.Severity: {}}
		}
		if exposure[vuln.ProductName][vuln.Severity] == nil {
			exposure[vuln.ProductName][vuln.Severity] = map[string]bool{}
		}
		exposure[vuln.ProductName][vuln.Severity][vuln.MachineId] = true
	}
	for productName, v := range exposure {
		for severity, machineIds := range v {
			gauge, err := machineExposureGauge.GetMetricWithLabelValues(productName, severity)
			if err != nil {
				log.Fatal(err)
			}
			gauge.Set(float64(len(machineIds)))
		}
	}
}

func refreshExposureScore(authClient *azureauth.AuthClient) {
	exposureScoreData, err := vulnerabilities.GetExposureScore(authClient)
	if err != nil {
		log.Println("Error while fetching machine vulnerabilities")
		log.Println(err)
	}
	exposureScoreGauge.Set(exposureScoreData)
}

func StartExporter(authClient *azureauth.AuthClient, ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				refreshData(authClient)
			case <-ctx.Done():
				log.Println(ctx.Err())
				return
			}
		}
	}()

	// To start with fresh data
	go refreshData(authClient)
}
