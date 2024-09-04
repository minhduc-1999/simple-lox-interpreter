package parser

import "lox/lexer"

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

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()
	for p.match(lexer.BANG, lexer.BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()
	for p.match(lexer.GREATER, lexer.GREATER_EQUAL, lexer.LESS, lexer.LESS_EQUAL) {
		op := p.previous()
		right := p.term()
		expr = NewBinary(expr, op, right)
	}
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match(lexer.MINUS, lexer.PLUS) {
		op := p.previous()
		right := p.factor()
		expr = NewBinary(expr, op, right)
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()
	for p.match(lexer.SLASH, lexer.STAR) {
		op := p.previous()
		right := p.unary()
		expr = NewBinary(expr, op, right)
	}
	return expr
}

func (p *Parser) unary() Expr {
	if p.match(lexer.BANG, lexer.MINUS) {
		op := p.previous()
		right := p.unary()
		return NewUnary(op, right)
	}
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(lexer.FALSE) {
		return NewLiteral(false)
	}
	if p.match(lexer.TRUE) {
		return NewLiteral(true)
	}
	if p.match(lexer.NIL) {
		return NewLiteral(nil)
	}
	if p.match(lexer.NUMBER, lexer.STRING) {
		return NewLiteral(p.previous().Literal())
	}
	if p.match(lexer.LEFT_PAREN) {
		expr := p.expression()
		return NewGrouping(expr)
	}
	return nil
}

func (p Parser) previous() lexer.Token {
	return p.tokens[p.current-1]
}

func (p Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p Parser) match(tokenTypes ...lexer.TokenType) bool {
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
