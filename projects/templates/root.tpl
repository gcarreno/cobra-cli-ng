package cmd

import (
{{- if .Viper }}
	"fmt"{{ end }}
	"os"

	"github.com/spf13/cobra"
{{- if .Viper }}
	"github.com/spf13/viper"{{ end }}
)

{{ if .Viper -}}
const (
	cFlagConfigShort   = "c"
	cFlagConfigLong    = "config"
	cFlagConfigDefault = ""
	cFlagConfigUsage   = "config file (default is $HOME/.{{ .Parent.Name }}.yaml)"
)
{{- end }}

var (
{{ if .Viper -}}
    cfgFile string
{{- end -}}
    // {{ .Name }}Cmd represents the base command when called without any subcommands
    {{ .Name }}Cmd = &cobra.Command{
    	Use:   "{{ .Parent.Name }}",
	    Short: "A brief description of your application",
	    Long: `A longer description that spans multiple lines. 
For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
        Example: `Add example`,
        // Uncomment the following line if your bare application
        // has an action associated with it:
        // Run: func(cmd *cobra.Command, args []string) { },
    }
)
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
{{- if .Viper }}
	cobra.OnInitialize(initConfig)
{{ end }}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
{{ if .Viper }}
	rootCmd.PersistentFlags().StringVarP(&cfgFile, cFlagConfigLong, cFlagConfigShort, cFlagConfigDefault, cFlagConfigUsage)
{{ else }}
	// rootCmd.PersistentFlags().StringVarP(&cfgFile, cFlagConfigLong, cFlagConfigShort, cFlagConfigDefault, cFlagConfigUsage)
{{ end }}
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

{{ if .Viper -}}
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".{{ .Parent.Name }}" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".{{ .Parent.Name }}")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
{{- end }}
