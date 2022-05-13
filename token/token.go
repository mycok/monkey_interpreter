package token

const (
	// ILLEGAL represents any character not understood by the language lexer.
	ILLEGAL = "ILLEGAL"

	// EOF represents end of file.
	EOF = "EOF"

	// IDENT such as add, foobar, x,y, ....etc
	IDENT = "IDENT"

	// INT such as 1234567890
	INT = "INT"

	// ASSIGN ... are some of the operators implemented in the language.
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

	// COMMA ... are some of the delimiters implemented in the language.
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION ... are some of the keywords implemented in the language.
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
