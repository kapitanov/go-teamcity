package teamcity_test

import (
	"testing"

	. "github.com/kapitanov/go-teamcity"
)

const (
	TC_URL                  = "https://teamcity.jetbrains.com"
	PROJECT_ID              = "TeamCityRestApiClients_RubyClient"
	PROJECT_WITH_CHANGES_ID = "TeamCityThirdPartyPlugins_TeamCityGithub"
	PROJECT_NAME            = "Femah"
	BUILD_TYPE_ID           = "TeamCityRestApiClients_RubyClient_BuildGem"
)

func Test_GetProjects(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	projects, err := client.GetProjects()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(projects) == 0 {
		t.Error("No projects have been fetched")
		t.Fail()
		return
	}

	for _, p := range projects {
		if p.ID == "" {
			t.Error("Got a projects with no ID")
			t.Fail()
			return
		}
	}
}

func Test_GetProjectById(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	project, err := client.GetProjectByID(PROJECT_ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if project.ID == "" {
		t.Error("Got a project with no ID")
		t.Fail()
		return
	}
}

func Test_GetProjectByName(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	project, err := client.GetProjectByName(PROJECT_NAME)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if project.ID == "" {
		t.Error("Got a project with no ID")
		t.Fail()
		return
	}
}

func Test_GetBuildTypes(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	buildTypes, err := client.GetBuildTypes()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(buildTypes) == 0 {
		t.Error("No build types have been fetched")
		t.Fail()
		return
	}

	for _, bt := range buildTypes {
		if bt.ID == "" {
			t.Error("Got a build type with no ID")
			t.Fail()
			return
		}
	}
}

func Test_GetBuildTypeById(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	buildType, err := client.GetBuildTypeByID(BUILD_TYPE_ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if buildType.ID == "" {
		t.Error("Got a build type with no ID")
		t.Fail()
		return
	}
}

func Test_GetBuildTypesForProject(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	project, err := client.GetProjectByID(PROJECT_ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	buildTypes, err := client.GetBuildTypesForProject(project.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(buildTypes) == 0 {
		t.Error("No build types have been fetched")
		t.Fail()
		return
	}

	for _, bt := range buildTypes {
		if bt.ID == "" {
			t.Error("Got a build type with no ID")
			t.Fail()
			return
		}
	}
}

func Test_GetBuilds(t *testing.T) {
	const COUNT = 1
	client := NewClient(TC_URL, GuestAuth())
	builds, err := client.GetBuilds(COUNT)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(builds) == 0 {
		t.Error("Got no builds")
		t.Fail()
		return
	}

	if len(builds) > COUNT {
		t.Error("Got too many builds")
		t.Fail()
		return
	}
}

func Test_GetBuildById(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	builds, err := client.GetBuilds(1)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(builds) == 0 {
		t.Error("Got no builds")
		t.Fail()
		return
	}

	build, err := client.GetBuildByID(builds[0].ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if build.ID == 0 {
		t.Error("Got build with no ID")
		t.Fail()
		return
	}
}

func Test_GetChanges(t *testing.T) {
	const COUNT = 1
	client := NewClient(TC_URL, GuestAuth())
	changes, err := client.GetChanges(COUNT)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(changes) == 0 {
		t.Error("Got no changes")
		t.Fail()
		return
	}

	if len(changes) > COUNT {
		t.Error("Got too many changes")
		t.Fail()
		return
	}

	for _, change := range changes {
		if change.ID == 0 {
			t.Error("Got change with no ID")
			t.Fail()
			return
		}
	}
}

func Test_GetChangeByID(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	changes, err := client.GetChanges(1)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(changes) == 0 {
		t.Error("Got no changes")
		t.Fail()
		return
	}

	change, err := client.GetChangeByID(changes[0].ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if change.ID == 0 {
		t.Error("Got change with no ID")
		t.Fail()
		return
	}
}

func Test_GetChangesForProject(t *testing.T) {
	const COUNT = 1
	client := NewClient(TC_URL, GuestAuth())
	project, err := client.GetProjectByID(PROJECT_WITH_CHANGES_ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	changes, err := client.GetChangesForProject(project.ID, COUNT)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(changes) == 0 {
		t.Error("Got no changes")
		t.Fail()
		return
	}

	if len(changes) > COUNT {
		t.Error("Got too many changes")
		t.Fail()
		return
	}

	for _, change := range changes {
		if change.ID == 0 {
			t.Error("Got change with no ID")
			t.Fail()
			return
		}
	}
}

func Test_GetChangesForBuild(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	builds, err := client.GetBuilds(1)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(builds) == 0 {
		t.Error("Got no builds")
		t.Fail()
		return
	}

	changes, err := client.GetChangesForBuild(builds[0].ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(changes) == 0 {
		t.Error("Gon no changes")
		t.Fail()
		return
	}

	for _, change := range changes {
		if change.ID == 0 {
			t.Error("Got change with no ID")
			t.Fail()
			return
		}
	}
}

func Test_GetChangesForBuildTypeSinceChange(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	project, err := client.GetProjectByID(PROJECT_WITH_CHANGES_ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	buildTypes, err := client.GetBuildTypesForProject(project.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(buildTypes) == 0 {
		t.Error("Got no build types")
		t.Fail()
		return
	}

	prChanges, err := client.GetChangesForProject(project.ID, 5)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(prChanges) == 0 {
		t.Error("Got no changes for project")
		t.Fail()
		return
	}

	changes, err := client.GetChangesForBuildTypeSinceChange(buildTypes[0].ID, prChanges[len(prChanges)-1].ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(changes) < len(prChanges)-1 {
		t.Error("Got to few changes")
		t.Fail()
		return
	}
}

func Test_GetChangesForBuildTypePending(t *testing.T) {
	client := NewClient(TC_URL, GuestAuth())
	project, err := client.GetProjectByID(PROJECT_WITH_CHANGES_ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	buildTypes, err := client.GetBuildTypesForProject(project.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(buildTypes) == 0 {
		t.Error("Got no build types")
		t.Fail()
		return
	}

	_, err = client.GetChangesForBuildTypePending(buildTypes[0].ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}
