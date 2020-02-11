package repo

import (
	"fmt"
	"time"

	"github.com/unravela/gitwrk/worklog"
	"gopkg.in/src-d/go-git.v4"
)

// GetWorkLogFromRepo go through all commits matching time window and
// extrac work logs for each commit.
func GetWorkLogFromRepo(dir string, since time.Time, till time.Time) (worklog.WorkLogs, error) {

	// open the repository and get log iterator
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, fmt.Errorf("Cannot open the %s directory. Check if it's git repository", dir)
	}

	iterator, err := repo.Log(&git.LogOptions{
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, fmt.Errorf("Error in getting data from repository. %s", err)
	}

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

	return output, nil
}
