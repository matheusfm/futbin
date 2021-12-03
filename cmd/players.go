package cmd

import (
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
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
	Short: "players",
	Example: `1.  # Brazilian players in LaLiga:
    futbin players --nation 54 --league 53
2.  # OTW players (see options in 'futbin cardversions' command):
    futbin players --version otw`,
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

func init() {
	rootCmd.AddCommand(playersCmd)

	playersCmd.PersistentFlags().IntVar(&fPage, "page", 1, "page")
	playersCmd.PersistentFlags().StringVar(&fVersion, "version", "", "card version")
	playersCmd.PersistentFlags().IntVar(&fClubID, "club", 0, "club ID")
	playersCmd.PersistentFlags().IntVar(&fNationID, "nation", 0, "nation ID")
	playersCmd.PersistentFlags().IntVar(&fLeagueID, "league", 0, "league ID")
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
			colorStat(player.Pace),
			colorStat(player.Shooting),
			colorStat(player.Passing),
			colorStat(player.Dribbling),
			colorStat(player.Defending),
			colorStat(player.Physicality),
		}
	}
	printTable(header, data)
}

func colorStat(stat int) string {
	if stat <= 50 {
		return color.RedString("%d", stat)
	}
	if stat <= 70 {
		return color.YellowString("%d", stat)
	}
	return color.GreenString("%d", stat)
}
