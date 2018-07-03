package integration

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
	"github.com/satori/go.uuid"
)

var (
	octopusURL    string
	octopusAPIKey string
	client        *octopusdeploy.Client
)

func initTest() *octopusdeploy.Client {
	octopusURL = os.Getenv("OCTOPUS_URL")
	octopusAPIKey = os.Getenv("OCTOPUS_APIKEY")

	if octopusURL == "" || octopusAPIKey == "" {
		log.Fatal("Please make sure to set the env variables 'OCTOPUS_URL' and 'OCTOPUS_APIKEY' before running this test")
	}

	httpClient := http.Client{}
	client := octopusdeploy.NewClient(&httpClient, octopusURL, octopusAPIKey)

	return client
}

func getRandomName() string {
	fullName := fmt.Sprintf("go-octopusdeploy %s", uuid.NewV4())

	return fullName
}
