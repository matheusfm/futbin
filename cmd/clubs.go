package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// clubsCmd represents the clubs command
var clubsCmd = &cobra.Command{
	Use:     "clubs",
	Short:   "List clubs",
	Example: "futbin clubs --league 13",
	Run: func(cmd *cobra.Command, args []string) {
		leagues, err := leaguesAndClubsClient.Get()
		cobra.CheckErr(err)
		data := make([][]string, 1)
		for _, league := range leagues {
			// 0 -> percorre e adiciona
			//
			if fLeagueID != 0 && fLeagueID != league.ID {
				continue
			}
			for _, club := range league.Clubs {
				data = append(data, []string{
					strconv.Itoa(club.ID),
					club.Name,
					fmt.Sprintf("%s (%d)", league.Name, league.ID),
					strconv.Itoa(club.PlayerCount),
					strconv.Itoa(club.SpecialCount),
					strconv.Itoa(club.GoldCount),
					strconv.Itoa(club.SilverCount),
					strconv.Itoa(club.BronzeCount),
				})
			}
		}
		printTable([]string{"ID", "NAME", "LEAGUE (ID)", "PLAYERS", "SPECIAL", "GOLD", "SILVER", "BRONZE"}, data)
	},
}

func init() {
	rootCmd.AddCommand(clubsCmd)
	clubsCmd.PersistentFlags().IntVar(&fLeagueID, "league", 0, "league ID")
}
