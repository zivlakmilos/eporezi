package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zivlakmilos/eporezi/private/gui"
)

var guiCmd = &cobra.Command{
	Use:   "gui [url]",
	Short: "Start GUI",
	Long:  "Start ePorezi GUI",
	Run: func(cmd *cobra.Command, args []string) {
		url := ""
		if len(args) > 0 {
			url = args[0]
		}
		gui.Run(url)
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
