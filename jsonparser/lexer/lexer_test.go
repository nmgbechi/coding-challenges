package lexer

import (
	"testing"
)

func TestPeek(t *testing.T) {
	l := New("")
	if l.peek() != 0 {
		t.Error("peek doesn't return 0 for empty string")
	}

	l = New("waterbambam")
	if l.peek() != 'w' {
		t.Errorf("peek doesn't return position 0 character the first time, returned %v", l.peek())
	}

	l.advance()
	l.advance()

	if l.peek() != 't' {
		t.Errorf("peek doesn't return advanced element, return %v", l.peek())
	}
}

func TestAdvance(t *testing.T) {
	l := New("")
	l.advance()
	if l.position != 1 {
		t.Error("advance didn't move")
	}
}

func TestSkipWhitespace(t *testing.T) {
	l := New("                   abc")
	l.skipWhitespace()
	if l.peek() != 'a' {
		t.Error("skip whitespace doesn't skip until first occurring non-whitespace character")
	}
}
