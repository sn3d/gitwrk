package gitwrk

import (
	"testing"
	"time"
)

func TestSimpleSpent(t *testing.T) {

	commit := Commit{Message: "Spent 1h40m"}

	d := commit.Spent()
	if len(d) != 1 {
		t.Error("The parse function must return only one duration")
	}

	if d[0].Minutes() != 100 {
		t.Error("The duration need to be 1h40m (100m)")
	}
}

func TestMultiSpent(t *testing.T) {
	commit := Commit{Message: "Spent 1h40m, 30m, 3h20m"}
	d := commit.Spent()

	if len(d) != 3 {
		t.Error("The parse function must return only one duration")
	}

	if d[1].Minutes() != 30 {
		t.Error("The second duration need to be 30m")
	}

	if d[2].Minutes() != ((3 * 60) + 20) {
		t.Error("The second duration need to be 3h20m")
	}
}

func TestParseAsGitTrailerLine(t *testing.T) {
	// the git trailer line is convention described here: https://git-scm.com/docs/git-interpret-trailers
	commit := Commit{Message: "spent: 1h"}
	d := commit.Spent()

	if len(d) != 1 || d[0].Hours() != 1 {
		t.Error("The tailer line cannot be parsed correctly")
	}
}

// Scenario:
//   given commit with 'date:' in message
//    when we get the commit's date
//    then the date must be value from commit message
func TestGetDate(t *testing.T) {
	commit := Commit{
		Message: "hello world \nspent: 1h\ndate:2021-03-18\n",
	}

	d := commit.Date()
	if d.Year() != 2021 || d.Month() != 3 || d.Day() != 18 {
		t.Error("The date should be 2021-03-18 and it's not")
	}
}

// Scenario:
//   given commit without 'date:' in message
//    when we get the Date
//    then the date must be commit's when date
func TestGetDefaultDate(t *testing.T) {
	timestamp := time.Now()

	commit := Commit{
		Message: "hello world \nspent: 1h",
		When:    timestamp,
	}

	d := commit.Date()
	if d != timestamp {
		t.Error("The date should be same as commit's When and it's not")
	}
}
