package cmd

import (
	"strconv"

	"github.com/matheusfm/futbin/nations"
	"github.com/spf13/cobra"
)

// nationsCmd represents the nations command
var nationsCmd = &cobra.Command{
	Use:     "nations",
	Short:   "nations",
	Example: "futbin nations",
	Run: func(cmd *cobra.Command, args []string) {
		data := make([][]string, len(nations.Nations))
		for id, name := range nations.Nations {
			data = append(data, []string{strconv.Itoa(id), name})
		}
		printTable([]string{"ID", "NATION"}, data)
	},
}

func init() {
	rootCmd.AddCommand(nationsCmd)
}
