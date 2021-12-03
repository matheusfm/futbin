package cmd

import (
	"github.com/spf13/cobra"
)

// latestCmd represents the latest command
var latestCmd = &cobra.Command{
	Use:     "latest",
	Short:   "latest added players",
	Example: "futbin latest",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := playersClient.Latest()
		cobra.CheckErr(err)
		printPlayers(p)
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)
}
