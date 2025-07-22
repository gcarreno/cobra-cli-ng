package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gcarreno/cobra-cli-ng/projects"
)

const (
	cFlagProjectLong    = "project"
	cFlagProjectUsage   = "a project name"
	cFlagProjectDefault = ""

	cFlagParentLong    = "parent"
	cFlagParentUsage   = "a parent command name"
	cFlagParentDefault = ""
)

var (
	// Will contain the string value for the flag "project"
	proj string
	// Will contain the string value for the flag "parent"
	parent string
	// addCmd represents the add command
	addCmd = &cobra.Command{
		Use:     "add [command name]",
		Aliases: []string{"a", "command"},
		Short:   "Add a command to a Cobra Application",
		Long: `Add (cobra-cli-ng add) will create a new command, with the 
appropriate structure for a Cobra-based CLI application, and 
register it to its parent (default rootCmd).

Unless you use the "--project" flag, it will add to the first project in "cobra-cli-ng.json".
To find what projects exist in the list, please use "cobra-cli-ng projects".

If you want your command to be public, pass in the command name
with an initial uppercase letter.`,
		Example: `  # Adding a new command named server, resulting in a new cmd/server.go
  cobra-cli-ng add server
`,
		ValidArgsFunction: validArgsAdd,
		PreRunE:           runAddPreE,
		RunE:              runAddE,
	}
)

func init() {
	// Flag project
	addCmd.Flags().StringVar(&proj, cFlagProjectLong, cFlagProjectDefault, cFlagProjectUsage)

	// Flag parent
	addCmd.Flags().StringVar(&parent, cFlagParentLong, cFlagParentDefault, cFlagParentUsage)
}

// Make some validations on flags or args
func runAddPreE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify the name for the new command")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments specified")
	}

	return nil
}

// The main body of the "add" command
func runAddE(cmd *cobra.Command, args []string) error {
	prjs := &projects.Projects{}

	// Internal error
	cobra.CheckErr(prjs.Load(cmd))

	var prj *projects.Project
	if proj == "" {
		prj = prjs.Projects[0]
	} else {
		prj = prjs.Get(proj)
		if prj == nil {
			cobra.CheckErr(fmt.Errorf("cannot find project '%s'", proj))
		}
	}

	// Test if command already exists
	if !force {
		for _, cmd := range prj.Commands {
			if strings.EqualFold(args[0], cmd.Name) {
				cobra.CheckErr(fmt.Errorf("command '%s' already exists. use --force to overwrite", args[0]))
			}
		}
	}

	if parent == "" {
		parent = projects.DefaultRootCommand
	}

	command := projects.NewCommand(prj, args[0], parent, useViper)

	// Internal error
	cobra.CheckErr(command.Create(cmd, force))

	prj.Add(command)

	cobra.CheckErr(prjs.Save(cmd, force))

	return nil
}

// A function to validate arguments
func validArgsAdd(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	if len(args) == 0 {
		comps = cobra.AppendActiveHelp(comps, "Please specify the name for the new command")
	} else if len(args) == 1 {
		comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments (but may accept flags)")
	} else {
		comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
	}
	return comps, cobra.ShellCompDirectiveNoFileComp
}
