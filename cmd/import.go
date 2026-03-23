/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var filePath string

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a list from a json file, make sure to run init first",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Instance.ImportList(cmd.Context(), filePath)
		if err != nil {
			return fmt.Errorf("error importing your list: %s", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVar(&filePath, "file", "./todos.json", "Set the filepath for the list")
}
