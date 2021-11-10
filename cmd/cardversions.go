package cmd

import (
	"github.com/matheusfm/futbin/cardversions"
	"github.com/spf13/cobra"
)

var cardversionsClient = cardversions.DefaultClient()

// cardversionsCmd represents the cardversions command
var cardversionsCmd = &cobra.Command{
	Use:     "cardversions",
	Short:   "Card versions",
	Example: "futbin cardversions",
	Run: func(cmd *cobra.Command, args []string) {
		cv, err := cardversionsClient.Get()
		cobra.CheckErr(err)
		data := make([][]string, len(cv))
		for i, version := range cv {
			data[i] = []string{version.ID, version.Name}
		}
		printTable([]string{"ID", "NAME"}, data)
	},
}

func init() {
	rootCmd.AddCommand(cardversionsCmd)
}
