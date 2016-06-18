package teamcity

import (
	"fmt"
	"net/url"
)

type buildList struct {
	Count  int     `json:"count"`
	Builds []Build `json:"build"`
}

// Get build by its ID
func (c client) GetBuildByID(id int) (Build, error) {
	debugf("GetBuildByID(%d)", id)
	uri := fmt.Sprintf("/builds/id:%d", id)

	var build Build
	err := c.httpGet(uri, nil, &build)
	if err != nil {
		errorf("GetBuildByID(%d) failed with %s", id, err)
		return Build{}, err
	}

	debugf("GetBuildByID(%d): OK", id)
	return build, nil
}

// Get N latest builds
func (c client) GetBuilds(count int) ([]Build, error) {
	debugf("GetBuilds(%d)", count)
	args := url.Values{}
	args.Set("locator", fmt.Sprintf("count:%d,running:any", count))

	var list buildList
	err := c.httpGet("/builds", &args, &list)
	if err != nil {
		errorf("GetBuilds(%d) failed with %s", count, err)
		return nil, err
	}

	debugf("GetBuilds(%d): OK", count)
	return list.Builds, nil
}

// Get N latest builds for a build type
func (c client) GetBuildsForBuildType(buildType BuildType, count int) ([]Build, error) {
	debugf("GetBuildsForBuildType('%s', %d)", buildType.ID, count)
	args := url.Values{}
	args.Set("locator", fmt.Sprintf("buildType:%s,count:%d,running:any", url.QueryEscape(buildType.ID), count))

	var list buildList
	err := c.httpGet("/builds", &args, &list)
	if err != nil {
		errorf("GetBuildsForBuildType('%s', %d) failed with %s", buildType.ID, count, err)
		return nil, err
	}

	debugf("GetBuildsForBuildType('%s', %d): OK", buildType.ID, count)
	return list.Builds, nil
}
