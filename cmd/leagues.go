package cmd

import (
	"strconv"

	"github.com/matheusfm/futbin/leagueclubs"
	"github.com/spf13/cobra"
)

var leaguesAndClubsClient = leagueclubs.DefaultClient()

// leaguesCmd represents the leagues command
var leaguesCmd = &cobra.Command{
	Use:     "leagues",
	Short:   "leagues",
	Example: "futbin leagues",
	Run: func(cmd *cobra.Command, args []string) {
		leagues, err := leaguesAndClubsClient.Get()
		cobra.CheckErr(err)
		data := make([][]string, len(leagues))
		for i, league := range leagues {
			data[i] = []string{strconv.Itoa(league.ID), league.Name}
		}
		printTable([]string{"ID", "NAME"}, data)
	},
}

func init() {
	rootCmd.AddCommand(leaguesCmd)
}
