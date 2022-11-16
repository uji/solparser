package token

type PosError struct {
	Pos Pos
	Msg string
}

var _ error = &PosError{}

func NewPosError(pos Pos, msg string) *PosError {
	return &PosError{
		Pos: pos,
		Msg: msg,
	}
}

func (e *PosError) Error() string {
	if e.Pos.IsValid() {
		return e.Pos.String() + ": " + e.Msg
	}
	return e.Msg
}
