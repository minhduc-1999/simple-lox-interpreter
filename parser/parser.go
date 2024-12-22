package parser

import (
	"errors"
	"fmt"
	"lox/lexer"
)

type Parser struct {
	tokens  []lexer.Token
	current int
}

func NewParser(tokens []lexer.Token) Parser {
	return Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) expression() (Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.match(lexer.BANG, lexer.BANG_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = NewBinary(expr, operator, right)
	}
	return expr, nil
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(lexer.GREATER, lexer.GREATER_EQUAL, lexer.LESS, lexer.LESS_EQUAL) {
		op := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = NewBinary(expr, op, right)
	}
	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(lexer.MINUS, lexer.PLUS) {
		op := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = NewBinary(expr, op, right)
	}
	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(lexer.SLASH, lexer.STAR) {
		op := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = NewBinary(expr, op, right)
	}
	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(lexer.BANG, lexer.MINUS) {
		op := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return NewUnary(op, right), nil
	}
	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	if p.match(lexer.FALSE) {
		return NewLiteral(false), nil
	}
	if p.match(lexer.TRUE) {
		return NewLiteral(true), nil
	}
	if p.match(lexer.NIL) {
		return NewLiteral(nil), nil
	}
	if p.match(lexer.NUMBER, lexer.STRING) {
		return NewLiteral(p.previous().Literal()), nil
	}
	if p.match(lexer.LEFT_PAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		return NewGrouping(expr), nil
	}
	return nil, fmt.Errorf("not expected expression: %v", p.peek())
}

func (p Parser) previous() lexer.Token {
	return p.tokens[p.current-1]
}

func (p Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p *Parser) match(tokenTypes ...lexer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p Parser) check(tokenType lexer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType() == tokenType
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current += 1
	}
	return p.previous()
}

func (p Parser) isAtEnd() bool {
	return p.peek().TokenType() == lexer.EOF
}

func (p *Parser) consume(t lexer.TokenType, msg string) (lexer.Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}
	return lexer.Token{}, errors.New(msg)
}

func (p Parser) error(token lexer.Token, msg string) error {
	if token.TokenType() == lexer.EOF {
		return fmt.Errorf("line %v at end %v", token.Line(), msg)
	}
	return fmt.Errorf("line at %v, %v", token.Line(), msg)
}

func (p Parser) report(line int, where string, msg string) {
	fmt.Printf("[line %d] Error %v: %v", line, where, msg)
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType() == lexer.SEMICOLON {
			return
		}
		switch p.peek().TokenType() {
		case lexer.CLASS:
		case lexer.FUN:
		case lexer.VAR:
		case lexer.FOR:
		case lexer.IF:
		case lexer.WHILE:
		case lexer.RETURN:
		case lexer.PRINT:
			return
		}
		p.advance()
	}
}

func (p *Parser) Parse() Expr {
	result, err := p.expression()
	if err != nil {
		// TODO: enter panic mode
		fmt.Printf("%v\n", err)
		return nil
	}
	return result
}
