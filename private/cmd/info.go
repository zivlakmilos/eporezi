package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zivlakmilos/eporezi/pkg/scard"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show card inforamtions",
	Long:  "Show informations about sertificate",
	Run: func(cmd *cobra.Command, args []string) {
		card := scard.NewSCard(cfg.PKCS11Module)
		err := card.Open("")
		if err != nil {
			ExitWithError(err)
		}
		defer card.Close()

		info, err := card.Info()
		if err != nil {
			ExitWithError(err)
		}

		fmt.Printf("Label:\t\t%s\n", info.Label)
		fmt.Printf("Serial Number:\t%s\n", info.SerialNumber)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
