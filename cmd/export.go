/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports the current list to json format",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := Instance.ExportList(cmd.Context())
		if err != nil {
			return fmt.Errorf("error exporting your list: %s", err)
		}
		fmt.Println("Your list was exported to todos.json")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
