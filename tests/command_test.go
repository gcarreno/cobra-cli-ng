package tests

import (
	"path/filepath"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/gcarreno/cobra-cli-ng/projects"
)

const (
	cmdName          = "test"
	cmdParentCommand = ""
	cmdFilename      = "test.go"
)

func TestCommand(t *testing.T) {
	// Test runs in parallel
	t.Parallel()

	// Create projects.Projects
	prjs, err := projects.NewProjects(true)

	// Assert no errors
	assert.NilError(t, err)

	// Create projects.Project
	prj := projects.NewProject(prjs, "", projects.DefaultCommandPath)

	// Create projects.Command
	command := projects.NewCommand(prj, cmdName, "", false)

	// Assert command.Filename is cmdFilename
	filename := filepath.Base(command.Filename)
	assert.Equal(t, cmdFilename, filename)

	// Assert command.Name is cmdName
	assert.Equal(t, cmdName, command.Name)

	// Assert command.ParentCommand is empty
	assert.Equal(t, cmdParentCommand, command.ParentCommand)

	// Assert command.Parent is prj
	assert.DeepEqual(t, command.Parent, prj)
}
