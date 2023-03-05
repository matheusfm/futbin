package cmd

import "github.com/spf13/cobra"

var (
	version    = "dev"
	commit     = ""
	date       = ""
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version",
		Run: func(cmd *cobra.Command, args []string) {
			t := newTable()
			t.SetHeader([]string{"Version", "Date", "Commit"})
			t.Append([]string{version, date, commit})
			t.Render()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
