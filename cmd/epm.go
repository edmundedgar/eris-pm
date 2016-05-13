package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eris-ltd/eris-pm/definitions"
	"github.com/eris-ltd/eris-pm/packages"
	"github.com/eris-ltd/eris-pm/util"
	"github.com/eris-ltd/eris-pm/version"

	log "github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	. "github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
	logger "github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
	cfg "github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/eris-ltd/tendermint/config"
	"github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/spf13/cobra"
)

const VERSION = version.VERSION

// Global Do struct
var do *definitions.Do

// Defining the root command
var EPMCmd = &cobra.Command{
	Use:   "epm",
	Short: "The Eris Package Manager Deploys and Tests Smart Contract Systems",
	Long: `The Eris Package Manager Deploys and Tests Smart Contract Systems

Made with <3 by Eris Industries.

Complete documentation is available at https://docs.erisindustries.com
` + "\nVersion:\n  " + VERSION,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// TODO: make this better.... need proper epm config
		// need to be able to have variable writers (eventually)
		log.SetFormatter(logger.ErisFormatter{})

		log.SetLevel(log.WarnLevel)
		if do.Verbose {
			log.SetLevel(log.InfoLevel)
		} else if do.Debug {
			log.SetLevel(log.DebugLevel)
		}

		// clears epm.log file
		util.ClearJobResults()

		// Welcomer....
		log.Info("Hello! I'm EPM.")

		// Fixes path issues and controls for mint-client / eris-keys assumptions
		util.BundleHttpPathCorrect(do)
		util.PrintPathPackage(do)

		// Populates chainID from the chain (if its not passed)
		IfExit(util.GetChainID(do))

		// Populates the tendermint config object for proper websocket connection
		config.Set("chain_id", do.ChainID)
		config.Set("log_level", "error")
		cfg.ApplyConfig(config)
	},

		PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Ensure that errors get written to screen and generally flush the log
		// log.Flush()
	},
}

func Execute() {
	InitEPM()
	AddGlobalFlags()
	AddComands()
	EPMCmd.Execute()
}

func AddComands() {
	buildInitCommand()
	EPMCmd.AddComand(Init)
	buildInstallCommand()
	EPMCmd.AddComand(Install)
	EPMCmd.AddComand(Run)
}

//restrict flag behaviour when needed (rare but used sometimes)
func FlagCheck(num int, comp string, cmd *cobra.Command, flags []string) error {
	switch comp {
	case "eq":
		if len(flags) != num {
			cmd.Help()
			return fmt.Errorf("\n**Note** you sent our marmots the wrong number of flags.\nPlease send the marmots %d flags only.", num)
		}
	case "ge":
		if len(flags) < num {
			cmd.Help()
			return fmt.Errorf("\n**Note** you sent our marmots the wrong number of flags.\nPlease send the marmots at least %d flag(s).", num)
		}
	}
	return nil
}