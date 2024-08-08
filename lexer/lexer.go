package lexer

import (
	"fmt"
)

type Lexer struct {
	source               string
	tokens               []Token
	errors               []string
	start, current, line int
	keywords             map[string]TokenType
}

func NewLexer(src string) Lexer {
	return Lexer{
		source: src,
		tokens: []Token{},
		errors: []string{},
		keywords: map[string]TokenType{
			"and":    AND,
			"class":  CLASS,
			"else":   ELSE,
			"false":  FALSE,
			"for":    FOR,
			"fun":    FUN,
			"if":     IF,
			"nil":    NIL,
			"or":     OR,
			"print":  PRINT,
			"return": RETURN,
			"super":  SUPER,
			"this":   THIS,
			"true":   TRUE,
			"var":    VAR,
			"while":  WHILE,
		},
	}
}

func (s *Lexer) ScanToken() ([]Token, []string) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", "", s.line))
	return s.tokens, s.errors
}

func (s *Lexer) advance() uint8 {
	c := s.source[s.current]
	s.current += 1
	return c
}

func (s *Lexer) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addSingleToken(LEFT_PAREN)
	case ')':
		s.addSingleToken(RIGHT_PAREN)
	case '{':
		s.addSingleToken(LEFT_BRACE)
	case '}':
		s.addSingleToken(RIGHT_BRACE)
	case ',':
		s.addSingleToken(COMMA)
	case '.':
		s.addSingleToken(DOT)
	case '-':
		s.addSingleToken(MINUS)
	case '+':
		s.addSingleToken(PLUS)
	case ';':
		s.addSingleToken(SEMICOLON)
	case '*':
		s.addSingleToken(STAR)
	case '!':
		if s.match('=') {
			s.addSingleToken(BANG_EQUAL)
		} else {
			s.addSingleToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addSingleToken(EQUAL_EQUAL)
		} else {
			s.addSingleToken(EQUAL)
		}
	case '>':
		if s.match('=') {
			s.addSingleToken(GREATER_EQUAL)
		} else {
			s.addSingleToken(GREATER)
		}
	case '<':
		if s.match('=') {
			s.addSingleToken(LESS_EQUAL)
		} else {
			s.addSingleToken(LESS)
		}
	case '/':
		if s.match('/') {
			// skip comment
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addSingleToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line += 1
		break
	case '"':
		s.string()
		break
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			s.report("", fmt.Sprintf("Unexpected character %c", c))
		}
	}
}

func (s *Lexer) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType, found := s.keywords[text]
	if found {
		s.addSingleToken(tokenType)
	} else {
		s.addToken(IDENTIFIER, text)
	}

}

func isAlphaNumeric(c uint8) bool {
	return isAlpha(c) || isDigit(c)
}

func isAlpha(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func (s *Lexer) peekNext() uint8 {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Lexer) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	s.addToken(NUMBER, s.source[s.start:s.current])
}

func (s *Lexer) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line += 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.report("", fmt.Sprintf("Unterminated string"))
		return
	}

	// Consume closing "
	s.advance()

	text := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, text)
}

func (s Lexer) peek() uint8 {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Lexer) report(where string, msg string) {
	err := fmt.Sprintf("[line %d] Error %v: %v", s.line, where, msg)
	s.errors = append(s.errors, err)
}

func (s *Lexer) addSingleToken(t TokenType) {
	s.addToken(t, "")
}

func (s *Lexer) addToken(t TokenType, literal string) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(t, text, literal, s.line))
}

func (s Lexer) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Lexer) match(expected uint8) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current += 1
	return true
}
