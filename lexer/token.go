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
	False
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
	True
	Try
	Type
	Using
	View
	Virtual
	While
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
		return False
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
		return True
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
