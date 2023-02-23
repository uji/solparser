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

type ModifierList struct {
	StateMutability *StateMutability
}

type Visibility = token.Token

// Parameter is type of ParameterList elements
type Parameter struct {
	TypeName TypeName
}

type ParameterList []*Parameter

type StateMutability = token.Token

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
// TypeName Nodes

type TypeName interface {
	Node
	typeNameNode()
}

type ElementaryTypeName []token.Token

func (e ElementaryTypeName) Pos() token.Pos {
	return e[0].Pos
}

func (e ElementaryTypeName) End() token.Pos {
	// TODO
	return token.Pos{
		Column: 1,
		Line:   1,
	}
}

func (e ElementaryTypeName) typeNameNode() {}

// ----------------------------------------------------------------------------
// Expression Nodes

type Expression interface {
	Node
	expressionNode()
}

// ----------------------------------------------------------------------------

type Identifier = token.Token

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

// EmptyStringLiteral | NonEmptyStringLiteral
type StringLiteral struct {
	From, To token.Pos
	Value    string
}

func (s StringLiteral) Pos() token.Pos {
	return s.From
}

func (s StringLiteral) End() token.Pos {
	return s.To
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

// ----------------------------------------------------------------------------

type Block struct {
	LBracePos token.Pos
	RBracePos token.Pos
	Nodes     []Node // statement | unblocked-block
}

func (b Block) Pos() token.Pos {
	return b.LBracePos
}

func (b Block) End() token.Pos {
	return b.RBracePos
}

// ----------------------------------------------------------------------------
// Statement Nodes

type Statement interface {
	Node
	statementNode()
}

type ReturnStatement struct {
	From       token.Pos
	SemiPos    token.Pos
	Expression Expression
}

func (r *ReturnStatement) Pos() token.Pos {
	return r.From
}
func (r *ReturnStatement) End() token.Pos {
	return r.SemiPos
}

func (s *ReturnStatement) statementNode() {}
