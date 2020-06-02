package gitwrk

import (
	"regexp"
	"time"
)

var (
	spentReg    *regexp.Regexp
	durationReg *regexp.Regexp
)

func init() {
	spentReg, _ = regexp.Compile("(?mi:spen[t|d][ |\\t]+(\\d.*)$)")
	durationReg, _ = regexp.Compile("(([0-9]+|m|h)+)")
}

// ParseSpent function consume commit message and
// parse 'Spent XXXXX' text into Durations
func parseSpent(t string) []time.Duration {

	// first, we check if there is present 'Spent' with duration part.
	// If yes, we continue only with duration part
	r := spentReg.FindStringSubmatchIndex(t)
	if len(r) != 4 {
		return []time.Duration{}
	}
	d := t[r[2]:r[3]]

	// the duration part might have multiple durations e.g. '1h30m, 2h50m, 14m'
	// this for loop is going through all durations
	r = durationReg.FindStringSubmatchIndex(d)
	output := make([]time.Duration, 0)
	for len(r) > 0 {
		durationTxt := d[r[0]:r[1]]
		duration, _ := time.ParseDuration(durationTxt)
		output = append(output, duration)

		d = d[r[1]:]
		r = durationReg.FindStringSubmatchIndex(d)
	}

	return output
}
