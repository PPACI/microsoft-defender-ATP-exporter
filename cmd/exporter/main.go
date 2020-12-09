package main

import (
	"github.com/PPACI/microsoft-defender-ATP-exporter/pkg/azureauth"
	"log"
	"os"
)

var (
	AzureTenantId     string
	AzureClientId     string
	AzureClientSecret string
)

func init() {
	var ok bool
	AzureTenantId, ok = os.LookupEnv("AZURE_TENANT_ID")
	if !ok {
		log.Fatalln("Missing AZURE_TENANT_ID variable")
	}
	AzureClientId, ok = os.LookupEnv("AZURE_CLIENT_ID")
	if !ok {
		log.Fatalln("Missing AZURE_CLIENT_ID variable")
	}
	AzureClientSecret, ok = os.LookupEnv("AZURE_CLIENT_SECRET")
	if !ok {
		log.Fatalln("Missing AZURE_CLIENT_SECRET variable")
	}
}

func main() {
	log.Println("Init Azure Auth client")
	client := azureauth.NewAuthClient(AzureTenantId, AzureClientId, AzureClientSecret)
	log.Println("Get Token")
	token, err := client.GetToken()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received token : %v\n", token)
}
