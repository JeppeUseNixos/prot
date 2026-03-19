/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/z3co/prot/cmd"
	"github.com/z3co/prot/db/data"
)

func main() {
	configHome, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Could not find config dir: %s", err)
	}
	configDir := filepath.Join(configHome, "prot")
	err = os.MkdirAll(configDir, 0775)
	if err != nil {
		log.Fatalf("Could not create config dir: %s", err)
	}
	err = data.InitDB(filepath.Join(configDir, "database.sqlite"))
	if err != nil {
		log.Fatalf("Could not create db: %s", err)
	}
	cmd.Execute()
}
