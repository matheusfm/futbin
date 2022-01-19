package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/matheusfm/futbin/nations"
	"github.com/matheusfm/futbin/players"
	"github.com/spf13/cobra"
)

var (
	playersClient        = players.DefaultClient()
	flagPage             int
	flagVersion          string
	flagClubID           int
	flagNationID         int
	flagLeagueID         int
	flagPosition         []string
	flagPrice            string
	flagWeakFoot         string
	flagSkills           string
	flagRating           string
	flagPace             string
	flagAcceleration     string
	flagSprintSpeed      string
	flagShooting         string
	flagPositioning      string
	flagFinishing        string
	flagShotPower        string
	flagLongShots        string
	flagVolleys          string
	flagPenalties        string
	flagDribbling        string
	flagAgility          string
	flagBalance          string
	flagReactions        string
	flagBallControl      string
	flagComposure        string
	flagDefending        string
	flagInterceptions    string
	flagHeadingAccuracy  string
	flagMarking          string
	flagStandingTackle   string
	flagSlidingTackle    string
	flagPassing          string
	flagVision           string
	flagCrossing         string
	flagFreeKickAccuracy string
	flagShortPassing     string
	flagLongPassing      string
	flagCurve            string
	flagPhysicality      string
	flagJumping          string
	flagStamina          string
	flagStrength         string
	flagAggression       string
)

// playersCmd represents the players command
var playersCmd = &cobra.Command{
	Use:   "players",
	Short: "players",
	Example: `1.  # Brazilian players in LaLiga:
    futbin players --nation 54 --league 53
2.  # OTW players (see options in 'futbin cardversions' command):
    futbin players --version otw
3.  # Brazilian players with more than 90 of passing:
    futbin players --nation 54 --passing 90-
4.  # Icons with a maximum price of 300K:
    futbin players --league 2118 --price -300000
5.  # Players with 5 weak foot and 5 skills:
    futbin players --wf 5 --skills 5`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := playersClient.Get(&players.Options{
			Platform:         platform,
			Page:             flagPage,
			Version:          flagVersion,
			ClubID:           flagClubID,
			NationID:         flagNationID,
			LeagueID:         flagLeagueID,
			Position:         flagPosition,
			Price:            flagToRange(flagPrice),
			WeakFoot:         flagToRange(flagWeakFoot),
			Skills:           flagToRange(flagSkills),
			Rating:           flagToRange(flagRating),
			Pace:             flagToRange(flagPace),
			Acceleration:     flagToRange(flagAcceleration),
			SprintSpeed:      flagToRange(flagSprintSpeed),
			Shooting:         flagToRange(flagShooting),
			Positioning:      flagToRange(flagPositioning),
			Finishing:        flagToRange(flagFinishing),
			ShotPower:        flagToRange(flagShotPower),
			LongShots:        flagToRange(flagLongShots),
			Volleys:          flagToRange(flagVolleys),
			Penalties:        flagToRange(flagPenalties),
			Dribbling:        flagToRange(flagDribbling),
			Agility:          flagToRange(flagAgility),
			Balance:          flagToRange(flagBalance),
			Reactions:        flagToRange(flagReactions),
			BallControl:      flagToRange(flagBallControl),
			Composure:        flagToRange(flagComposure),
			Defending:        flagToRange(flagDefending),
			Interceptions:    flagToRange(flagInterceptions),
			HeadingAccuracy:  flagToRange(flagHeadingAccuracy),
			Marking:          flagToRange(flagMarking),
			StandingTackle:   flagToRange(flagStandingTackle),
			SlidingTackle:    flagToRange(flagSlidingTackle),
			Passing:          flagToRange(flagPassing),
			Vision:           flagToRange(flagVision),
			Crossing:         flagToRange(flagCrossing),
			FreeKickAccuracy: flagToRange(flagFreeKickAccuracy),
			ShortPassing:     flagToRange(flagShortPassing),
			LongPassing:      flagToRange(flagLongPassing),
			Curve:            flagToRange(flagCurve),
			Physicality:      flagToRange(flagPhysicality),
			Jumping:          flagToRange(flagJumping),
			Stamina:          flagToRange(flagStamina),
			Strength:         flagToRange(flagStrength),
			Aggression:       flagToRange(flagAggression),
		})
		cobra.CheckErr(err)
		printPlayers(p)
	},
}

