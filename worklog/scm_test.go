package worklog

import "testing"

func TestMessageWithScope(t *testing.T) {
	msg := "chore(server): add support for request"
	semanticMsg := parseSemanticCommitMessage(msg)

	if semanticMsg.Type != "chore" {
		t.Error("The type must be 'chore'")
	}

	if semanticMsg.Scope != "server" {
		t.Error("The scope must be 'server'")
	}

	if semanticMsg.Subject != " add support for request" {
		t.Error("The subject doesn't match")
	}
}

func TestMessageWithoutScope(t *testing.T) {
	msg := "fix: add support for request"
	semanticMsg := parseSemanticCommitMessage(msg)

	if semanticMsg.Type != "fix" {
		t.Error("The type must be 'fix'")
	}

	if semanticMsg.Scope != "" {
		t.Error("The scope must be empty string")
	}
}

func TestNonSemanticMessage(t *testing.T) {
	semanticMsg := parseSemanticCommitMessage("this is some commit message")
	if semanticMsg.Type != "none" {
		t.Error("The type must be 'none'")
	}
}
