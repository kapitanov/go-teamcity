package teamcity

import "net/url"

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

	// Get change by its ID
	GetChangeByID(id int) (Change, error)
	// Get N latest changes
	GetChanges(count int) ([]Change, error)
	// Get N latest changes for a project
	GetChangesForProject(id string, count int) ([]Change, error)
	// Get changes for a build
	GetChangesForBuild(id int) ([]Change, error)
	// Get changes for build type since a particular change
	GetChangesForBuildTypeSinceChange(btId string, cId int) ([]Change, error)
	// Get pending changes for build type
	GetChangesForBuildTypePending(id string) ([]Change, error)
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

// BuildStatus is a build status enum
type BuildStatus int

const (
	// StatusUnknown is a zero value of BuildStatus
	StatusUnknown BuildStatus = iota

	// StatusSuccess is a status of successful build
	StatusSuccess

	// StatusRunning is a status of build that is currently running
	StatusRunning

	// StatusFailure is a status of failed build
	StatusFailure
)

// Build is a TeamCity project build
type Build struct {
	// Build ID
	ID int `json:"id"`
	// Build Number
	Number string `json:"number"`
	// Build Status
	Status BuildStatus `json:"status"`
	// Build Status Text
	StatusText string `json:"statusText"`
	// Build Progress Percentage
	Progress int `json:"progress"`
	// Build type ID
	BuildTypeID string `json:"buildTypeId"`
}

// Change is a TeamCity project change
type Change struct {
	// Change ID
	ID int `json:"id"`
	// VCS revision id
	Version string `json:"version"`
	// Change author username
	Username string `json:"username"`
	// Change date
	Date string `json:"date"`
}
