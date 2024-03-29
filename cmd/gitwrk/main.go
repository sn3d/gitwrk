package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/unravela/gitwrk"
	"github.com/unravela/gitwrk/export"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "gitwrk",
		Usage:   "Get work log from Git repository",
		Action:  mainCmd,
		Version: gitwrk.GetVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "git",
				Usage: "Path to git repository. Default value is current dir.",
				Value: ".",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Type of output. Can be 'json', 'csv', 'markdown' or 'table'. The default is 'markdown'",
				Value:   "markdown",
			},
			&cli.StringFlag{
				Name:  "author",
				Usage: "Filter worklogs for given author",
				Value: "",
			},
			&cli.StringFlag{
				Name:  "type",
				Usage: "Filter worklogs for given type by semantic commit message (e.g. chore, fix etc)",
				Value: "",
			},
			&cli.StringFlag{
				Name:  "scope",
				Usage: "Filter worklogs for given scope by semantic commit message",
				Value: "",
			},
			&cli.StringFlag{
				Name:  "since",
				Usage: "Lower boundary of time. Older commits will be ignored. It's in ISO form YYYY-MM-DD",
				Value: "1970-01-01",
			},
			&cli.StringFlag{
				Name:  "till",
				Usage: "Upper boundary of time. Newest commits will be ignored. It's in ISO form YYYY-MM-DD",
				Value: "2099-01-01",
			},
			&cli.BoolFlag{
				Name:  "last-month",
				Usage: "get all work logs for last month",
			},
			&cli.BoolFlag{
				Name:  "current-month",
				Usage: "get all work logs for current month",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainCmd(ctx *cli.Context) error {

	var since time.Time
	var till time.Time

	// Process options
	since, err := time.Parse("2006-01-02", ctx.String("since"))
	if err != nil {
		return err
	}

	till, err = time.Parse("2006-01-02", ctx.String("till"))
	if err != nil {
		return err
	}

	if ctx.Bool("last-month") {
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		currentLocation := now.Location()

		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		since = firstOfMonth.AddDate(0, -1, 0)
		till = since.AddDate(0, 1, 0).Add(-1)
	}

	if ctx.Bool("current-month") {
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		currentLocation := now.Location()

		since = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		till = since.AddDate(0, 1, -0).Add(-1)
	}

	// get and transform commits to worklogs
	gitDir := ctx.String("git")
	commits, err := gitwrk.GetCommits(gitDir, since, till)

	wlogs := make([]gitwrk.WorkLog, 0)
	for _, c := range commits {
		wl := gitwrk.CreateWorkLogs(c)
		wlogs = append(wlogs, wl...)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// filter worklogs
	wlogs = gitwrk.WorkLogs(wlogs).Filter(func(w gitwrk.WorkLog) bool {

		// filter by 'author' if it's set
		author := strings.ToLower(ctx.String("author"))
		if author != "" && strings.ToLower(w.Author) != author {
			return false
		}

		// filter by 'type' if it's set
		scmType := strings.ToLower(ctx.String("type"))
		if scmType != "" && strings.ToLower(w.Scm.Type) != scmType {
			return false
		}

		// filter by 'scope' if it's set
		scmScope := strings.ToLower(ctx.String("scope"))
		if scmScope != "" && strings.ToLower(w.Scm.Scope) != scmScope {
			return false
		}

		// we need to filter by since&till values once more
		// because wlogs might contain older logs. It's caused
		// when since is 2020-07-04, the commit is with '2020-07-04', but
		// commit contains 'spent: 5m 6h 8h). That means 5m - 2020-07-04, 6h - 2020-07-03
		// and 8h - 2020-07-02.
		//
		// This is best place if we want to avoid multiple loops
		//
		// the bug: https://github.com/unravela/gitwrk/issues/1
		if w.When.Before(since) {
			return false
		}
		if w.When.After(till) {
			return false
		}

		return true
	})

	// render output (by desired type)
	switch strings.ToLower(ctx.String("output")) {
	case "table":
		export.Table(wlogs, os.Stdout)
		break
	case "markdown":
		export.Markdown(wlogs, os.Stdout)
		break
	case "json":
		export.JSON(wlogs, os.Stdout)
		break
	case "csv":
		export.Csv(wlogs, os.Stdout)
		break

	}

	return nil
}
