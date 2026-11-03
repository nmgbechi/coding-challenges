package token

type TokenType int

const (
	// Structural Symbols
	LeftBrace TokenType = iota
	RightBrace
	LeftBracket
	RightBracket
	Comma
	Colon

	// data types
	String
	Number

	// Keywords
	True
	False
	Null

	// housekeeping
	EOF
	Illegal
)

type Token struct {
	Value     string
	TokenType TokenType
}
