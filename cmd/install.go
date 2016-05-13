package cmd

import (
	"os"
	"path/filepath"

	"github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command {
	Use:   "install"
	Short: "install a EPM package and/or dependencies of a package"
	Long:  "initialize EPM package"
	Run: Initialize,
}

func buildInstallCommand() {
	addInstallFlags()
	addInstallArgs()
}

func addInitFlags() {
	InstallCmd.Flags().BoolVarP(, "yaml", "-y", true, "default value, initializes with epm.yaml file.")
}

func Initialize() {

}