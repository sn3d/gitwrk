package gitwrk

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
)

// returns you all commits in time frame between since and till
func GetCommits(repoDir string, since time.Time, till time.Time) ([]Commit, error) {
	// open the repository and get log iterator
	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		return nil, fmt.Errorf("Cannot open the %s directory. Check if it's git repository", repoDir)
	}

	iterator, err := repo.Log(&git.LogOptions{
		Order: git.LogOrderCommitterTime,
	})

	if err != nil {
		return nil, fmt.Errorf("Error in getting data from repository. %s", err)
	}

	// iterate over all commits between 'since' and 'till'
	output := make([]Commit, 0)
	for {
		c, err := iterator.Next()
		if err != nil {
			break
		}

		when := c.Author.When
		if when.After(since) && when.Before(till) {
			commit := Commit{
				Author:  c.Author.Email,
				Message: c.Message,
				When:    when,
			}
			output = append(output, commit)
		}
	}

	return output, nil
}
