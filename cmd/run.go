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

	. "github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
	logger "github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
	cfg "github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/eris-ltd/tendermint/config"
	"github.com/eris-ltd/eris-pm/Godeps/_workspace/src/github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{

	Use: "run",
	Sort: "run epm jobs file"
	Run: RunPackage,


}

func InitEPM() {
	do = definitions.NowDo()
}

// Flags that are to be used by commands are handled by the Do struct
// Define the persistent commands (globals)
func AddGlobalFlags() {
	RunCmd.PersistentFlags().StringVarP(&do.JobFilePath, "file", "f", defaultFile(), "path to package file which EPM should use; default respects $EPM_FILE")
	RunCmd.PersistentFlags().StringVarP(&do.ContractsPath, "contracts-path", "p", defaultContracts(), "path to the contracts EPM should use; default respects $EPM_CONTRACTS_PATH")
	RunCmd.PersistentFlags().StringVarP(&do.ABIPath, "abi-path", "a", defaultAbi(), "path to the abi directory EPM should use when saving ABIs after the compile process; default respects $EPM_ABI_PATH")
	RunCmd.PersistentFlags().StringVarP(&do.Chain, "chain", "c", defaultChain(), "<ip:port> of chain which EPM should use; default respects $EPM_CHAIN_ADDR")
	RunCmd.PersistentFlags().StringVarP(&do.Signer, "sign", "s", defaultSigner(), "<ip:port> of signer daemon which EPM should use; default respects $EPM_SIGNER_ADDR")
	RunCmd.PersistentFlags().StringVarP(&do.DefaultGas, "gas", "g", defaultGas(), "default gas to use; can be overridden for any single job; default respects $EPM_GAS")
	RunCmd.PersistentFlags().StringVarP(&do.Compiler, "compiler", "m", defaultCompiler(), "<ip:port> of compiler which EPM should use; default respects $EPM_COMPILER_ADDR")
	RunCmd.PersistentFlags().StringVarP(&do.DefaultAddr, "address", "r", defaultAddr(), "default address to use; operates the same way as the [account] job, only before the epm file is ran; default respects $EPM_ADDRESS")
	RunCmd.PersistentFlags().StringSliceVarP(&do.DefaultSets, "set", "e", defaultSets(), "default sets to use; operates the same way as the [set] jobs, only before the epm file is ran (and after default address; default respects $EPM_SETS")
	RunCmd.PersistentFlags().StringVarP(&do.DefaultFee, "fee", "n", defaultFee(), "default fee to use; default respects $EPM_FEE")
	RunCmd.PersistentFlags().StringVarP(&do.DefaultAmount, "amount", "u", defaultAmount(), "default amount to use; default respects $EPM_AMOUNT")
	RunCmd.PersistentFlags().StringVarP(&do.DefaultOutput, "output", "o", defaultOutput(), "output format which epm should use [csv,json]; default respects $EPM_OUTPUT_FORMAT")
	RunCmd.PersistentFlags().BoolVarP(&do.Overwrite, "overwrite", "w", defaultOverwrite(), "overwrite jobs of similar names; defaults respects $EPM_OVERWRITE_APPROVE")
	RunCmd.PersistentFlags().BoolVarP(&do.SummaryTable, "summary", "t", defaultSummaryTable(), "output a table summarizing epm jobs; default respects $EPM_SUMMARY_TABLE")
	RunCmd.PersistentFlags().BoolVarP(&do.Verbose, "verbose", "v", defaultVerbose(), "verbose output; more output than no output flags; less output than debug level; default respects $EPM_VERBOSE")
	RunCmd.PersistentFlags().BoolVarP(&do.Debug, "debug", "d", defaultDebug(), "debug level output; the most output available for epm; if it is too chatty use verbose flag; default respects $EPM_DEBUG")
}

//----------------------------------------------------
func RunPackage(cmd *cobra.Command, args []string) {
	IfExit(packages.RunPackage(do))
}

// ---------------------------------------------------
// Defaults

func defaultFile() string {
	return setDefaultString("EPM_FILE", "./epm.yaml")
}

func defaultContracts() string {
	return setDefaultString("EPM_CONTRACTS_PATH", "./contracts")
}

func defaultAbi() string {
	return setDefaultString("EPM_ABI_PATH", "./abi")
}

func defaultChain() string {
	return setDefaultString("EPM_CHAIN_ADDR", "localhost:46657")
}

func defaultSigner() string {
	return setDefaultString("EPM_SIGNER_ADDR", "localhost:4767")
}

func defaultCompiler() string {
	verSplit := strings.Split(version.VERSION, ".")
	maj, _ := strconv.Atoi(verSplit[0])
	min, _ := strconv.Atoi(verSplit[1])
	pat, _ := strconv.Atoi(verSplit[2])
	return setDefaultString("EPM_COMPILER_ADDR", fmt.Sprintf("https://compilers.eris.industries:1%01d%02d%01d", maj, min, pat))
}

func defaultAddr() string {
	return setDefaultString("EPM_ADDRESS", "")
}

func defaultFee() string {
	return setDefaultString("EPM_FEE", "1234")
}

func defaultAmount() string {
	return setDefaultString("EPM_AMOUNT", "9999")
}

func defaultSets() []string {
	return setDefaultStringSlice("EPM_SETS", []string{})
}

func defaultGas() string {
	return setDefaultString("EPM_GAS", "1111111111")
}

func defaultOutput() string {
	return setDefaultString("EPM_OUTPUT_FORMAT", "csv")
}

func defaultSummaryTable() bool {
	return setDefaultBool("EPM_SUMMARY_TABLE", true)
}

func defaultVerbose() bool {
	return setDefaultBool("EPM_VERBOSE", false)
}

func defaultOverwrite() bool {
	return setDefaultBool("EPM_OVERWRITE_APPROVE", false)
}

func defaultDebug() bool {
	return setDefaultBool("EPM_DEBUG", false)
}

func setDefaultBool(envVar string, def bool) bool {
	env := os.Getenv(envVar)
	if env != "" {
		i, _ := strconv.ParseBool(env)
		return i
	}
	return def
}

func setDefaultString(envVar, def string) string {
	env := os.Getenv(envVar)
	if env != "" {
		return env
	}
	return def
}

func setDefaultStringSlice(envVar string, def []string) []string {
	env := os.Getenv(envVar)
	if env != "" {
		return strings.Split(env, ";")
	}
	return def
}