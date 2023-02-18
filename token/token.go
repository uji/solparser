package token

import (
	"fmt"

	"github.com/SteelSeries/bufrr"
)

type TokenType int

const (
	Invalid TokenType = iota
	Unknown

	EOS // End of source indicator.

	LParen      // (
	RParen      // )
	LBrack      // [
	RBrack      // ]
	LBrace      // {
	RBrace      // }
	Colon       // :
	Semicolon   // ;
	Period      // .
	Conditional // ?
	DoubleArrow // =>
	RightArrow  // ->

	Assign       // =
	AssignBitOr  // |=
	AssignBitXor // ^=
	AssignBitAnd // &=
	AssignShl    // <<=
	AssignSar    // >>=
	AssignShr    // >>>=
	AssignAdd    // +=
	AssignSub    // -=
	AssignMul    // *=
	AssignDiv    // /=
	AssignMod    // %=

	Comma  // ,
	Or     // ||
	And    // &&
	BitOr  // |
	BitXor // ^
	BitAnd // &
	Shl    // <<
	Sar    // >>
	Shr    // >>>
	Add    // +
	Sub    // -
	Mul    // *
	Div    // /
	Mod    // %
	Exp    // **

	Equal              // =
	NotEqual           // !=
	LessThan           // <
	GreaterThan        // >
	LessThanOrEqual    // <=
	GreaterThanOrEqual // >=
	Not                // !
	BitNot             // ~
	Inc                // ++
	Dec                // --
	DoubleQuote        // "
	SingleQuote        // ' or \'

	// Reserved Keyword
	After
	Alias
	Apply
	Auto
	Byte
	Case
	Copyof
	Default
	Define
	Final
	Implements
	In
	Inline
	Let
	Macro
	Match
	Mutable
	Null
	Of
	Partial
	Promise
	Reference
	Relocatable
	Sealed
	Sizeof
	Static
	Supports
	Switch
	Typedef
	Typeof
	Var

	// Keyword
	Abstract
	Address
	Anonymous
	As
	Assembly
	Bool
	Break
	Bytes
	Calldata
	Catch
	Constant
	Constructor
	Continue
	Contract
	Delete
	Do
	Else
	Emit
	Enum
	Error
	Event
	External
	Fallback
	Fixed
	For
	From
	Function
	Global
	Hex
	If
	Immutable
	Import
	Indexed
	Interface
	Internal
	Is
	Library
	Mapping
	Memory
	Modifier
	NewKeyword
	Override
	Payable
	Pragma
	Private
	Public
	Pure
	Receive
	Return
	Returns
	Revert
	Storage
	String
	Struct
	Try
	Type
	Using
	View
	Virtual
	While

	// Literal
	TrueLiteral
	FalseLiteral
	Number
	StringLiteral
	UnicodeStringLiteral
	HexStringLiteral
	CommentLiteral
)

var EOSString string = string([]rune{bufrr.EOF})

