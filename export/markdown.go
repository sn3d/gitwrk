package export

import (
	"github.com/unravela/gitwrk"
	"github.com/unravela/gitwrk/table"
	"io"
	"strings"
	"time"
)

// Markdown render work logs as markdown text with
// table.
func Markdown(wlogs gitwrk.WorkLogs, out io.Writer) {
	tbl := table.NewTable(
		table.Header{
			Name: "WHEN",
			Align: table.AligmentLeft,
		},
		table.Header{
			Name: "AUTHOR",
			Align: table.AligmentLeft,
		},
		table.Header{
			Name: "TYPE",
			Align: table.AligmentLeft,
		},
		table.Header{
			Name: "SCOPE",
			Align: table.AligmentLeft,
		},
		table.Header{
			Name: "DESCRIPTION",
			Align: table.AligmentLeft,
		},
		table.Header{
			Name:  "SPENT",
			Align: table.AligmentRight,
		})

	spentTotal := int64(0)
	for _, wlog := range wlogs {
		when := wlog.When.Format("2006-01-02")
		author := formatAuthor(wlog.Author)
		spent := wlog.Spent.String()
		scmType := wlog.Scm.Type
		scmScope := wlog.Scm.Scope
		descr := formatDescription(wlog.Scm)
		tbl.Append(when, author, scmType, scmScope, descr, spent)

		spentTotal += int64(wlog.Spent)
	}

	io.WriteString(out, "\n")
	tbl.Print(out)

	// write 'total'
	io.WriteString(out, "\n**Spent total**: ")
	io.WriteString(out, time.Duration(spentTotal).String())
	io.WriteString(out, "\n\n")
}

// Function format the long subject to text that can be used in table. It's returns
// you first line of SCM's subject and  trim the text to max. 25 characters.
func formatDescription(s *gitwrk.SemanticCommitMessage) string {
	descr := strings.TrimSpace(s.Subject)

	// get only first line (if it's multiline text)
	newLineIdx := strings.Index(s.Subject, "\n")
	if newLineIdx > 0 {
		descr = s.Subject[:newLineIdx]
	}

	// trim to max 50 characters
	if len(descr) > 50 {
		descr = descr[:50] + "..."
	}

	return descr
}

// Function cutout the email and keep only left part (name) of email.
func formatAuthor(author string) string {
	idx := strings.Index(author, "@")
	if idx > 0 {
		return author[:idx]
	}
	return author
}

