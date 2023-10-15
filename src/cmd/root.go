// metabypass
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var version = "1.000-0 (2023.xx.yy)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "metabypass",
	Short:   "Add a short description here",
	Long:    "Add a long description here",
	Version: version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
