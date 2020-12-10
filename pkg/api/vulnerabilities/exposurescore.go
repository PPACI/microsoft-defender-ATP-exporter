package vulnerabilities

import (
	"encoding/json"
	"fmt"
	"github.com/PPACI/microsoft-defender-ATP-exporter/pkg/azureauth"
	"io/ioutil"
	"net/http"
)

type exposureScoreApiAnswer struct {
	Context string `json:"@odata.context"`
	Score   float64
}

func GetExposureScore(authClient *azureauth.AuthClient) (float64, error) {
	accessToken, err := authClient.GetToken()
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest(http.MethodGet, "https://api.securitycenter.windows.com/api/exposureScore", nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf(string(body))
	}
	data := exposureScoreApiAnswer{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data.Score, nil
}
