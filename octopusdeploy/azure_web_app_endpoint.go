package octopusdeploy

// AzureWebAppEndpoint represents the an Azure web app-based endpoint.
type AzureWebAppEndpoint struct {
	ResourceGroupName string `json:"ResourceGroupName,omitempty"`
	WebAppName        string `json:"WebAppName,omitempty"`
	WebAppSlotName    int    `json:"WebAppSlotName"`

	endpoint
}

// NewAzureWebAppEndpoint creates a new endpoint for Azure web apps.
func NewAzureWebAppEndpoint() *AzureWebAppEndpoint {
	azureWebAppEndpoint := &AzureWebAppEndpoint{
		endpoint: *newEndpoint("AzureWebApp"),
	}

	return azureWebAppEndpoint
}

var _ IResource = &AzureWebAppEndpoint{}
var _ IEndpoint = &AzureWebAppEndpoint{}
