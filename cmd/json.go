package cmd

import (
	"fmt"
	"io"
	"os"

	gogeta "github.com/rajenderK7/gogeta/internal/gogeta_json"
	"github.com/spf13/cobra"
)

// If output file is not provided then we fallback to default.
// However, a better default output file should be chosen.
const defaultOutputFile string = "gogeta.txt"

var (
	inputFile    string
	outputFile   string
	outputToFile bool
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "The json command tells gogeta to convert JSON to Go types",
	Run: func(cmd *cobra.Command, args []string) {
		geneartedType, err := gogeta.GenerateGoStruct(inputFile)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if !outputToFile && outputFile == "" {
			fmt.Println(geneartedType)
			return
		}
		err = writeToFile(&geneartedType)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

// writeToFile function creates a file if the file is not already present
// otherwise, it will truncate all the data and rewrites the file.
func writeToFile(data *string) error {
	fp := outputFile
	if outputFile == "" {
		fp = defaultOutputFile
	}
	file, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, *data)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	jsonCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Input JSON filename")
	jsonCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "File to write the genearted types to")
	jsonCmd.Flags().BoolVarP(&outputToFile, "ouput-to-file", "f", false, "Output the generated type to a file")

	// The input flag is required because the actual
	// JSON is sourced from a JSON file.
	jsonCmd.MarkPersistentFlagRequired("input")
}
