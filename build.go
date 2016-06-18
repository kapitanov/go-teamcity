package teamcity

import (
	"fmt"
	"net/url"
)

type buildListItem struct {
	ID          int    `json:"id"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	StatusText  string `json:"statusText"`
	Running     bool   `json:"running"`
	Progress    int    `json:"percentageComplete"`
	BuildTypeID string `json:"buildTypeId"`
}

type buildList struct {
	Count  int             `json:"count"`
	Builds []buildListItem `json:"build"`
}

// Get build by its ID
func (c client) GetBuildByID(id int) (Build, error) {
	debugf("GetBuildByID(%d)", id)
	uri := fmt.Sprintf("/builds/id:%d", id)

	var build buildListItem
	err := c.httpGet(uri, nil, &build)
	if err != nil {
		errorf("GetBuildByID(%d) failed with %s", id, err)
		return Build{}, err
	}

	debugf("GetBuildByID(%d): OK", id)
	return createBuildFromJSON(build), nil
}

// Get N latest builds
func (c client) GetBuilds(count int) ([]Build, error) {
	debugf("GetBuilds(%d)", count)
	args := url.Values{}
	args.Set("locator", fmt.Sprintf("count:%d,running:any", count))
	args.Set("fields", "build(id,number,status,state,buildTypeId,statusText,running,percentageComplete)")

	var list buildList
	err := c.httpGet("/builds", &args, &list)
	if err != nil {
		errorf("GetBuilds(%d) failed with %s", count, err)
		return nil, err
	}

	debugf("GetBuilds(%d): OK", count)
	return createBuildsFromJSON(list.Builds), nil
}

// Get N latest builds for a build type
func (c client) GetBuildsForBuildType(id string, count int) ([]Build, error) {
	debugf("GetBuildsForBuildType('%s', %d)", id, count)
	args := url.Values{}
	args.Set("locator", fmt.Sprintf("buildType:%s,count:%d,running:any", url.QueryEscape(id), count))
	args.Set("fields", "build(id,number,status,state,buildTypeId,statusText,running,percentageComplete)")

	var list buildList
	err := c.httpGet("/builds", &args, &list)
	if err != nil {
		errorf("GetBuildsForBuildType('%s', %d) failed with %s", id, count, err)
		return nil, err
	}

	debugf("GetBuildsForBuildType('%s', %d): OK", id, count)
	return createBuildsFromJSON(list.Builds), nil
}

func createBuildFromJSON(item buildListItem) Build {
	status := StatusFailure
	if item.Running {
		status = StatusRunning
	} else if item.Status == "SUCCESS" {
		status = StatusSuccess
	}

	return Build{
		ID:          item.ID,
		Number:      item.Number,
		Status:      status,
		StatusText:  item.StatusText,
		Progress:    item.Progress,
		BuildTypeID: item.BuildTypeID,
	}
}

func createBuildsFromJSON(items []buildListItem) []Build {
	builds := make([]Build, len(items))

	for i, build := range items {
		builds[i] = createBuildFromJSON(build)
	}

	return builds
}