func init() {
	rootCmd.AddCommand(playersCmd)
	playersCmd.PersistentFlags().IntVar(&flagPage, "page", 1, "Page")
	playersCmd.PersistentFlags().StringVar(&flagVersion, "version", "", "Card version")
	playersCmd.PersistentFlags().IntVar(&flagClubID, "club", 0, "Club ID")
	playersCmd.PersistentFlags().IntVar(&flagNationID, "nation", 0, "Nation ID")
	playersCmd.PersistentFlags().IntVar(&flagLeagueID, "league", 0, "League ID")
	playersCmd.PersistentFlags().StringVar(&flagWeakFoot, "wf", "", "Weak Foot")
	playersCmd.PersistentFlags().StringVar(&flagSkills, "skills", "", "Skills")
	playersCmd.PersistentFlags().StringVar(&flagRating, "ovr", "", "Rating")
	playersCmd.PersistentFlags().StringVar(&flagPrice, "price", "", "Price")
	playersCmd.PersistentFlags().StringSliceVar(&flagPosition, "position", []string{}, "Position")
	playersCmd.PersistentFlags().StringVar(&flagPace, "pace", "", "Pace")
	playersCmd.PersistentFlags().StringVar(&flagAcceleration, "acceleration", "", "Acceleration")
	playersCmd.PersistentFlags().StringVar(&flagSprintSpeed, "sprint-speed", "", "Sprint Speed")
	playersCmd.PersistentFlags().StringVar(&flagShooting, "shooting", "", "Shooting")
	playersCmd.PersistentFlags().StringVar(&flagPositioning, "positioning", "", "Positioning")
	playersCmd.PersistentFlags().StringVar(&flagFinishing, "finishing", "", "Finishing")
	playersCmd.PersistentFlags().StringVar(&flagShotPower, "shot-power", "", "Shot Power")
	playersCmd.PersistentFlags().StringVar(&flagLongShots, "long-shots", "", "Long Shots")
	playersCmd.PersistentFlags().StringVar(&flagVolleys, "volleys", "", "Volleys")
	playersCmd.PersistentFlags().StringVar(&flagPenalties, "penalties", "", "Penalties")
	playersCmd.PersistentFlags().StringVar(&flagDribbling, "dribbling", "", "Dribbling")
	playersCmd.PersistentFlags().StringVar(&flagAgility, "agility", "", "Agility")
	playersCmd.PersistentFlags().StringVar(&flagBalance, "balance", "", "Balance")
	playersCmd.PersistentFlags().StringVar(&flagReactions, "reactions", "", "Reactions")
	playersCmd.PersistentFlags().StringVar(&flagBallControl, "ball-control", "", "Ball Control")
	playersCmd.PersistentFlags().StringVar(&flagComposure, "composure", "", "Composure")
	playersCmd.PersistentFlags().StringVar(&flagDefending, "defending", "", "Defending")
	playersCmd.PersistentFlags().StringVar(&flagInterceptions, "interceptions", "", "Interceptions")
	playersCmd.PersistentFlags().StringVar(&flagHeadingAccuracy, "heading-accuracy", "", "Heading Accuracy")
	playersCmd.PersistentFlags().StringVar(&flagMarking, "marking", "", "Marking")
	playersCmd.PersistentFlags().StringVar(&flagStandingTackle, "standing-tackle", "", "Standing Tackle")
	playersCmd.PersistentFlags().StringVar(&flagSlidingTackle, "sliding-tackle", "", "Sliding Tackle")
	playersCmd.PersistentFlags().StringVar(&flagPassing, "passing", "", "Passing")
	playersCmd.PersistentFlags().StringVar(&flagVision, "vision", "", "Vision")
	playersCmd.PersistentFlags().StringVar(&flagCrossing, "crossing", "", "Crossing")
	playersCmd.PersistentFlags().StringVar(&flagFreeKickAccuracy, "free-kick-accuracy", "", "Free Kick Accuracy")
	playersCmd.PersistentFlags().StringVar(&flagShortPassing, "short-passing", "", "Short Passing")
	playersCmd.PersistentFlags().StringVar(&flagLongPassing, "long-passing", "", "Long Passing")
	playersCmd.PersistentFlags().StringVar(&flagCurve, "curve", "", "Curve")
	playersCmd.PersistentFlags().StringVar(&flagPhysicality, "physicality", "", "Physicality")
	playersCmd.PersistentFlags().StringVar(&flagJumping, "jumping", "", "Jumping")
	playersCmd.PersistentFlags().StringVar(&flagStamina, "stamina", "", "Stamina")
	playersCmd.PersistentFlags().StringVar(&flagStrength, "strength", "", "Strength")
	playersCmd.PersistentFlags().StringVar(&flagAggression, "aggression", "", "Aggression")
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
			color.HiYellowString(humanize.Comma(int64(player.Price(platform)))),
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

func flagToRange(s string) players.Range {
	values := strings.SplitN(s, "-", 2)
	minS := strings.ReplaceAll(values[0], ",", "") // to support comma separated prices
	min, err := strconv.Atoi(minS)
	minOK := err == nil
	if !minOK && minS != "" {
		return players.Range{} //invalid character
	}
	max := 0
	maxOK := false
	if len(values) > 1 {
		maxS := strings.ReplaceAll(values[1], ",", "")
		max, err = strconv.Atoi(maxS)
		maxOK = err == nil
	}
	r := players.Range{}
	if minOK {
		r.Min = min
	}
	if maxOK {
		r.Max = max
	} else if !strings.Contains(s, "-") {
		r.Max = min
	} else if values[1] != "" {
		return players.Range{} //invalid character
	}
	return r
}
