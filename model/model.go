/*
Package model contains abstractions that will be used throughout the module.
*/
package model

// VisitFunction is called when visiting a Node.
type VisitFunction func(n Node) interface{}

// Node represents a node in a data structure.
type Node interface {
	Data() interface{}
	Neighborhood() []Node
}

// Void is a void type
type Void struct{}
