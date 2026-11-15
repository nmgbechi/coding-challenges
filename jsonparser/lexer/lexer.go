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
	case '"':
		// start of multi-character string
		tok = l.parseString()
	default:

	}

	// consume
	l.advance()
	return tok
}

func (l *Lexer) parseString() token.Token {
	curr := l.peek()
	if curr != '"' {
		return newToken(token.Illegal, string(curr))
	}
	// go to the next character after the first encountered '"'
	l.advance()
	start := l.position
	for l.position < len(l.input) {
		if l.peek() == '"' {
			break
		}
		l.advance()
	}
	val := string(l.input[start:l.position])
	if l.peek() != '"' {
		return newToken(token.Illegal, val)
	}
	return newToken(token.String, val)
}
