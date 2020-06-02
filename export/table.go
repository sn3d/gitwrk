package export

import (
	"github.com/unravela/gitwrk"
	"io"
	"time"

	"github.com/olekukonko/tablewriter"
)

// Table render the collection
// of work logs into table.
func Table(wlogs gitwrk.WorkLogs, out io.Writer) {
	table := tablewriter.NewWriter(out)

	// set format of table
	table.SetHeaderAlignment(tablewriter.ALIGN_RIGHT)
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)

	// header
	table.SetHeader([]string{"When", "Author", "Type", "Scope", "Spent"})

	// rows
	spentTotal := int64(0)
	for _, wlog := range wlogs {
		when := wlog.When.Format("2006-01-02 15:04:05")
		spent := wlog.Spent.String()
		spentTotal += int64(wlog.Spent)
		scmType := wlog.Scm.Type
		scmScope := wlog.Scm.Scope
		table.Append([]string{when, wlog.Author, scmType, scmScope, spent})
	}

	// total footer
	table.SetFooter([]string{"", "", "", "Total", time.Duration(spentTotal).String()})
	table.Render()
}
