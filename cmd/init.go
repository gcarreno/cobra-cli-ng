package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/gcarreno/cobra-cli-ng/projects"
)

var (
	// initCmd represents the init command
	initCmd = &cobra.Command{
		Use:     "init [path]",
		Aliases: []string{"i", "initialize", "initialise", "create"},
		Short:   "Initialize a Cobra Application",
		Long: `Initialize (cobra-cli-ng init) will create a new application, with the 
appropriate structure for a Cobra-based CLI application.

This "init" command must be run inside of a go module (please run "go mod init <MODNAME>" first).`,
		Example: `  # Creating a new project with defaults: path is cmd
  cobra-cli-ng init

  # Creating a new project with defaults and viper: path is cmd
  cobra-cli-ng init --viper

  # Creating a new project named serverd, with path cli/serverd
  cobra-cli-ng init cli/serverd`,
		ValidArgsFunction: validArgsInit,
		PreRunE:           runInitPreE,
		RunE:              runInitE,
	}
)

// Make some validations on flags or args
func runInitPreE(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("too many arguments specified")
	}

	return nil
}

// The main body of the "init" command
func runInitE(cmd *cobra.Command, args []string) error {
	// New projects
	prjs, err := projects.NewProjects(force)
	// Internal error
	cobra.CheckErr(err)

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

	// Run "go mod tidy"
	cmd.Println("Executing 'go mod tidy'...")
	// Internal error
	out, err := exec.Command("go", "mod", "tidy").Output()
	cobra.CheckErr(err)
	if len(out) > 0 {
		cmd.Println("Done:")
		cmd.Println(string(out))
	} else {
		cmd.Println("Done.")
	}

	return nil
}

// A function to validate arguments
func validArgsInit(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	var directive cobra.ShellCompDirective
	if len(args) == 0 {
		comps = cobra.AppendActiveHelp(comps, "Optionally specify the path of the go module to initialize")
		directive = cobra.ShellCompDirectiveDefault
	} else if len(args) == 1 {
		comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments (but may accept flags)")
		directive = cobra.ShellCompDirectiveNoFileComp
	} else {
		comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
		directive = cobra.ShellCompDirectiveNoFileComp
	}
	return comps, directive
}
