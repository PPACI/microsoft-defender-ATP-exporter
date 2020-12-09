package vulnerabilities

import (
	"encoding/json"
	"fmt"
	"github.com/PPACI/microsoft-defender-ATP-exporter/pkg/azureauth"
	"io/ioutil"
	"net/http"
)

type machineVulnerabilitiesApiAnswer struct {
	Context string `json:"@odata.context"`
	Value   []MachineVulnerability
}

type MachineVulnerability struct {
	Id             string
	CveId          string
	MachineId      string
	FixingKbId     string
	ProductName    string
	ProductVendor  string
	ProductVersion string
	Severity       string
}

func GetMachineVulnerabilities(authClient *azureauth.AuthClient) ([]MachineVulnerability, error) {
	accessToken, err := authClient.GetToken()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, "https://api.securitycenter.windows.com/api/vulnerabilities/machinesVulnerabilities", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}
	data := machineVulnerabilitiesApiAnswer{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data.Value, nil
}
