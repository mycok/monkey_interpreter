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
	NOTEQ    = "!="

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
	ELSEIF   = "ELSE IF"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"elseif": ELSEIF,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// TokenType represents the type of a token.
type TokenType string

// Token represents an type created by a lexer type.
type Token struct {
	Type    TokenType
	Literal string
}

// LookupIndentifier performs a map lookup based on the provided identifier string and
// returns the matching type constant.
func LookupIndentifier(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}
