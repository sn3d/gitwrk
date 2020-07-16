package gitwrk

import (
	"time"
)

// WorkLog is main structure representing working log
type WorkLog struct {
	Author string
	When   time.Time
	Spent  time.Duration
	Scm    *SemanticCommitMessage
}

// WorkLogs is array of WorkLog-s
type WorkLogs []WorkLog

// Create construct the single working log or multiple
// working logs, depends on message
func Create(author string, when time.Time, msg string) WorkLogs {

	scm := ParseSemanticCommitMessage(msg)
	durations := parseSpent(msg)
	output := make([]WorkLog, 0)

	for i, d := range durations {
		// the 'when' is based on duration's index.
		// The durations are ordered like: now, now - 1 day, now - 2 days, now - 3 days
		w := when.AddDate(0, 0, (i * -1))

		log := WorkLog{
			Author: author,
			When:   w,
			Scm:    scm,
			Spent:  d.Round(time.Minute),
		}

		output = append(output, log)
	}

	return output
}

// Filter go through worklogs and return only these,
// they're matching by passed 'match' function
func (wlogs WorkLogs) Filter(match func(WorkLog) bool) WorkLogs {
	output := make([]WorkLog, 0)
	for _, wlog := range wlogs {
		if match(wlog) {
			output = append(output, wlog)
		}
	}
	return output
}
