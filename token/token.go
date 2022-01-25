package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// User defined identifiers + literals
	IDENT = "IDENT" // add, foobar, x,y, ....etc
	INT   = "INT"   // 1234567890

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	ELSE_IF  = "ELSE IF"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"elseif": ELSE_IF,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIndentifier(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}
