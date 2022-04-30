package ast

type pragmaValue struct {
	version    string
	expression string
}

type PragmaDirective struct {
	pragmaName  string
	pragmaValue pragmaValue
}

type FunctionDescriptor struct {
	name string
}

type StateMutability struct {
	pure bool
}

type ModifierList struct {
	stateMutability *StateMutability
}

type TypeName struct {
	elementalyTypeName string
}

type EventParameter struct {
	typeName *TypeName
}

type ParameterList struct {
	eventParameter *EventParameter
}

type ReturnParameters struct {
	parameterList *ParameterList
}

type FunctionDefinition struct {
	functionDescriptor *FunctionDescriptor
	modifierList       *ModifierList
	returnParameters   *ReturnParameters
}

type ContractPart struct {
	functionDefinition *FunctionDefinition
}

type ContractDefinition struct {
	contractPart *ContractPart
}

// A File node represents a Solidity source file.
type SourceUnit struct {
	PragmaDirective    *PragmaDirective
	ContractDefinition *ContractDefinition
}
