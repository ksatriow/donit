package cmd

import (
	"fmt"
	"os"
	"donit/pkg/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "donit",
	Short: "Donit is a CLI for initializing Docker projects",
	Long:  `Donit helps you initialize Docker and Docker Compose files for various programming languages.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Donit",
	Long:  `All software has versions. This is Donit's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.GetBuildInfo())
	},
}
