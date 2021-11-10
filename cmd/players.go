package cmd

import (
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/matheusfm/futbin/nations"
	"github.com/matheusfm/futbin/players"
	"github.com/spf13/cobra"
)

var (
	playersClient = players.DefaultClient()
	fPage         int
	fVersion      string
	fClubID       int
	fNationID     int
	fLeagueID     int
)

// playersCmd represents the players command
var playersCmd = &cobra.Command{
	Use:   "players",
	Short: "List players",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := playersClient.Get(&players.Options{
			Platform: platform,
			Page:     fPage,
			Version:  fVersion,
			ClubID:   fClubID,
			NationID: fNationID,
			LeagueID: fLeagueID,
		})
		cobra.CheckErr(err)
		printPlayers(p)
	},
}

func printPlayers(p []players.Player) {
	header := []string{
		"ID",
		"NAME",
		"RAT",
		"CLUB (ID)",
		"NATION (ID)",
		"LEAGUE (ID)",
		"POS",
		fmt.Sprintf("PRICE (%s)", platform),
		"PAC",
		"SHO",
		"PAS",
		"DRI",
		"DEF",
		"PHY",
	}
	data := make([][]string, len(p))
	for i, player := range p {
		data[i] = []string{
			strconv.Itoa(player.ID),
			player.CommonName,
			strconv.Itoa(player.Rating),
			fmt.Sprintf("%s (%d)", player.ClubName, player.ClubID),
			fmt.Sprintf("%s (%d)", nations.Nations[player.NationID], player.NationID),
			fmt.Sprintf("%s (%d)", player.League, player.LeagueID),
			player.Position,
			humanize.Comma(int64(player.Price(platform))),
			strconv.Itoa(player.Pace),
			strconv.Itoa(player.Shooting),
			strconv.Itoa(player.Passing),
			strconv.Itoa(player.Dribbling),
			strconv.Itoa(player.Defending),
			strconv.Itoa(player.Physicality),
		}
	}
	printTable(header, data)
}

func init() {
	rootCmd.AddCommand(playersCmd)

	playersCmd.PersistentFlags().IntVar(&fPage, "page", 1, "page")
	playersCmd.PersistentFlags().StringVar(&fVersion, "version", "", "card version")
	playersCmd.PersistentFlags().IntVar(&fClubID, "club", 0, "club ID")
	playersCmd.PersistentFlags().IntVar(&fNationID, "nation", 0, "nation ID")
	playersCmd.PersistentFlags().IntVar(&fLeagueID, "league", 0, "league ID")
}
