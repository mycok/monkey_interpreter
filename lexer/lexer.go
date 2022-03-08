package lexer

import (
	"github.com/mycok/monkey_interpreter/token"
)

// Lexer represents a Lexer object / type.
type Lexer struct {
	input        string
	position     int  // current position in input (points to the current char / position of that char in the input)
	readPosition int  // current reading's position in input (after current char)
	char         byte // current char under examination
}

// New returns an initialized instance of a Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// Set all the remaining lexer fields by calling l.readChar.
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// Out of range scenario.
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the current token based on the value of l.char.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.char {
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '!':
		tok = l.makeTwoCharToken('=', token.NOTEQ, token.BANG)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case '=':
		tok = l.makeTwoCharToken('=', token.EQ, token.ASSIGN)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			literal := l.readIdentifiersAndNumbers(isLetter)

			// TODO: figure out a way to get rid of these if/else statements for easy readability.
			if literal == "else" && l.peekChar() == 'i' {
				l.readChar()
				nextLiteral := l.readIdentifiersAndNumbers(isLetter)

				tok.Literal = literal + nextLiteral
				tok.Type = token.LookupIndentifier(tok.Literal)

				return tok
			}

			tok.Literal = literal
			tok.Type = token.LookupIndentifier(tok.Literal)

			return tok
		} else if isDigit(l.char) {
			tok.Literal = l.readIdentifiersAndNumbers(isDigit)
			tok.Type = token.INT

			return tok

		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifiersAndNumbers(fn func(ch byte) bool) string {
	position := l.position

	for fn(l.char) {
		l.readChar()
	}

	// Return a slice of the input string.
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) makeTwoCharToken(char byte, twoCharType token.TokenType, defaultType token.TokenType) token.Token {
	if l.peekChar() == char {
		ch := l.char
		l.readChar()
		return token.Token{Type: twoCharType, Literal: string(ch) + string(l.char)}
	}

	return newToken(defaultType, l.char)
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
