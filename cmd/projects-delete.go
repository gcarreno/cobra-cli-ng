package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/gcarreno/cobra-cli-ng/projects"
	"github.com/gcarreno/cobra-cli-ng/utils"
)

const (
	cFlagDeleteFilesShort   = "d"
	cFlagDeleteFilesLong    = "delete-files"
	cFlagDeleteFilesDefault = false
	cFlagDeleteFilesUsage   = "delete the project files"
)

var (
	// Will contain the boolean value of the persistent flag "delete-files"
	deleteFiles bool
	// projectsDeleteCmd represents the projects delete command
	projectsDeleteCmd = &cobra.Command{
		Use:     "delete [project name]",
		Aliases: []string{"d"},
		Short:   "Delete a project from the project list",
		Long: `Projects Delete (cobra-cli-ng projects delete) will delete a project from the project list.

To find what projects exist in the list, please use "cobra-cli-ng projects".`,
		Example: `  # Delete a project named "serverd" from the list
  cobra-cli-ng projects delete serverd`,
		ValidArgsFunction: validArgsProjectsDelete,
		PreRunE:           runProjectsDeletePreE,
		RunE:              runProjectsDeleteE,
	}
)

func init() {
	// Flag delete files
	projectsDeleteCmd.Flags().BoolVarP(&deleteFiles, cFlagDeleteFilesLong, cFlagDeleteFilesShort, cFlagDeleteFilesDefault, cFlagDeleteFilesUsage)
}

// Make some validations on flags or args
func runProjectsDeletePreE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please specify the project name to delete")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments specified")
	}

	return nil
}

// The main body of the "projects delete" command
func runProjectsDeleteE(cmd *cobra.Command, args []string) error {
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

	if len(prjs.Projects) == 1 {
		// Internal error
		cobra.CheckErr(fmt.Errorf(`cannot delete the only project left.
  Please add a new project, and then you can delete '%s'`, args[0]))
	}

	prj := prjs.Get(args[0])

	if prj == nil {
		// Internal error
		cobra.CheckErr(fmt.Errorf(`could not find project '%s'.
  Please use "cobra-cli-ng projects" to get a list of projects`, args[0]))
	}

	if deleteFiles {
		cmd.Println("Deleting files...")
		// TODO: Delete files
		folder := filepath.Join(prjs.Path, prj.Path)
		if err := os.RemoveAll(folder); err != nil {
			cobra.CheckErr(err)
		}
		cmd.Println("Done.")
	}

	prjs.Delete(prj)

	cobra.CheckErr(prjs.Save(cmd, true))

	cmd.Printf("Project '%s' has been deleted\n", args[0])

	return nil
}

func validArgsProjectsDelete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string
	if len(args) == 0 {
		comps = cobra.AppendActiveHelp(comps, "Please specify the name of the project to delete")
	} else if len(args) == 1 {
		comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments (but may accept flags)")
	} else {
		comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
	}
	return comps, cobra.ShellCompDirectiveNoFileComp
}
