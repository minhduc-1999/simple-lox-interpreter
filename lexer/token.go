package lexer

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func (t Token) Line() int {
	return t.line
}

func (t Token) TokenType() TokenType {
	return t.tokenType
}

func (t Token) Literal() any {
	return t.literal
}

func NewToken(tokenType TokenType, lexeme string, literal string, line int) Token {
	return Token{
		tokenType,
		lexeme,
		literal,
		line,
	}
}

func (t Token) String() string {
	switch t.tokenType {
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	case DOT:
		return "."
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SEMICOLON:
		return ";"
	case SLASH:
		return "/"
	case STAR:
		return "*"
	case BANG:
		return "!"
	case BANG_EQUAL:
		return "!="
	case EQUAL:
		return "="
	case EQUAL_EQUAL:
		return "=="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case IDENTIFIER:
	case STRING:
	case NUMBER:
		return t.lexeme
	case AND:
		return "&&"
	case CLASS:
		return "class"
	case ELSE:
		return "else"
	case FALSE:
		return "false"
	case FUN:
		return "fun"
	case FOR:
		return "for"
	case IF:
		return "if"
	case NIL:
		return "nil"
	case OR:
		return "or"
	case PRINT:
		return "print"
	case RETURN:
		return "return"
	case SUPER:
		return "super"
	case THIS:
		return "this"
	case TRUE:
		return "true"
	case VAR:
		return "var"
	case WHILE:
		return "while"
	default:
		return ""
	}
	return ""
}
