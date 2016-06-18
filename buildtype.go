package teamcity

import (
	"fmt"
	"net/url"
)

type buildTypeList struct {
	Count      int         `json:"count"`
	BuildTypes []BuildType `json:"buildType"`
}

// Get build type by its ID
func (c client) GetBuildTypeByID(id string) (BuildType, error) {
	debugf("GetBuildTypeByID('%s')", id)
	uri := fmt.Sprintf("/buildTypes/id:%s", url.QueryEscape(id))

	var buildType BuildType
	err := c.httpGet(uri, nil, &buildType)
	if err != nil {
		errorf("GetBuildTypeByID('%s') failed with %s", id, err)
		return BuildType{}, err
	}

	debugf("GetBuildTypeByID('%s'): OK", id)
	return buildType, nil
}

// Get list of all build types
func (c client) GetBuildTypes() ([]BuildType, error) {
	debugf("GetBuildTypes()")
	var list buildTypeList
	err := c.httpGet("/buildTypes", nil, &list)
	if err != nil {
		errorf("GetBuildTypes() failed with %s", err)
		return nil, err
	}

	debugf("GetBuildTypes(): OK")
	return list.BuildTypes, nil
}

// Get list of build types for a project
func (c client) GetBuildTypesForProject(project Project) ([]BuildType, error) {
	debugf("GetBuildTypesForProject('%s')", project.ID)
	args := url.Values{}
	args.Set("locator", fmt.Sprintf("project:%s", url.QueryEscape(project.ID)))

	var list buildTypeList
	err := c.httpGet("/buildTypes", &args, &list)
	if err != nil {
		errorf("GetBuildTypesForProject('%s') failed with %s", project.ID, err)
		return nil, err
	}

	debugf("GetBuildTypesForProject('%s'): OK", project.ID)
	return list.BuildTypes, nil
}
