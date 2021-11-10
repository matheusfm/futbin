package cmd

import (
	"github.com/spf13/cobra"
)

// totwCmd represents the totw command
var totwCmd = &cobra.Command{
	Use:     "totw",
	Short:   "List the players of the current TOTW",
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
