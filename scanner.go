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
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		scanner.line++
	default:
		Error(scanner.line, fmt.Sprintf("Unexpected character. %c", c))
	}
}

func (scanner *Scanner) addToken(tokenType TokenKind) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, &Token{tokenType, text, nil, scanner.line})
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
		return '0'
	}

	return scanner.source[scanner.current]
}
