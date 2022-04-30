package lexer

type TokenType int

const (
	// Misc characters
	Invalid TokenType = iota
	Hat
	Tilde
	TGREATER
	TLESS
	Equal
	Colon
	ParenL
	ParenR
	BraceL
	BraceR
	SEMICOLON

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
		return TLESS
	case '>':
		return TGREATER
	case '=':
		return Equal
	case ':':
		return Colon
	case ';':
		return SEMICOLON
	case '(':
		return ParenL
	case ')':
		return ParenR
	case '{':
		return BraceL
	case '}':
		return BraceR
	}

	return Invalid
}

func asKeyword(str string) TokenType {
	switch str {
	case ">":
		return TGREATER
	case "=":
		return Equal
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

	return Invalid
}

type Token struct {
	TokenType TokenType
	Text      string
}

func NewToken(ch string) Token {
	return Token{
		TokenType: asKeyword(ch),
		Text:      ch,
	}
}
