package cmd

import (
	"github.com/spf13/cobra"
)

// totwCmd represents the totw command
var totwCmd = &cobra.Command{
	Use:     "totw",
	Short:   "current TOTW players",
	Example: "futbin totw",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := playersClient.TOTW()
		cobra.CheckErr(err)
		printPlayers(p)
	},
}

func init() {
	rootCmd.AddCommand(totwCmd)
}
