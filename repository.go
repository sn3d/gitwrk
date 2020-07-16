package gitwrk

import (
	"fmt"
	"sort"
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

		wlogs := Create(c.Author.Email, c.Author.When, c.Message)
		output = append(output, wlogs...)
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].When.After(output[j].When)
	})

	return output, nil
}
