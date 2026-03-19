/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/z3co/prot/db/data"
	db "github.com/z3co/prot/db/gen"
	"github.com/z3co/prot/internal/logic"
)

var Instance *logic.Instance
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prot",
	Short: "Prot is a project based todolist",
	Long: `Prot is a project based todolist, that means that it is specifick to the folder or git branch you are on`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		configHome, err := os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("could not find user config dir: %s", err)
		}
		configDir := filepath.Join(configHome, "prot", "database.sqlite")
		dbConn, err := data.OpenDB(configDir)
		if err != nil {
			return fmt.Errorf("could not open database: %s", err)
		}
		Instance = &logic.Instance{
			Store: db.New(dbConn),
		}
		return nil
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prot.yaml)")
}
