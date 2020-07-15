package gitwrk

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
)

// GetWorkLogFromRepo go through all commits matching time window and
// extrac work logs for each commit.
func GetWorkLogFromRepo(dir string, since time.Time, till time.Time) (WorkLogs, error) {

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
	output := make([]WorkLog, 0)
	for {
		c, err := iterator.Next()
		if err != nil {
			break
		}

		when := c.Author.When
		if when.Before(since) {
			break
		}

		if when.After(till) {
			continue
		}

		// commit belong to since-till range
		wlogs := Create(c.Author.Email, c.Author.When, c.Message)

		// we need to filter by since&till values once more
		// because wlogs might contain older logs. It's caused
		// when since is 2020-07-04, the commit is with '2020-07-04', but
		// commit contains 'spent: 5m 6h 8h). That means 5m - 2020-07-04, 6h - 2020-07-03
		// and 8h - 2020-07-02.
		//
		// the bug: https://github.com/unravela/gitwrk/issues/1
		wlogs = wlogs.Filter(func(log WorkLog) bool {
			if log.When.Before(since) {
				return false
			}

			if log.When.After(till) {
				return false
			}
			return true
		})

		output = append(output, wlogs...)
	}

	return output, nil
}
