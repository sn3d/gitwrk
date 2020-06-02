package export

import (
	"encoding/csv"
	"github.com/unravela/gitwrk"
	"io"
	"strconv"
)

// Csv is responsible for rendering output
// when '-o csv' is set.
func Csv(wlogs gitwrk.WorkLogs, out io.Writer) {
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
