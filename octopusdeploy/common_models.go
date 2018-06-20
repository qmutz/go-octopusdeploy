package octopusdeploy

type PagedResults struct {
	ItemType       string `json:"ItemType"`
	TotalResults   int    `json:"TotalResults"`
	NumberOfPages  int    `json:"NumberOfPages"`
	LastPageNumber int    `json:"LastPageNumber"`
	ItemsPerPage   int    `json:"ItemsPerPage"`
	IsStale        bool   `json:"IsStale"`
	Links          Links  `json:"Links"`
}

type Links struct {
	Self        string `json:"Self"`
	Template    string `json:"Self"`
	PageAll     string `json:"Page.All"`
	PageCurrent string `json:"Page.Current"`
	PageLast    string `json:"Page.Last"`
	PageNext    string `json:"Page.Next"`
}

type SensitiveValue struct {
	HasValue bool   `json:"HasValue"`
	NewValue string `json:"NewValue"`
}

type PropertyValueResource struct {
	IsSensitive    bool           `json:"IsSensitive"`
	Value          string         `json:"Value"`
	SensitiveValue SensitiveValue `json:"SensitiveValue"`
}

type DeploymentStepResource struct {
	Id                 string                     `json:"Id"`
	Name               string                     `json:"Name"`
	PackageRequirement string                     `json:"PackageRequirement"` // may need its own model / enum
	Properties         PropertyValueResource      `json:"Properties"`         // may need its own model
	Condition          string                     `json:"Condition"`          // needs enum
	StartTrigger       string                     `json:"StartTrigger"`       // needs enum
	Actions            []DeploymentActionResource `json:"Actions"`
}

type DeploymentActionResource struct {
	Id                            string                `json:"Id"`
	Name                          string                `json:"Name"`
	ActionType                    string                `json:"ActionType"`
	IsDisabled                    bool                  `json:"IsDisabled"`
	CanBeUsedForProjectVersioning bool                  `json:"CanBeUsedForProjectVersioning"`
	Environments                  []string              `json:"Environments"`
	ExcludedEnvironments          []string              `json:"ExcludedEnvironments"`
	Channels                      []string              `json:"Channels"`
	TenantTags                    []string              `json:"TenantTags"`
	Properties                    PropertyValueResource `json:"Properties"`
	LastModifiedOn                string                `json:"LastModifiedOn"` // datetime
	LastModifiedBy                string                `json:"LastModifiedBy"`
	Links                         Links                 `json:"Links"` // may be wrong
}