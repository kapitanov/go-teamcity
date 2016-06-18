package teamcity

import (
	"fmt"
	"net/url"
)

type projectList struct {
	Count    int       `json:"count"`
	Projects []Project `json:"project"`
}

// Get a project by its ID
func (c client) GetProjectByID(id string) (Project, error) {
	debugf("GetProjectByID('%s')", id)
	uri := fmt.Sprintf("/projects/id:%s", url.QueryEscape(id))

	var project Project
	err := c.httpGet(uri, nil, &project)
	if err != nil {
		errorf("GetProjectByID('%s') failed: %s", id, err)
		return Project{}, err
	}

	debugf("GetProjectByID('%s'): OK", id)
	return project, nil
}

// Get a project by its name
func (c client) GetProjectByName(name string) (Project, error) {
	debugf("GetProjectByName('%s')", name)

	uri := fmt.Sprintf("/projects/name:%s", url.QueryEscape(name))

	var project Project
	err := c.httpGet(uri, nil, &project)
	if err != nil {
		errorf("GetProjectByName('%s') failed: %s", name, err)
		return Project{}, err
	}

	debugf("GetProjectByName('%s'): OK", name)
	return project, nil
}

// Get list of projects
func (c client) GetProjects() ([]Project, error) {
	debugf("GetProjects()")
	var list projectList
	err := c.httpGet("/projects", nil, &list)
	if err != nil {
		errorf("GetProjects() failed: %s", err)
		return nil, err
	}

	debugf("GetProjects(): OK")
	return list.Projects, nil
}
