package projects

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gcarreno/cobra-cli-ng/utils"
	"github.com/spf13/cobra"
)

const (
	ProjectsFile = "cobra-cli-ng.json"
)

//go:embed templates/*.tpl
var templates embed.FS

type Projects struct {
	Module   string     `json:"module"`
	Path     string     `json:"path"`
	Projects []*Project `json:"projects"`
}

func NewProjects(force bool) (*Projects, error) {
	projects := &Projects{
		Projects: []*Project{},
	}

	goList, err := NewGoList()
	if err != nil {
		return nil, err
	}
	projects.Module = goList.Path
	projects.Path = goList.Dir

	filename := filepath.Join(projects.Path, ProjectsFile)
	// Check if file exists and avoid re-writing it
	if !force && utils.FileExists(filename) {
		return nil, fmt.Errorf("file '%s' already exists. use --force to overwrite", filename)
	}

	return projects, nil
}

func (ps *Projects) Add(project *Project) {
	ps.Projects = append(ps.Projects, project)
}

func (ps *Projects) Get(project string) *Project {
	for _, prj := range ps.Projects {
		if strings.EqualFold(project, prj.Name) {
			return prj
		}
	}

	return nil
}

func (ps *Projects) Delete(project *Project) {
	for index, prj := range ps.Projects {
		if prj == project {
			ps.Projects = append(ps.Projects[:index], ps.Projects[index+1:]...)
		}
	}
}

func (ps *Projects) Save(cmd *cobra.Command, force bool) error {
	if len(ps.Projects) < 1 {
		return fmt.Errorf("there are no projects to save")
	}

	filename := filepath.Join(ps.Path, ProjectsFile)
	// Create the file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Obtain a formatted JSON verion of the structure
	content, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		return err
	}

	// Save the JSON content to the file
	if _, err := f.Write(content); err != nil {
		return err
	}

	cmd.Printf("Saved projects to '%s'\n", filename)

	return nil
}

func (ps *Projects) Load(cmd *cobra.Command) error {
	filename := filepath.Join(ps.Path, ProjectsFile)
	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	// Load the structure
	if err := json.NewDecoder(f).Decode(ps); err != nil {
		return err
	}

	// Re-instate the parents
	for _, prj := range ps.Projects {
		prj.Parent = ps
		for _, cmd := range prj.Commands {
			cmd.Parent = prj
		}
	}

	cmd.Printf("Loaded projects from '%s'\n", filename)

	return nil
}
