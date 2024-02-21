// Package object contains the internal object system used when evaluating monkey source code.
//
// Every value in monkey is represented as a struct which implements an Object interface.
package object

import (
	"bytes"
	"fmt"
	"monkey-interpreter/ast"
	"strings"
)

type ObjectType string

const (
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
)

func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s, enclosingEnvironment: nil}
}

func ExtendEnvironment(enclosingEnv *Environment) *Environment {
	s := make(map[string]Object)

	return &Environment{store: s, enclosingEnvironment: enclosingEnv}
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Environment struct {
	store                map[string]Object
	enclosingEnvironment *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.enclosingEnvironment != nil {
		obj, ok = e.enclosingEnvironment.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
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

type Error struct {
	Message string
}

func (e *Error) Inspect() string  { return e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {/n}")
	out.WriteString(f.Body.String())
	out.WriteString("/n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return FUNCTION_OBJ }
func (s *String) Inspect() string  { return s.Value }
