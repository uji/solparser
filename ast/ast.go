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

// ----------------------------------------------------------------------------
// Expression Nodes

type Expression interface {
	Node
	expressionNode()
}

// ----------------------------------------------------------------------------
// Literal Nodes

type Literal interface {
	Node
	literalNode()
	expressionNode()
}

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
	List []Node // EmptyStringLiteral | NonEmptyStringLiteral
}

type HexStringLiteral []*HexString

type UnicordStringLiteral []*UnicordStrings

type NumberLiteral struct {
	Number     Node // DecimalNumber | HexNumber
	NumberUnit *NumberUnit
}

func (*BooleanLiteral) literalNode()       {}
func (*StringLiteral) literalNode()        {}
func (*HexStringLiteral) literalNode()     {}
func (*UnicordStringLiteral) literalNode() {}
func (*NumberLiteral) literalNode()        {}

func (*BooleanLiteral) expressionNode()       {}
func (*StringLiteral) expressionNode()        {}
func (*HexStringLiteral) expressionNode()     {}
func (*UnicordStringLiteral) expressionNode() {}
func (*NumberLiteral) expressionNode()        {}

// ----------------------------------------------------------------------------

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
	List []Printable
}

// ----------------------------------------------------------------------------
// Printable Nodes

type Printable interface {
	Node
	printableNode()
}

type SingleQuotedPrintable struct {
	Begin  token.Pos
	String string
}

func (s *SingleQuotedPrintable) Pos() token.Pos {
	return s.Begin
}

func (s *SingleQuotedPrintable) End() token.Pos {
	return token.Pos{
		Column: s.Begin.Column + len(s.String) + 1,
		Line:   s.Begin.Line,
	}
}

type DoubleQuotedPrintable struct {
	Begin  token.Pos
	String string
}

func (d *DoubleQuotedPrintable) Pos() token.Pos {
	return d.Begin
}

func (d *DoubleQuotedPrintable) End() token.Pos {
	return token.Pos{
		Column: d.Begin.Column + len(d.String) + 1,
		Line:   d.Begin.Line,
	}
}

type EscapeSequence struct {
	Begin  token.Pos
	String string
}

func (e *EscapeSequence) Pos() token.Pos {
	return e.Begin
}

func (e *EscapeSequence) End() token.Pos {
	return token.Pos{
		Column: e.Begin.Column + len(e.String),
		Line:   e.Begin.Line,
	}
}

func (*SingleQuotedPrintable) printableNode() {}
func (*DoubleQuotedPrintable) printableNode() {}
func (*EscapeSequence) printableNode()        {}

// ----------------------------------------------------------------------------

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
	Return     token.Pos
	Expression Expression
}

func (r *ReturnStatement) Pos() token.Pos {
	return r.Return
}
func (r *ReturnStatement) End() token.Pos {
	return token.Pos{
		Column: r.Expression.End().Column + 1,
		Line:   r.Expression.End().Line,
	}
}
