package cmd

import (
	"os"

	"github.com/danielmesquitta/wire-config/internal/usecase"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wire-config",
	Short: "Generate wire.go file from a config file",
	Long:  `wire-config is a tool to generate wire.go files for different environments from a config file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		mainPackageImport, err := cmd.Flags().GetString("main-package")
		if err != nil {
			return err
		}
		environmentNames, err := cmd.Flags().GetStringSlice("environments")
		if err != nil {
			return err
		}
		return usecase.ParseWireConfig(
			configFile,
			outputFile,
			mainPackageImport,
			environmentNames,
		)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("config", "c", "", "Path to the wire config file")
	rootCmd.Flags().StringP("output", "o", "", "Path to the output file")
	rootCmd.Flags().
		StringP("main-package", "m", "", "Import path of the main package")
	rootCmd.Flags().
		StringSliceP("environments", "e", []string{}, "Environment names to generate")
}
