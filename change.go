package teamcity

import (
	"fmt"
	"net/url"
)

type changeList struct {
	Count   int      `json:"count"`
	Changes []Change `json:"change"`
}

// Get change by its ID
func (c client) GetChangeByID(id int) (Change, error) {
	debugf("GetChangeByID(%d)", id)
	uri := fmt.Sprintf("/changes/id:%d", id)

	var change Change
	err := c.httpGet(uri, nil, &change)
	if err != nil {
		errorf("GetChangeByID(%d) failed: %s", id, err)
		return Change{}, err
	}

	debugf("GetChangeByID(%d): OK", id)
	return change, nil
}

// Get N latest changes
func (c client) GetChanges(count int) ([]Change, error) {
	debugf("GetChanges(%d)", count)
	args := url.Values{}
	args.Set("locator", fmt.Sprintf("count:%d", count))

	var list changeList
	err := c.httpGet("/changes", &args, &list)
	if err != nil {
		errorf("GetChanges(%d) failed with %s", count, err)
		return nil, err
	}

	debugf("GetChanges(%d): OK", count)
	return list.Changes, nil
}

func (c client) abstractListFillByLocator(locator, uri, debugFuncName string, result interface{}) error {
	debugf(debugFuncName)
	args := url.Values{}
	args.Set("locator", locator)

	err := c.httpGet(uri, &args, result)
	if err != nil {
		errorf("%s failed with %s", debugFuncName, err)
		return err
	}

	debugf("%s: OK", debugFuncName)
	return nil
}

// Get N latest changes for a project
func (c client) GetChangesForProject(id string, count int) ([]Change, error) {
	var list changeList
	err := c.abstractListFillByLocator(
		fmt.Sprintf("project:%s,count:%d", url.QueryEscape(id), count),
		"/changes",
		fmt.Sprintf("GetChangesForProject('%s', %d)", id, count),
		&list,
	)
	if err != nil {
		return nil, err
	}
	return list.Changes, nil
}

// Get changes for a build
func (c client) GetChangesForBuild(id int) ([]Change, error) {
	var list changeList
	err := c.abstractListFillByLocator(
		fmt.Sprintf("build:(id:%d)", id),
		"/changes",
		fmt.Sprintf("GetChangesForBuild(%d)", id),
		&list,
	)
	if err != nil {
		return nil, err
	}
	return list.Changes, nil
}

// Get changes for build type since a particular change
func (c client) GetChangesForBuildTypeSinceChange(btId string, cId int) ([]Change, error) {
	var list changeList
	err := c.abstractListFillByLocator(
		fmt.Sprintf("buildType:(id:%s),sinceChange:(id:%d)", url.QueryEscape(btId), cId),
		"/changes",
		fmt.Sprintf("GetChangesFotBuildTypeSinceChange('%s', %d)", btId, cId),
		&list,
	)
	if err != nil {
		return nil, err
	}
	return list.Changes, nil
}

// Get pending changes for build type
func (c client) GetChangesForBuildTypePending(id string) ([]Change, error) {
	var list changeList
	err := c.abstractListFillByLocator(
		fmt.Sprintf("buildType:(id:%s),pending:true", url.QueryEscape(id)),
		"/changes",
		fmt.Sprintf("GetChangesForBuildTypePending('%s')", id),
		&list,
	)
	if err != nil {
		return nil, err
	}
	return list.Changes, nil
}
