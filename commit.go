package gitwrk

import (
	"regexp"
	"time"
)

// Most of WorkLog data are originated in Commit. Commit is produced by
// Repository and provide all needed parsing functionality
type Commit struct {
	// Who is author of the commit
	Author string

	// Commit message
	Message string

	// When was commit made, don't use it directly, prefer GetDate
	When time.Time
}

var (
	spentReg    *regexp.Regexp
	durationReg *regexp.Regexp
	dateReg     *regexp.Regexp
)

func init() {
	spentReg, _ = regexp.Compile("(?mi:spen[t|d]:?\\s*(\\d.*)$)")
	dateReg, _ = regexp.Compile("(?mi:date:?\\s*(\\d{4}-\\d{2}-\\d{2}))")
	durationReg, _ = regexp.Compile("(([0-9]+|m|h)+)")
}

// function returns 'spent: XXXX' in commit message as Duration
func (c Commit) Spent() []time.Duration {

	// first, we check if there is present 'Spent' with duration part.
	// If yes, we continue only with duration part
	r := spentReg.FindStringSubmatchIndex(c.Message)
	if len(r) != 4 {
		return []time.Duration{}
	}
	spent := c.Message[r[2]:r[3]]

	// the duration part might have multiple durations e.g. '1h30m, 2h50m, 14m'
	// this for loop is going through all durations
	r = durationReg.FindStringSubmatchIndex(spent)
	output := make([]time.Duration, 0)
	for len(r) > 0 {
		durationTxt := spent[r[0]:r[1]]
		duration, _ := time.ParseDuration(durationTxt)
		output = append(output, duration)

		spent = spent[r[1]:]
		r = durationReg.FindStringSubmatchIndex(spent)
	}

	return output
}

// returns you 'date: XXXX-XX-XX' in commit message as Time. If
// it's not defined or it's missing, then it's used When value
func (c Commit) Date() time.Time {
	var err error

	r := dateReg.FindStringSubmatch(c.Message)
	if len(r) != 2 {
		return c.When
	}

	dateStr := r[1]
	when, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.When
	}

	return when
}
