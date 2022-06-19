package ast

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

type ContractPart struct {
	FunctionDefinition *FunctionDefinition
}

type ContractDefinition struct {
	ContractPart *ContractPart
}

// A File node represents a Solidity source file.
type SourceUnit struct {
	PragmaDirective    *PragmaDirective
	ContractDefinition *ContractDefinition
	FunctionDefinition *FunctionDefinition
}
