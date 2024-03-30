package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	PKCS11Module string
}

var cfg Config = Config{}

var rootCmd = &cobra.Command{
	Use:   "ePorezi",
	Short: "ePorezi - Unofficial app for access Serbian electronic tax system ePorezi",
	Long: `Unofficial app for access Serbian electronic tax system ePorezi.
Allows Linux users to access Serbian tax system ePorezi.`,
	Version: "1.0.0",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfg.PKCS11Module, "module", "m", "/usr/lib/libaetpkss.so.3", "PKCS11 Module")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ExitWithError(err)
	}
}

func ExitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
