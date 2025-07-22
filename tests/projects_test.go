package tests

import (
	"path/filepath"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/gcarreno/cobra-cli-ng/projects"
)

const (
	module = "github.com/gcarreno/cobra-cli-ng"
	path   = "cobra-cli-ng"
)

func TestProjects(t *testing.T) {
	// Test runs in parallel
	t.Parallel()

	// Create projects.Projects
	prjs, err := projects.NewProjects(true)

	// Assert we have no error
	assert.NilError(t, err)

	// Assert Module
	assert.Equal(t, module, prjs.Module)

	// Assert Path
	basePath := filepath.Base(prjs.Path)
	assert.Equal(t, path, basePath)

	// Assert we have zero projects
	assert.Equal(t, 0, len(prjs.Projects))
}
