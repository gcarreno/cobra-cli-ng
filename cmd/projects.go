package cmd

import (
	"github.com/spf13/cobra"

	"github.com/gcarreno/cobra-cli-ng/projects"
)

var (
	// projectsCmd represents the projects command
	projectsCmd = &cobra.Command{
		Use:     "projects",
		Aliases: []string{"p"},
		Short:   "Lists the saved projects",
		Long:    `List (cobra-cli-ng projects) will list all the saved projects in "cobra-cli-ng.json".`,
		Run:     runProjects,
	}
)

// The main body of the "projects" command
func runProjects(cmd *cobra.Command, args []string) {
	prjs := projects.Projects{}
	if err := prjs.Load(cmd); err != nil {
		cobra.CheckErr(err)
	}
	cmd.Println("Available projects:")
	for _, prj := range prjs.Projects {
		cmd.Printf(" - %s: '%s'\n", prj.Name, prj.Path)
	}
}
