package gitwrk

import (
	"testing"
	"time"
)

func TestCreateSimpleWorkLog(t *testing.T) {
	wlog := Create("justin.bieber", time.Now(), "chore(module): this is some modification\n Spent 1h30m")

	if len(wlog) != 1 {
		t.Error("There is only one duration and wlog must be only one")
	}

	if wlog[0].Scm.Type != "chore" {
		t.Error("The commit must be 'ScmTypeChore'")
	}

	if wlog[0].Spent.Minutes() != 90 {
		t.Error("The duration must be 90 minutes (1h30m)")
	}
}

func TestCreateMultipleWorkLogs(t *testing.T) {
	wlog := Create("justin.bieber", time.Unix(1578614400, 0), "chore(module): this is some modification\n Spent 1h30m 45m 2h30m 55m")

	if len(wlog) != 4 {
		t.Error("Must be 4 logs for 4 durations in commit message")
	}

	if wlog[0].When.Day() != 10 {
		t.Error("The first record belong to 10.January 2020")
	}

	if wlog[1].When.Day() != 9 {
		t.Error("The second record belong to 10.January 2020 - 1 day = 9.Jan.2020")
	}

	if wlog[3].When.Day() != 7 {
		t.Error("The fourth record belong to 10.January 2020 - 3 days = 7.Jan.2020")
	}
}
