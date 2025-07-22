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

const (
	DefaultCommandPath = "cmd"
	DefaultRootCommand = "root"
)

type Command struct {
	Filename      string   `json:"filename"`
	Name          string   `json:"name"`
	ParentCommand string   `json:"parent"`
	Viper         bool     `json:"viper"`
	Parent        *Project `json:"-"` // Needed to avoid repeating fields
}

func NewCommand(parent *Project, cmdName string, parentCmdName string, viper bool) *Command {
	filename := filepath.Join(
		parent.Parent.Path,
		parent.Path,
		parent.CommandsPath,
		fmt.Sprintf("%s.go", cmdName),
	)

	return &Command{
		Filename:      filename,
		Name:          cmdName,
		ParentCommand: parentCmdName,
		Viper:         viper,
		Parent:        parent,
	}
}

func (c *Command) Create(cmd *cobra.Command, force bool) error {
	if !force && utils.FileExists(c.Filename) {
		return fmt.Errorf("file '%s' already exists. use --force to overwrite", c.Filename)
	}
	var tpl_content []byte
	var err error
	if c.Name == DefaultRootCommand {
		// Read the template
		tpl_content, err = templates.ReadFile("templates/root.tpl")
		if err != nil {
			return err
		}
	} else {
		// Read the template
		tpl_content, err = templates.ReadFile("templates/command.tpl")
		if err != nil {
			return err
		}
	}

	// Parse the template
	tpl, err := template.New(c.Name).Parse(string(tpl_content))
	if err != nil {
		return err
	}

	dir := filepath.Dir(c.Filename)
	// Make sure we have all the folder structure in place
	if err := utils.EnsureDir(dir, 0775); err != nil {
		return err
	}

	// Create file
	cmdFile, err := os.Create(c.Filename)
	if err != nil {
		return err
	}
	defer cmdFile.Close()

	// Execute the template
	err = tpl.Execute(cmdFile, c)
	if err != nil {
		return err
	}

	cmd.Printf("Created file: '%s'\n", c.Filename)

	cmd.Printf("Executing 'gofmt -s -w %s'...\n", c.Filename)
	// Running "gofmt -s -w filename"
	cobra.CheckErr(exec.Command("gofmt", "-s", "-w", c.Filename).Run())
	cmd.Println("Done.")

	return nil
}
