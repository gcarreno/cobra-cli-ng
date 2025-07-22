package projects

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/gcarreno/cobra-cli-ng/utils"
	"github.com/spf13/cobra"
)

type Project struct {
	Name         string     `json:"name"`
	Path         string     `json:"path"`
	CommandsPath string     `json:"commands-path"`
	Commands     []*Command `json:"commands"`
	Parent       *Projects  `json:"-"` // Needed to avoid repeating fields
}

func NewProject(parent *Projects, path string, cmdPath string) *Project {
	// Find the name of the folder containing this module
	var name string
	if path == "" {
		name = filepath.Base(parent.Path)
	} else {
		name = filepath.Base(path)
	}

	project := &Project{
		Name:         name,
		Path:         path,
		CommandsPath: cmdPath,
		Commands:     []*Command{},
		Parent:       parent,
	}

	return project
}

func (p *Project) Create(cmd *cobra.Command, force bool) error {
	filename := filepath.Join(p.Parent.Path, p.Path, "main.go")
	// Check if file exists and avoid re-writing it
	if !force && utils.FileExists(filename) {
		return fmt.Errorf("file '%s' already exists. use --force to overwrite", filename)
	}

	// Read the template
	tpl_content, err := templates.ReadFile("templates/main.tpl")
	if err != nil {
		return err
	}

	// Parse the template
	tpl, err := template.New("root").Parse(string(tpl_content))
	if err != nil {
		return err
	}

	dir := filepath.Join(p.Parent.Path, p.Path)
	// Make sure we have all the folder structure in place
	if err := utils.EnsureDir(dir, 0775); err != nil {
		return err
	}

	// Create "main.go"
	mainFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer mainFile.Close()

	// Execute the template
	err = tpl.Execute(mainFile, p)
	if err != nil {
		return err
	}

	cmd.Printf("Created file: '%s'\n", filename)

	cmd.Printf("Executing 'gofmt -s -w %s'...\n", filename)
	// Running "gofmt -s -w filename"
	cobra.CheckErr(exec.Command("gofmt", "-s", "-w", filename).Run())
	cmd.Println("Done.")

	return nil
}

func (p *Project) Add(command *Command) {
	p.Commands = append(p.Commands, command)
}
