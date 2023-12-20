// Package object contains the internal object system used when evaluating monkey source code.
//
// Every value in monkey is represented as a struct which implements an Object interface.
package object

import "fmt"

type ObjectType string

const (
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Inspect() string  { return r.Value.Inspect() }
func (r *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }
