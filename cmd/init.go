package cmd

import (
	"os"
	"path/filepath"

	"github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/spf13/cobra"
)

var InitCmd = &cobra.Command {
	Use:   "init"
	Short: "initialize EPM package"
	Long:  "initialize EPM package"
	Run: Initialize,
}

func buildInitCommand() {
	addInitFlags()
	FlagCheck(1, "eq", InitCmd, InitCmd.Flags())
}

func addInitFlags() {
	InitCmd.Flags().BoolVarP(, "yaml", "-y", true, "default value, initializes with epm.yaml file.")
	InitCmd.Flags().BoolVarP(, "json", "-j", false, "initializes with epm.json file.")
	InitCmd.Flags().BoolVarP(, "toml", "-t", false, "initializes with epm.toml file.")
}

func Initialize() {

}