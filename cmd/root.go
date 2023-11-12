package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gogeta",
	Short: "Gogeta generates valid Go types (structs) from JSON.",
	Long: `Gogeta is a CLI to generate valid Go types (structs) from JSON.
This tool takes any valid JSON file as input and generates valid Go types (structs).
The genereated types are written to the STDOUT or any file specified.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
