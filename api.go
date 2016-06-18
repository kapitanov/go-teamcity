package teamcity

import "net/url"

const (
	// StatusSuccess means successful build status
	StatusSuccess = "SUCCESS"
	// StatusFailure means failed build status
	StatusFailure = "FAILURE"
	// StatusError means error build status
	StatusError = "ERROR"
)

// Authorizer is a TeamCity client authorizer
type Authorizer interface {
	// ResolveUrl provides a full absolute root URL.
	// It should use the following format: baseURL + PREFIX.
	// PREFIX might be either "/guestAuth/app/rest" or "/httpAuth/app/rest" depending on authorization mode.
	ResolveBaseURL(baseURL string) string

	// GetUserInfo provides credentials for HTTP basic auth.
	// It returns nil for guest access mode.
	GetUserInfo() *url.Userinfo
}

// Client is a TeamCity client
type Client interface {
	// Get a project by its ID
	GetProjectByID(id string) (Project, error)
	// Get a project by its name
	GetProjectByName(name string) (Project, error)
	// Get list of projects
	GetProjects() ([]Project, error)

	// Get build type by its ID
	GetBuildTypeByID(id string) (BuildType, error)
	// Get list of all build types
	GetBuildTypes() ([]BuildType, error)
	// Get list of build types for a project
	GetBuildTypesForProject(id string) ([]BuildType, error)

	// Get build by its ID
	GetBuildByID(id int) (Build, error)
	// Get N latest builds
	GetBuilds(count int) ([]Build, error)
	// Get N latest builds for a build type
	GetBuildsForBuildType(id string, count int) ([]Build, error)
}

// Project is a TeamCity project
type Project struct {
	// Project ID
	ID string `json:"id"`
	// Project name
	Name string `json:"name"`
	// Project description
	Description string `json:"description"`
	// Parent project ID
	ParentProjectID string `json:"parentProjectId"`
}

// BuildType is a TeamCity project build configuration
type BuildType struct {
	// Project ID
	ID string `json:"id"`
	// Project name
	Name string `json:"name"`
	// Project description
	Description string `json:"description"`
	// Project ID
	ProjectID string `json:"projectId"`
}

// Build is a TeamCity project build
type Build struct {
	// Build ID
	ID int `json:"id"`
	// Build Number
	Number string `json:"number"`
	// Build Status
	Status string `json:"status"`
	// Build State
	State string `json:"state"`
	// Build type ID
	BuildTypeID string `json:"buildTypeId"`
}
