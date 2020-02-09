package render

import (
	"encoding/json"
	"io"
	"time"

	"github.com/unravela/gitwrk/worklog"
)

type jsonRecord struct {
	When     time.Time `json:"when"`
	Author   string    `json:"author"`
	ScmType  string    `json:"scm_type"`
	ScmScope string    `json:"scm_scope"`
	Duration int       `json:"duration_min"`
}

// JSON render the collection of work logs as JSON.
// For better control over JSON format and names, we're mapping WorkLog into
// jsonRecord first. Then we're marshalling jsonRecords.
func JSON(wlogs worklog.WorkLogs, out io.Writer) {

	records := make([]jsonRecord, len(wlogs))
	for i, wlog := range wlogs {
		records[i] = jsonRecord{
			When:     wlog.When,
			Author:   wlog.Author,
			ScmType:  wlog.Scm.Type,
			ScmScope: wlog.Scm.Scope,
			Duration: int(wlog.Duration.Minutes()),
		}
	}

	bytes, _ := json.MarshalIndent(records, "", "\t")
	out.Write(bytes)
}
