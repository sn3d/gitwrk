package render

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/unravela/gitwrk/worklog"
)

// Csv is responsible for rendering output
// when '-o csv' is set.
func Csv(wlogs worklog.WorkLogs, out io.Writer) {
	csvWriter := csv.NewWriter(out)

	// write header
	csvWriter.Write([]string{"when", "author", "type", "scope", "spent", "spent_minutes"})

	for _, wlog := range wlogs {
		csvWriter.Write([]string{
			wlog.When.String(),
			wlog.Author,
			wlog.Scm.Type,
			wlog.Scm.Scope,
			wlog.Spent.String(),
			strconv.Itoa(int(wlog.Spent.Minutes())),
		})
	}

	csvWriter.Flush()
}
