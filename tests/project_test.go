package tests

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/gcarreno/cobra-cli-ng/projects"
)

const (
	prjName         = "cobra-cli-ng"
	prjPath         = ""
	prjCommandsPath = "cmd"
)

func TestProject(t *testing.T) {
	// Test runs in parallel
	t.Parallel()

	// Create projects.Projects
	prjs, err := projects.NewProjects(true)

	// Assert no errors
	assert.NilError(t, err)

	// Create projects.Project
	prj := projects.NewProject(prjs, "", projects.DefaultCommandPath)

	// Assert prj.name is prjName
	assert.Equal(t, prjName, prj.Name)

	// Assert prj.Path is empty
	assert.Equal(t, prjPath, prj.Path)

	// Assert prj.Path is empty
	assert.Equal(t, prjCommandsPath, prj.CommandsPath)

	// Assert we have 0 commands
	assert.Equal(t, 0, len(prj.Commands))

	// Assert Project.Parent is prjs
	assert.DeepEqual(t, prjs, prj.Parent)
}
