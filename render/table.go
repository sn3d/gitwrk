package render

import (
	"io"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/unravela/gitwrk/worklog"
)

// Table render the collection
// of work logs into table.
func Table(wlogs worklog.WorkLogs, out io.Writer) {
	table := tablewriter.NewWriter(out)

	// set format of table
	table.SetHeaderAlignment(tablewriter.ALIGN_RIGHT)
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)

	// header
	table.SetHeader([]string{"When", "Author", "Type", "Scope", "Duration (min)"})

	// rows
	durationTotal := 0
	for _, wlog := range wlogs {
		when := wlog.When.Format("2006-01-02 15:04:05")
		duration := int(wlog.Duration.Minutes())
		durationTotal += duration
		scmType := wlog.Scm.Type
		scmScope := wlog.Scm.Scope
		table.Append([]string{when, wlog.Author, scmType, scmScope, strconv.Itoa(duration)})
	}

	// total footer
	table.SetFooter([]string{"", "", "", "Total", strconv.Itoa(durationTotal)})
	table.Render()
}
