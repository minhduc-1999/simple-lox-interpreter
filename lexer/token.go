package lexer

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   string
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal string, line int) Token {
	return Token{
		tokenType,
		lexeme,
		literal,
		line,
	}
}
