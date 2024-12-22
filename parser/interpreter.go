package parser

import (
	"fmt"
	"lox/lexer"
	"reflect"
	"strconv"
)

type Interpreter struct {
}

func (a *Interpreter) evaluate(expr Expr) (any, error) {
	return expr.accept(a)
}

func isTruthy(obj any) bool {
	if obj == nil {
		return false
	}
	switch v := obj.(type) {
	case bool:
		return v
	default:
		return true
	}
}

func (a *Interpreter) visitBinary(expr *Binary) (any, error) {
	left, err := a.evaluate(expr.left)
	if err != nil {
		return nil, err
	}
	valueLeft, err := strconv.ParseFloat(fmt.Sprintf("%v", left), 64)
	if err != nil {
		return nil, err
	}
	right, err := a.evaluate(expr.right)
	if err != nil {
		return nil, err
	}
	valueRight, err := strconv.ParseFloat(fmt.Sprintf("%v", right), 64)
	if err != nil {
		return nil, err
	}

	switch expr.operator.TokenType() {
	case lexer.MINUS:
		return valueLeft - valueRight, nil
	case lexer.PLUS:
		switch v := left.(type) {
		case string:
			strRight := right.(string)
			return v + strRight, nil
		case float64:
			return valueLeft + valueRight, nil
		default:
			return nil, NewRuntimeErr("invalid operand value")
		}
	case lexer.STAR:
		return valueLeft * valueRight, nil
	case lexer.SLASH:
		if valueRight == 0 {
			return nil, NewRuntimeErr("division by zero error")
		}
		return valueLeft / valueRight, nil
	case lexer.GREATER_EQUAL:
		return valueLeft >= valueRight, nil
	case lexer.GREATER:
		return valueLeft > valueRight, nil
	case lexer.LESS:
		return valueLeft < valueRight, nil
	case lexer.LESS_EQUAL:
		return valueLeft <= valueRight, nil
	case lexer.BANG_EQUAL:
		return !reflect.DeepEqual(left, right), nil
	case lexer.EQUAL_EQUAL:
		return reflect.DeepEqual(left, right), nil
	default:
		return nil, NewRuntimeErr("invalid operator")
	}
}

func (a *Interpreter) visitGrouping(expr *Grouping) (any, error) {
	return a.evaluate(expr.expression)
}

func (a *Interpreter) visitLiteral(expr *Literal) (any, error) {
	return expr.value, nil
}

func (a *Interpreter) visitUnary(expr *Unary) (any, error) {
	val, err := a.evaluate(expr.right)
	if err != nil {
		return nil, err
	}
	right := fmt.Sprintf("%v", val)

	switch expr.operator.TokenType() {
	case lexer.MINUS:
		f, err := strconv.ParseFloat(right, 64)
		if err != nil {
			return nil, err
		}
		return -f, nil
	case lexer.BANG:
		return !isTruthy(expr.right), nil
	}

	return nil, fmt.Errorf("Unary expression is invalid")
}
