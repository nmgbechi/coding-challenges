package lexer

import (
	"unicode"

	"github.com/nmgbechi/coding-challenges/jsonparser/token"
)

type Lexer struct {
	input    []rune
	position int
}

func New(input string) *Lexer {
	return &Lexer{
		input: []rune(input),
	}
}

func (l *Lexer) peek() rune {
	if l.position >= len(l.input) {
		return 0 // NUL character
	}
	return l.input[l.position]
}

func (l *Lexer) advance() {
	l.position++
}

func (l *Lexer) skipWhitespace() {
	char := l.peek()
	for unicode.IsSpace(char) {
		l.advance()
		char = l.peek()
	}
}

func newToken(t token.TokenType, tVal string) token.Token {
	return token.Token{TokenType: t, Value: tVal}
}

// NextToken skips whitespaces then, decides what type of token the first non-whitespace character
// encountered is, advances the cursor and returns the character examined as a token
func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	curr := l.peek()
	var tok token.Token
	switch curr {
	case '{':
		tok = newToken(token.LeftBrace, "")
	case '}':
		tok = newToken(token.RightBrace, "")
	case '[':
		tok = newToken(token.LeftBracket, "")
	case ']':
		tok = newToken(token.RightBracket, "")
	case ',':
		tok = newToken(token.Comma, "")
	case ':':
		tok = newToken(token.Colon, "")
	case 0:
		// NUL character so, end of file
		tok = newToken(token.EOF, "")
	}

	// consume
	l.advance()
	return tok
}
