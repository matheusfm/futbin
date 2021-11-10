package cmd

import (
	"github.com/spf13/cobra"
)

// popularCmd represents the popular command
var popularCmd = &cobra.Command{
	Use:   "popular",
	Short: "Popular players",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := playersClient.Popular()
		cobra.CheckErr(err)
		printPlayers(p)
	},
}

func init() {
	rootCmd.AddCommand(popularCmd)
}
