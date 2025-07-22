package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
    // {{ .Name }}Cmd represents the {{ .Name }} command
    {{ .Name }}Cmd = &cobra.Command{
	    Use:   "{{ .Name }}",
	    Short: "A brief description of your command",
	    Long: `A longer description that spans multiple lines. 
For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
        Example: `Add example`,
	    Run: func(cmd *cobra.Command, args []string) {
		    fmt.Println("{{ .Name }} called")
	    },
    }
)

func init() {
	{{ .ParentCommand }}Cmd.AddCommand({{ .Name }}Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// {{ .Name }}Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// {{ .Name }}Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
