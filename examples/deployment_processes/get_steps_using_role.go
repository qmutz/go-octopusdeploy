package examples

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
)

func GetStepsUsingRoleExample() {
	var (
		// Declare working variables
		octopusURL    string = "https://youroctourl"
		octopusAPIKey string = "API-YOURAPIKEY"
		spaceName     string = "default"
		roleName      string = "My role"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)

	if err != nil {
		// TODO: handle error
	}

	// Get projects
	projects, err := client.Projects.GetAll()

	if err != nil {
		// TODO: handle error
	}

	// Loop through list
	for _, project := range projects {
		deploymentProcess, err := client.DeploymentProcesses.Get(project.DeploymentProcessID)

		if err != nil {
			// TODO: handle error
		}

		for _, step := range deploymentProcess.Steps {
			propertyValue := step.Properties["Octopus.Action.TargetRoles"]
			if len(propertyValue) > 0 {
				for _, role := range strings.Split(propertyValue, ",") {
					if strings.ToLower(role) == strings.ToLower(roleName) {
						fmt.Printf("Step [%s] from project [%s] is using the role [%s]\n", step.Name, project.Name, roleName)
					}
				}
			}
		}
	}
}
