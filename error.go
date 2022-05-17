package solparser

import "github.com/uji/solparser/lexer"

type Error struct {
	Token lexer.Token
	Msg   string
}

var _ error = &Error{}

func newError(token lexer.Token, msg string) *Error {
	return &Error{
		Token: token,
		Msg:   msg,
	}
}

func (e *Error) Error() string {
	if e.Token.Pos.IsValid() {
		return e.Token.Pos.String() + ": " + e.Msg
	}
	return e.Msg
}