func asKeyword(str string) TokenType {
	switch str {
	case EOSString:
		return EOS
	case "(":
		return LParen
	case ")":
		return RParen
	case "[":
		return LBrack
	case "]":
		return RBrack
	case "{":
		return LBrace
	case "}":
		return RBrace
	case ":":
		return Colon
	case ";":
		return Semicolon
	case ".":
		return Period
	case "?":
		return Conditional
	case "=>":
		return DoubleArrow
	case "->":
		return RightArrow
	case "=":
		return Assign
	case "|=":
		return AssignBitOr
	case "^=":
		return AssignBitXor
	case "&=":
		return AssignBitAnd
	case "<<=":
		return AssignShl
	case ">>=":
		return AssignSar
	case ">>>=":
		return AssignShr
	case "+=":
		return AssignAdd
	case "-=":
		return AssignSub
	case "*=":
		return AssignMul
	case "/=":
		return AssignDiv
	case "%=":
		return AssignMod
	case ",":
		return Comma
	case "||":
		return Or
	case "&&":
		return And
	case "|":
		return BitOr
	case "^":
		return BitXor
	case "&":
		return BitAnd
	case "<<":
		return Shl
	case ">>":
		return Sar
	case ">>>":
		return Shr
	case "+":
		return Add
	case "-":
		return Sub
	case "*":
		return Mul
	case "/":
		return Div
	case "%":
		return Mod
	case "**":
		return Exp
	case "==":
		return Equal
	case "!=":
		return NotEqual
	case "<":
		return LessThan
	case ">":
		return GreaterThan
	case "<=":
		return LessThanOrEqual
	case ">=":
		return GreaterThanOrEqual
	case "!":
		return Not
	case "~":
		return BitNot
	case "++":
		return Inc
	case "--":
		return Dec
	case `"`:
		return DoubleQuote
	case `\'`:
		return SingleQuote
	case "after":
		return After
	case "alias":
		return Alias
	case "apply":
		return Apply
	case "auto":
		return Auto
	case "byte":
		return Byte
	case "case":
		return Case
	case "copyof":
		return Copyof
	case "default":
		return Default
	case "define":
		return Define
	case "final":
		return Final
	case "implements":
		return Implements
	case "in":
		return In
	case "inline":
		return Inline
	case "let":
		return Let
	case "macro":
		return Macro
	case "match":
		return Match
	case "mutable":
		return Mutable
	case "null":
		return Null
	case "of":
		return Of
	case "partial":
		return Partial
	case "promise":
		return Promise
	case "reference":
		return Reference
	case "relocatable":
		return Relocatable
	case "sealed":
		return Sealed
	case "sizeof":
		return Sizeof
	case "static":
		return Static
	case "supports":
		return Supports
	case "switch":
		return Switch
	case "typedef":
		return Typedef
	case "typeof":
		return Typeof
	case "var":
		return Var
	case "abstract":
		return Abstract
	case "address":
		return Address
	case "anonymous":
		return Anonymous
	case "as":
		return As
	case "assembly":
		return Assembly
	case "bool":
		return Bool
	case "break":
		return Break
	case "bytes":
		return Bytes
	case "calldata":
		return Calldata
	case "catch":
		return Catch
	case "constant":
		return Constant
	case "constructor":
		return Constructor
	case "continue":
		return Continue
	case "contract":
		return Contract
	case "delete":
		return Delete
	case "do":
		return Do
	case "else":
		return Else
	case "emit":
		return Emit
	case "enum":
		return Enum
	case "error":
		return Error
	case "event":
		return Event
	case "external":
		return External
	case "fallback":
		return Fallback
	case "false":
		return FalseLiteral
	case "fixed":
		return Fixed
	case "for":
		return For
	case "from":
		return From
	case "function":
		return Function
	case "global":
		return Global
	case "hex":
		return Hex
	case "if":
		return If
	case "immutable":
		return Immutable
	case "import":
		return Import
	case "indexed":
		return Indexed
	case "interface":
		return Interface
	case "internal":
		return Internal
	case "is":
		return Is
	case "library":
		return Library
	case "mapping":
		return Mapping
	case "memory":
		return Memory
	case "modifier":
		return Modifier
	case "newKeyword":
		return NewKeyword
	case "override":
		return Override
	case "payable":
		return Payable
	case "pragma":
		return Pragma
	case "private":
		return Private
	case "public":
		return Public
	case "pure":
		return Pure
	case "receive":
		return Receive
	case "return":
		return Return
	case "returns":
		return Returns
	case "revert":
		return Revert
	case "storage":
		return Storage
	case "string":
		return String
	case "struct":
		return Struct
	case "true":
		return TrueLiteral
	case "try":
		return Try
	case "type":
		return Type
	case "using":
		return Using
	case "view":
		return View
	case "virtual":
		return Virtual
	case "while":
		return While
	}

	return Unknown
}

func asToken(str string) TokenType {
	return asKeyword(str)
}

type Token struct {
	TokenType TokenType
	Text      string
	Pos       Pos
}

func NewToken(ch string, pos Pos) Token {
	return Token{
		TokenType: asToken(ch),
		Text:      ch,
		Pos:       pos,
	}
}

type Pos struct {
	Column int
	Line   int
}

func (p Pos) IsValid() bool { return p.Line > 0 && p.Column > 0 }

func (p Pos) String() string {
	if !p.IsValid() {
		return "-"
	}

	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

// IsSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func IsSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\n', '\t', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
