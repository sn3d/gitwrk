package repo

import (
	"time"

	"github.com/unravela/gitwrk/worklog"
	"gopkg.in/src-d/go-git.v4"
)

// GetWorkLogFromRepo go through all commits matching time window and
// extrac work logs for each commit.
func GetWorkLogFromRepo(dir string, since time.Time, till time.Time) worklog.WorkLogs {

	// open the repository and get log iterator
	repo, _ := git.PlainOpen(dir)
	iterator, _ := repo.Log(&git.LogOptions{
		Order: git.LogOrderCommitterTime,
	})

	// iterate over all commits until
	// reach the 'since'
	output := make([]worklog.WorkLog, 0)
	for {
		c, err := iterator.Next()
		if err != nil {
			break
		}

		if c.Author.When.Before(since) {
			break
		}

		if c.Author.When.After(till) {
			continue
		}

		// commit belong to since-till range
		wlogs := worklog.Create(c.Author.Email, c.Author.When, c.Message)
		output = append(output, wlogs...)
	}

	return output
}
