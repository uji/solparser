package ast

import "github.com/uji/solparser/token"

type Node interface {
	Pos() token.Pos
	End() token.Pos
}

type PragmaValue struct {
	Version    string
	Expression string
}

type PragmaDirective struct {
	PragmaName  string
	PragmaValue PragmaValue
}

type FunctionDescriptor struct {
	Name string
}

type StateMutability struct {
	Pure bool
}

type ModifierList struct {
	StateMutability *StateMutability
}

type TypeName struct {
	ElementalyTypeName string
}

type EventParameter struct {
	TypeName *TypeName
}

type ParameterList struct {
	EventParameter *EventParameter
}

type ReturnParameters struct {
	ParameterList *ParameterList
}

type FunctionDefinition struct {
	FunctionDescriptor *FunctionDescriptor
	ModifierList       *ModifierList
	ReturnParameters   *ReturnParameters
}

type InheritanceSpecifier struct {
	IdentifierPath   string
	CallArgumentList CallArgumentList
}

type CallArgumentList struct{}

type ContractBodyElement Node

type ContractDefinition struct {
	Position              token.Pos
	Abstract              bool
	InheritanceSpecifiers []*InheritanceSpecifier
	ContractBodyElements  []*ContractBodyElement
}

// A File node represents a Solidity source file.
type SourceUnit struct {
	PragmaDirective    *PragmaDirective
	ContractDefinition *ContractDefinition
	FunctionDefinition *FunctionDefinition
}

type Literal Node // BooleanLiteral | StringLiteral | NumberLiteral | HexStringLiteral | UnicordStringLiteral

type BooleanLiteral struct {
	Token token.Token
}

func (b *BooleanLiteral) Pos() token.Pos {
	return b.Token.Pos
}

func (b *BooleanLiteral) End() token.Pos {
	return token.Pos{
		Column: b.Token.Pos.Column + len(b.Token.Text),
		Line:   b.Token.Pos.Line,
	}
}

type StringLiteral struct {
	Pos  token.Pos
	List []Node // EmptyStringLiteral | NonEmptyStringLiteral
}

type HexStringLiteral []*HexString

type UnicordStringLiteral []*UnicordStrings

type NumberLiteral struct {
	Number     Node // DecimalNumber | HexNumber
	NumberUnit *NumberUnit
}

type NumberUnit struct {
	Pos   token.Pos
	Value string
}

type EmptyStringLiteral struct {
	Pos          token.Pos
	SingleQuoted bool
}

type NonEmptyStringLiteral struct {
	Pos  token.Pos
	List []Node // SingleQuotedPrintable | DoubleQuotedPrintable | EscapeSequence
}

type SingleQuotedPrintable struct {
	Pos    token.Pos
	String string
}

type DoubleQuotedPrintable struct {
	Pos    token.Pos
	String string
}

type EscapeSequence struct {
	Pos    token.Pos
	String string
}

// unicode-string-literal (https://github.com/ethereum/solidity/blob/develop/docs/grammar/SolidityParser.g4#L407)
type UnicordStrings struct {
	Pos  token.Pos
	List []Node // string | EscapeSequence
}

type HexString struct {
	Pos          token.Pos
	SingleQuoted bool
	String       string
}

type ElementaryTypeNameKind int

const (
	ElementaryTypeName_Address ElementaryTypeNameKind = iota + 1
	ElementaryTypeName_AddressPayable
	ElementaryTypeName_Bool
	ElementaryTypeName_String
	ElementaryTypeName_Bytes
	ElementaryTypeName_SignedIntegerType
	ElementaryTypeName_UnsignedIntegerType
	ElementaryTypeName_FixedBytes
	ElementaryTypeName_Fixed
	ElementaryTypeName_Ufixed
)

type ElementaryTypeName struct {
	Pos  token.Pos
	Kind ElementaryTypeNameKind
}

type ReturnStatement struct {
	Position   token.Pos
	Expression Expression
}

func (r *ReturnStatement) Pos() token.Pos {
	return r.Position
}
func (r *ReturnStatement) End() token.Pos {
	return token.Pos{
		Column: r.Expression.End().Column + 1,
		Line:   r.Expression.End().Line,
	}
}

type Expression Node
