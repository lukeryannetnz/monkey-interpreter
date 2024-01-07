// The evaluator package evaluates monkey source code using a tree walking interpreter.
package evaluator

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/object"
	"monkey-interpreter/token"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

func Eval(node ast.Node, environment *object.Environment) object.Object {
	switch node := node.(type) {
	// statements
	case *ast.Program:
		return evalProgram(node.Statements, environment)
	case *ast.ExpressionStatement:
		return Eval(node.Value, environment)
	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements, environment)
	case *ast.ReturnStatement:
		return &object.ReturnValue{Value: Eval(node.Value, environment)}
	case *ast.LetStatement:
		return evalLetStatement(node, environment)

	// expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, environment)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, environment)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, environment)
		if isError(right) {
			return right
		}
		return evalInfixExperession(node.Operator, left, right)
	case *ast.IfExpression:
		condition := Eval(node.Condition, environment)

		if condition == nil {
			return NULL
		}

		boolCondition, ok := condition.(*object.Boolean)

		// conditions which contain a value (not nil) which is not a bool are truthy
		if !ok || boolCondition.Value {
			return Eval(node.Consequence, environment)
		} else {
			if node.Alternative == nil {
				return NULL
			}
			return Eval(node.Alternative, environment)
		}
	case *ast.Identifier:
		return evalIdentifier(node, environment)
	}

	return nil
}

func isError(obj object.Object) bool {
	if obj == nil {
		return false
	}

	return obj.Type() == object.ERROR_OBJ
}

func evalInfixExperession(operator string, left, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left, right)
	}

	if left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ {
		return evalBooleanInfixExpression(operator, left, right)
	}

	if left.Type() != right.Type() {
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}

	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	// no need to unwrap value here, can rely on pointer comparison as nativeBoolToBooleanObject implementation
	// ensures that boolean objects are pointers to same singleton memory address. this is faster,
	// although I have no benchmarked it.

	switch operator {
	case token.EQ:
		return nativeBoolToBooleanObject(left == right)
	case token.NOT_EQ:
		return nativeBoolToBooleanObject(left != right)
	}

	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case token.PLUS:
		return &object.Integer{Value: leftVal + rightVal}
	case token.MINUS:
		return &object.Integer{Value: leftVal - rightVal}
	case token.ASTERISK:
		return &object.Integer{Value: leftVal * rightVal}
	case token.SLASH:
		return &object.Integer{Value: leftVal / rightVal}
	case token.LT:
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case token.GT:
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case token.EQ:
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case token.NOT_EQ:
		return nativeBoolToBooleanObject(leftVal != rightVal)
	}

	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func evalPrefixExpression(prefix string, right object.Object) object.Object {
	switch prefix {
	case token.BANG:
		return evalBangOperatorExpression(right)
	case token.MINUS:
		return evalMinusOperatorExpression(right)
	}

	return newError("unknown operator: %s%s", prefix, right.Type())
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: %s%s", "-", right.Type())
	}

	return &object.Integer{Value: -right.(*object.Integer).Value}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Boolean:
		return &object.Boolean{Value: !right.Value}
	case *object.Integer:
		return FALSE

	}
	return NULL
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalProgram(stmts []ast.Statement, environment *object.Environment) object.Object {
	var result object.Object
	result = NULL
	for _, statement := range stmts {
		result = Eval(statement, environment)

		// if we encounter a return statement or error, break execution
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStatement(stmts []ast.Statement, environment *object.Environment) object.Object {
	var result object.Object
	result = NULL
	for _, statement := range stmts {
		result = Eval(statement, environment)

		if result != nil && (result.Type() == object.RETURN_VALUE_OBJ || result.Type() == object.ERROR_OBJ) {
			return result
		}
	}

	return result
}

func evalLetStatement(statement *ast.LetStatement, environment *object.Environment) object.Object {
	result := Eval(statement.Value, environment)

	if isError(result) {
		return result
	}

	environment.Set(statement.Name.Value, result)
	return result
}

func evalIdentifier(node *ast.Identifier, environment *object.Environment) object.Object {

	value, ok := environment.Get(node.Value)

	if !ok {
		return newError("unknown identifier: %s", node.Value)
	}

	return value
}
