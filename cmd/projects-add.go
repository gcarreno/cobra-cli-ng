package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/gcarreno/cobra-cli-ng/projects"
	"github.com/gcarreno/cobra-cli-ng/utils"
)

var (
	// projectsAddCmd represents the projects add command
	projectsAddCmd = &cobra.Command{
		Use:     "add [path]",
		Aliases: []string{"a"},
		Short:   "Add a project to the project list",
		Long: `Projects Add (cobra-cli-ng projects add) will create a new application, with the appropriate 
structure for a Cobra-based CLI application.

This "projects add" command must be run inside of a go module (please run "go mod init <MODNAME>" first)`,
		Example: `  # Adding a new project named "serverd", with path "cli/serverd"
  cobra-cli-ng projects add cli/serverd

  # Adding a new project named "serverd" and viper, with path "cli/serverd"
  cobra-cli-ng projects add cli/serverd --viper`,
		ValidArgsFunction: validArgsProjectsAdd,
		PreRunE:           runProjectsAddPreE,
		RunE:              runProjectsAddE,
	}
)

// Make some validations on flags or args
func runProjectsAddPreE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify the path for the new project")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments specified")
	}

	return nil
}

// The main body of the "projects add" command
func runProjectsAddE(cmd *cobra.Command, args []string) error {
	// Use projects.GoList to get the module's path
	goList, err := projects.NewGoList()
	if err != nil {
		// Internal error
		cobra.CheckErr(err)
	}

	projectsFile := filepath.Join(goList.Dir, projects.ProjectsFile)
	// Determine if we already have a projects file
	if !utils.FileExists(projectsFile) {
		// Internal error
		cobra.CheckErr(fmt.Errorf("can not find the '%s' file", projectsFile))
	}

	prjs := &projects.Projects{}
	// Load the projects
	// Internal error
	cobra.CheckErr(prjs.Load(cmd))

	if prjs.Projects[0].Path == "" {
		cmd.Println(`Warning: There's a project created with defaults.
  Please make sure the new project does not incur in errors.
  If need be, please delete the offending project.`)
	}

	var prjFolder string

	if len(args) > 0 && args[0] != "" {
		prjFolder = args[0]
	}

	if prjFolder == "." || prjFolder == "./" {
		prjFolder = ""
	}

	// New Project
	prj := projects.NewProject(
		prjs,                        // Parent
		prjFolder,                   // Path
		projects.DefaultCommandPath, // Command Path
	)

	// Internal error
	cobra.CheckErr(prj.Create(cmd, force))

	command := projects.NewCommand(
		prj,      // Parent
		"root",   // Command name
		"",       // Parent
		useViper, //Viper
	)

	// Internal error
	cobra.CheckErr(command.Create(cmd, force))

	prj.Add(command)
	prjs.Add(prj)

	// Internal error
	cobra.CheckErr(prjs.Save(cmd, force))

	return nil
}

// A function to validate arguments
func validArgsProjectsAdd(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	if len(args) == 0 {
		comps = cobra.AppendActiveHelp(comps, "Please specify the path for the new project")
	} else if len(args) == 1 {
		comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments (but may accept flags)")
	} else {
		comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
	}
	return comps, cobra.ShellCompDirectiveNoFileComp
}
