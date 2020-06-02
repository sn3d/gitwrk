package gitwrk

import (
	"regexp"
	"strings"
)

// SemanticCommitMessage is main structure that represent: 'type(scope): subject'
type SemanticCommitMessage struct {
	Type    string
	Scope   string
	Subject string
}

var (
	semanticReg *regexp.Regexp
)

func init() {
	// 1st group is type, 3th group is scope, 4th group is subject
	semanticReg, _ = regexp.Compile("^([a-zA-Z]+)(\\((.*)\\))?:(.*)$")
}

// ParseSemanticCommitMessage consume plain text of the commit message and
// produce the CommitMessage structure with type
// scope and subject.
//
// If commit message doesn't match to semantic commit message,
// then it's type NONE where subject is first line of messsage
func parseSemanticCommitMessage(t string) *SemanticCommitMessage {

	// get the first line
	idx := strings.Index(t, "\n")
	var line string
	if idx > 0 {
		line = t[0:idx]
	} else {
		line = t
	}

	// check if matching to semantic commit message
	res := semanticReg.FindStringSubmatchIndex(line)
	if len(res) == 10 { // it's 10 because 4 groups and main (4 + 1)*2 = 10
		scmType := line[res[2]:res[3]]

		scope := ""
		if res[6] >= 0 {
			scope = line[res[6]:res[7]]
		}

		subject := ""
		if res[8] >= 0 {
			subject = line[res[8]:res[9]]
		}

		return &SemanticCommitMessage{
			Type:    scmType,
			Scope:   scope,
			Subject: subject,
		}
	}

	// it's not a semantic commit message
	return &SemanticCommitMessage{
		Type:    "none",
		Scope:   "",
		Subject: t,
	}
}
