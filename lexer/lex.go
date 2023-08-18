package lexer

import (
	"fmt"
	"golang/utils"
	"os"
)

type lexer struct {
	source []rune
	tokens []Token
	loc    *Location
	mode   TokenType
	begin  *Location
	eof    bool
}

func (l *lexer) at() rune {
	return l.source[l.loc.Idx]
}

func (l lexer) peek() rune {
	return l.source[l.loc.Idx+1]
}

func (l *lexer) eat() rune {
	ch := l.source[l.loc.Idx]

	l.loc.Idx++
	l.loc.Col++
	if ch == '\n' {
		l.loc.Col = 0
		l.loc.Row++
	}

	if l.loc.Idx >= len(l.source) {
		l.eof = true
	}
	return ch
}

func (l lexer) all(start, end int) string {
	return string(l.source[start:(end + 1)])
}

func (l lexer) ahead(ahead int) string {
	return l.all(l.loc.Idx, l.loc.Idx)
}

func (l *lexer) start(mode TokenType) {
	l.mode = mode
	*l.begin = *l.loc
	l.eat()
}

func (l *lexer) end() *Token {
	value := l.all(l.begin.Idx, l.loc.Idx-1)

	var tt TokenType
	if l.mode == Identifier && IsKeyword(value) {
		tt = Keyword
	} else {
		tt = l.mode
	}

	token := CreateToken(
		tt,
		value,
		l.begin.Row,
		l.begin.Col,
		l.begin.Idx,
	)

	l.mode = None

	return token
}

func (l *lexer) add(token *Token) {
	l.tokens = append(l.tokens, *token)
}

func (l *lexer) singleToken(tokentype TokenType) {
	if l.mode == Identifier {
		l.add(l.end())
	}
	l.add(CreateToken(
		tokentype,
		string(l.at()),
		l.loc.Row,
		l.loc.Col,
		l.loc.Idx,
	))

	l.eat()
}

func (l *lexer) splitAdd(token *Token, endings ...TokenType) {
	if utils.Contains(endings, l.mode) {
		l.add(l.end())
	}
	l.add(token)

	l.eat()
}

func Lex(source string) *[]Token {
	l := &lexer{
		[]rune(source),
		[]Token{},
		CreateLocation(0, 0, 0),
		None,
		CreateLocation(0, 0, 0),
		false,
	}

	for !l.eof {
		// TODO: multi-character key tokens (>=, ++, etc)
		// TODO: add comments & multiline comments
		switch l.at() {
		case ' ':
			l.singleToken(Space)
		case '\n':
			l.singleToken(NewLine)
		case '(':
			l.singleToken(LeftParen)
		case ')':
			l.singleToken(RightParen)
		case '[':
			l.singleToken(LeftBracket)
		case ']':
			l.singleToken(RightBracket)
		case '{':
			l.singleToken(LeftBrace)
		case '}':
			l.singleToken(RightBrace)
		case '+', '-', '*', '/', '%':
			l.singleToken(Operator)
		case '>', '<':
			l.singleToken(Comparator)
		case ',':
			l.singleToken(Delimiter)
		default:
			if l.mode == None {
				l.start(Identifier)
			} else {
				l.eat()
			}
		}
	}
	if l.mode != None {
		l.add(l.end())
	}
	l.add(CreateToken(EOF, "EOF", l.loc.Row, l.loc.Col, l.loc.Idx))

	return &l.tokens
}

func Filter(tokens *[]Token) *[]Token {
	filter := func(item Token) bool {
		return item.Type != Space && item.Type != NewLine
	}
	filtered := utils.Filter(*tokens, filter)
	return &filtered
}

func Stringify(tokens *[]Token) string {
	str := "[]Token{\n"
	for _, token := range *tokens {
		str += "    " + fmt.Sprintln(token)
	}

	return str + "}"
}

func GetSourceCode(location string) (string, error) {
	content, err := os.ReadFile(location)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func Save(tokens *[]Token, location string) error {
	return utils.SaveFile([]byte(Stringify(tokens)), location)
}
