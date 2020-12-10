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
	exposureGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "defender_atp",
		Name:      "exposure",
		Help:      "Exposure of machines",
	}, []string{"product_name", "severity"})
	exposureScoreGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "defender_atp",
		Name:      "exposure_score",
		Help:      "Exposure score",
	})
)

func init() {
	prometheus.MustRegister(vulnerabilityGauge)
	prometheus.MustRegister(exposureGauge)
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

	exposure := make(map[string]map[string]int)
	for _, vuln := range machineVulnerabilitiesData {
		if exposure[vuln.ProductName] == nil {
			exposure[vuln.ProductName] = make(map[string]int)
		}
		exposure[vuln.ProductName][vuln.Severity]++
	}
	for productName, v := range exposure {
		for severity, count := range v {
			gauge, err := exposureGauge.GetMetricWithLabelValues(productName, severity)
			if err != nil {
				log.Fatal(err)
			}
			gauge.Set(float64(count))
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
