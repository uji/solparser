package lexer

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
	Constract
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
	case "constract":
		return Constract
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
		TokenType: asKeyword(ch),
		Text:      ch,
		Pos:       pos,
	}
}

type Position struct {
	Start int
	Size  int
	Line  int
}
