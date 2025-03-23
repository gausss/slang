package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var reservedWords = map[string]TokenKind{
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
}

type Scanner struct {
	source  string
	tokens  []*Token
	start   int
	current int
	line    int
}

func (scanner *Scanner) ScanTokens() []*Token {
	scanner.start = 0
	scanner.current = 0
	scanner.line = 1

	for !scanner.isAtEnd() {
		scanner.start = scanner.current
		scanner.scanToken()
	}

	scanner.tokens = append(scanner.tokens, &Token{EOF, "", nil, scanner.line})
	return scanner.tokens
}

func (scanner *Scanner) scanToken() {
	c := scanner.source[scanner.current]
	scanner.current++
	switch c {
	case '(':
		scanner.addToken(LEFT_PAREN)
	case ')':
		scanner.addToken(RIGHT_PAREN)
	case '{':
		scanner.addToken(LEFT_BRACE)
	case '}':
		scanner.addToken(RIGHT_BRACE)
	case ',':
		scanner.addToken(COMMA)
	case '.':
		scanner.addToken(DOT)
	case '-':
		scanner.addToken(MINUS)
	case '+':
		scanner.addToken(PLUS)
	case ';':
		scanner.addToken(SEMICOLON)
	case '*':
		scanner.addToken(STAR)
	case '!':
		if scanner.match('=') {
			scanner.addToken(BANG_EQUAL)
		} else {
			scanner.addToken(BANG)
		}
	case '=':
		if scanner.match('=') {
			scanner.addToken(EQUAL_EQUAL)
		} else {
			scanner.addToken(EQUAL)
		}
	case '<':
		if scanner.match('=') {
			scanner.addToken(LESS_EQUAL)
		} else {
			scanner.addToken(LESS)
		}
	case '>':
		if scanner.match('=') {
			scanner.addToken(GREATER_EQUAL)
		} else {
			scanner.addToken(GREATER)
		}
	case '/':
		if scanner.match('/') {
			// A comment goes until the end of the line.
			for scanner.peek() != '\n' && !scanner.isAtEnd() {
				scanner.advance()
			}
		} else {
			scanner.addToken(SLASH)
		}
	case '"':
		scanner.string()
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		scanner.line++
	default:
		if isDigit(c) {
			scanner.number()
		} else if isAlpha(c) {
			scanner.identifier()
		} else {
			Error(scanner.line, fmt.Sprintf("Unexpected character. %c", c))
		}
	}
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (scanner *Scanner) identifier() {
	for isAlphaNumeric(scanner.peek()) {
		scanner.advance()
	}

	value := scanner.source[scanner.start:scanner.current]
	kind, kindExists := reservedWords[value]

	if !kindExists {
		kind = IDENTIFIER
	}

	scanner.addToken(kind)
}

func (scanner *Scanner) number() {
	for isDigit(scanner.peek()) {
		scanner.advance()
	}

	// Look for a fractional part.
	if scanner.peek() == '.' && isDigit(scanner.peekNext()) {
		// Consume the "."
		scanner.advance()

		for isDigit(scanner.peek()) {
			scanner.advance()
		}
	}

	value, _ := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 64)
	scanner.addTokenLiteral(NUMBER, value)
}

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func (scanner *Scanner) addToken(tokenType TokenKind) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, &Token{tokenType, text, nil, scanner.line})
}

func (scanner *Scanner) addTokenLiteral(tokenType TokenKind, literal interface{}) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, &Token{tokenType, text, literal, scanner.line})
}

func (scanner *Scanner) advance() {
	scanner.current++
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) match(expected byte) bool {
	if scanner.isAtEnd() {
		return false
	}
	if scanner.source[scanner.current] != expected {
		return false
	}

	scanner.current++
	return true
}

func (scanner *Scanner) peek() byte {
	if scanner.isAtEnd() {
		return ' '
	}

	return scanner.source[scanner.current]
}

func (scanner *Scanner) peekNext() byte {
	if scanner.current+1 >= len(scanner.source) {
		return ' '
	}

	return scanner.source[scanner.current+1]
}

func (scanner *Scanner) string() {
	for scanner.peek() != '"' && !scanner.isAtEnd() {
		if scanner.peek() == '\n' {
			scanner.line++
		}
		scanner.advance()
	}

	if scanner.isAtEnd() {
		Error(scanner.line, "Unterminated string.")
		return
	}

	// The closing ".
	scanner.advance()

	// Trim the surrounding quotes.
	value := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.addTokenLiteral(STRING, value)
}
