package solparser

import (
	"github.com/uji/solparser/lexer"
)

type Error struct {
	Pos lexer.Position
	Msg string
}

var _ error = &Error{}

func newError(pos lexer.Position, msg string) *Error {
	return &Error{
		Pos: pos,
		Msg: msg,
	}
}

func (e *Error) Error() string {
	if e.Pos.IsValid() {
		return e.Pos.String() + ": " + e.Msg
	}
	return e.Msg
}
