package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func newTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("   ")
	table.SetNoWhiteSpace(true)

	return table
}

func printTable(header []string, data [][]string) {
	t := newTable()
	t.SetHeader(header)
	t.AppendBulk(data)
	t.Render()
}
