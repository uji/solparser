package solparser

import (
	"github.com/uji/solparser/token"
)

type Error struct {
	Pos token.Position
	Msg string
}

var _ error = &Error{}

func newError(pos token.Position, msg string) *Error {
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
