package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/gcarreno/cobra-cli-ng/version"
)

const (
	cFlagConfigShort   = "c"
	cFlagConfigLong    = "config"
	cFlagConfigDefault = ""
	cFlagConfigUsage   = "config file (default is $HOME/.cobra-cli-ng.yaml)"

	cFlagForceShort   = "f"
	cFlagForceLong    = "force"
	cFlagForceDefault = false
	cFlagForceUsage   = "force overwriting files"

	cFlagViperShort   = "v"
	cFlagViperLong    = "viper"
	cFlagViperDefault = false
	cFlagViperUsage   = "use viper for configuration"
)

var (
	// Will contain the string value of the persistent flag "config" if passed
	cfgFile string
	// Will contain the boolean value of the persistent flag "force"
	force bool
	// Will contain the boolean value of persistent flag "viper"
	useViper bool
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:     "cobra-cli-ng",
		Version: version.NewVersion().AsString(),
		Short:   "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

// Register commands, persistent or local flags
func init() {
	// Global call for viper config setup
	cobra.OnInitialize(initConfig)

	// Config File flag
	rootCmd.PersistentFlags().StringVarP(&cfgFile, cFlagConfigLong, cFlagConfigShort, cFlagConfigDefault, cFlagConfigUsage)

	// Force flag
	rootCmd.PersistentFlags().BoolVarP(&force, cFlagForceLong, cFlagForceShort, cFlagForceDefault, cFlagForceUsage)

	// Viper flag
	rootCmd.PersistentFlags().BoolVarP(&useViper, cFlagViperLong, cFlagViperShort, cFlagViperDefault, cFlagViperUsage)

	// Register our own usage function
	rootCmd.SetUsageFunc(usage)

	// Register all the available commands, with hierarchy
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)

	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(projectsAddCmd)
	projectsCmd.AddCommand(projectsDeleteCmd)
}

// Viper config setup
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra-cli-ng" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra-cli-ng")
	}

	// Load .env file
	err := godotenv.Load() // Automatically loads ".env"
	if err != nil {
		//initCmd.Println("No .env file found (that's okay)")
	}

	// Enable ENV binding
	viper.SetEnvPrefix("COBRACLING")
	viper.AutomaticEnv()

	// Make env vars like COBRACLING_SUB_SOMETHING map to "sub.something"
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()

	notFound := &viper.ConfigFileNotFoundError{}
	switch {
	case err != nil && !errors.As(err, notFound):
		cobra.CheckErr(err)
	case err != nil && errors.As(err, notFound):
		// The config file is optional, we shouldn't exit when the config is not found
		break
	default:
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// This will print our own version of the usage
// IMPORTANT: This has to be checked with the default usage function to maintain
// the same result that is printed by cobra's default usage function and the
// default usage template
func usage(cmd *cobra.Command) error {
	// Use cmd.Print and others to print to the same place cobra does
	cmd.Print("\033[1mUSAGE\033[0m")
	if cmd.Runnable() {
		cmd.Printf("\n  %s", cmd.UseLine())
	}
	if cmd.HasAvailableSubCommands() {
		cmd.Printf("\n  %s [command]", cmd.CommandPath())
	}
	if len(cmd.Aliases) > 0 {
		cmd.Printf("\n\n\033[1mALIASES\033[0m\n")
		cmd.Printf("  %s", cmd.NameAndAliases())
	}
	if cmd.HasExample() {
		cmd.Printf("\n\n\033[1mEXAMPLES\033[0m\n")
		cmd.Printf("%s", cmd.Example)
	}
	if cmd.HasAvailableSubCommands() {
		cmds := cmd.Commands()
		if len(cmd.Groups()) == 0 {
			cmd.Printf("\n\n\033[1mAVAILABLE COMMANDS\033[0m")
			for _, subcmd := range cmds {
				if subcmd.IsAvailableCommand() || subcmd.Name() == "help" {
					cmd.Printf("\n  %s %s", rpad(subcmd.Name(), subcmd.NamePadding()), subcmd.Short)
				}
			}
		} else {
			for _, group := range cmd.Groups() {
				cmd.Printf("\n\n%s", group.Title)
				for _, subcmd := range cmds {
					if subcmd.GroupID == group.ID && (subcmd.IsAvailableCommand() || subcmd.Name() == "help") {
						cmd.Printf("\n  %s %s", rpad(subcmd.Name(), subcmd.NamePadding()), subcmd.Short)
					}
				}
			}
			if !cmd.AllChildCommandsHaveGroup() {
				cmd.Printf("\n\n\033[1mADDITIONAL COMMANDS\033[0m")
				for _, subcmd := range cmds {
					if subcmd.GroupID == "" && (subcmd.IsAvailableCommand() || subcmd.Name() == "help") {
						cmd.Printf("\n  %s %s", rpad(subcmd.Name(), subcmd.NamePadding()), subcmd.Short)
					}
				}
			}
		}
	}
	if cmd.HasAvailableLocalFlags() {
		cmd.Printf("\n\n\033[1mFLAGS\033[0m\n")
		cmd.Print(trimRightSpace(cmd.LocalFlags().FlagUsages()))
	}
	if cmd.HasAvailableInheritedFlags() {
		cmd.Printf("\n\n\033[1mGLOBAL FLAGS\033[0m\n")
		cmd.Print(trimRightSpace(cmd.InheritedFlags().FlagUsages()))
	}
	if cmd.HasHelpSubCommands() {
		cmd.Printf("\n\n\033[1mADDITIONAL HELP TOPICS\033[0m")
		for _, subcmd := range cmd.Commands() {
			if subcmd.IsAdditionalHelpTopicCommand() {
				cmd.Printf("\n  %s %s", rpad(subcmd.CommandPath(), subcmd.CommandPathPadding()), subcmd.Short)
			}
		}
	}

	if cmd.HasAvailableSubCommands() {
		cmd.Printf("\n\nUse \"%s [command] --help\" for more information about a command.", cmd.CommandPath())
	}
	cmd.Println()
	return nil
}

// Helper function to trim spaces at the right of a string
func trimRightSpace(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// Helper function to right pad a string with spaces
func rpad(s string, padding int) string {
	formattedString := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(formattedString, s)
}
