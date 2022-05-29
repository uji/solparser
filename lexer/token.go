package lexer

import "fmt"

type TokenType int

const (
	// Misc characters
	Invalid TokenType = iota
	Unknown
	Hat
	Tilde
	Greater
	Less
	Equal
	Colon
	ParenL
	ParenR
	BraceL
	BraceR
	Semicolon

	// Keyword
	Pragma
	Contract
	Function
	Pubilc
	Pure
	Returns
	Return
)

func asMiscToken(r rune) TokenType {
	switch r {
	case '^':
		return Hat
	case '~':
		return Tilde
	case '<':
		return Less
	case '>':
		return Greater
	case '=':
		return Equal
	case ':':
		return Colon
	case ';':
		return Semicolon
	case '(':
		return ParenL
	case ')':
		return ParenR
	case '{':
		return BraceL
	case '}':
		return BraceR
	}

	return Unknown
}

func asKeyword(str string) TokenType {
	switch str {
	case "pragma":
		return Pragma
	case "contract":
		return Contract
	case "function":
		return Function
	case "pubilc":
		return Pubilc
	case "pure":
		return Pure
	case "returns":
		return Returns
	case "return":
		return Return
	}

	return Unknown
}

func asToken(str string) TokenType {
	r := []rune(str)
	if len(r) != 1 {
		return asKeyword(str)
	}
	return asMiscToken(r[0])
}

type Token struct {
	TokenType TokenType
	Text      string
	Pos       Position
}

func NewToken(ch string, pos Position) Token {
	return Token{
		TokenType: asToken(ch),
		Text:      ch,
		Pos:       pos,
	}
}

type Position struct {
	Column int
	Line   int
}

func (pos Position) IsValid() bool { return pos.Line > 0 && pos.Column > 0 }

func (pos Position) String() string {
	if !pos.IsValid() {
		return "-"
	}

	return fmt.Sprintf("%d:%d", pos.Line, pos.Column)
}
