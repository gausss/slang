package main

import "fmt"

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
	default:
		Error(scanner.line, fmt.Sprintf("Unexpected character. %c", c))
	}
}

func (scanner *Scanner) addToken(tokenType TokenType) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, &Token{tokenType, text, nil, scanner.line})
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}
