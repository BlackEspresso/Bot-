package parser

import (
	"errors"
	"strconv"
)

type TokenKind int

const (
	ArrayToken TokenKind = iota
	StringToken
	NumberToken
	IdentifierToken
	CommandToken
	ModuleToken
	BlockToken
	PropertyAccessToken
)

type Token struct {
	Kind       TokenKind
	Text       string
	ChildToken []*Token
}

func (t *Token) String() string {
	return "Token{" + strconv.Itoa(int(t.Kind)) + ", " + t.Text + "}"
}

func (t *Token) addChild(token *Token) {
	t.ChildToken = append(t.ChildToken, token)
	t.Text += " " + token.Text
}

type Parser struct {
	text     []rune
	position int
	current  rune
	Error    error
}

func NewToken(kind TokenKind, text string) *Token {
	return &Token{kind, text, []*Token{}}
}

func NewParser(text string) *Parser {
	runes := []rune(text)
	runes = append(runes, rune(0))
	return &Parser{runes, -1, 'm', nil}
}

func (p *Parser) hasNext() bool {
	return p.position < len(p.text)-1
}

func (p *Parser) revert() {
	p.position--
	p.current = p.text[p.position]
}

func (p *Parser) next() bool {
	if !p.hasNext() {
		return false
	}
	p.position++
	p.current = p.text[p.position]
	return true
}

func (p *Parser) nextChar() bool {
	for p.next() {
		if !isWhiteSpace(p.current) {
			return true
		}
	}
	return false
}

func (p *Parser) eatWhitespace() bool {
	if !isWhiteSpace(p.current) {
		return true
	}

	for p.next() {
		if !isWhiteSpace(p.current) {
			return true
		}
	}
	return false
}

func (p *Parser) skipLine() {
	newLineFound := false
	for p.next() {
		if isNewLine(p.current) {
			newLineFound = true
		}
		if !isNewLine(p.current) && newLineFound {
			return
		}
	}
}

func (p *Parser) parseComment() {
	p.skipLine()
}

func (p *Parser) Parse() (*Token, error) {
	token := NewToken(ModuleToken, "")
	for p.nextChar() {
		if p.current == '#' {
			p.parseComment()
		} else if isAlphabetical(p.current) {
			token.addChild(p.parseCommand())
		} else {
			p.Error = errors.New("didnt excpect " + string(p.current))
		}

		if p.Error != nil {
			break
		}
	}
	return token, p.Error
}

func (p *Parser) parseCommand() *Token {
	token := NewToken(CommandToken, "")
	for {
		if isAlphabetical(p.current) {
			token.addChild(p.parsePropertyAccess())
		} else if p.current == '[' {
			token.addChild(p.praseArray())
		} else if p.current == '{' {
			token.addChild(p.parseBlock())
		} else if p.current == '#' {
			p.parseComment()
		} else if p.current == '"' {
			token.addChild(p.parseString())
		} else if isNewLine(p.current) || p.current == '}' || p.current==0 {
			break
		} else {
			p.Error = errors.New("unkown command " + string(p.current))
			break
		}

		if !p.eatWhitespace() {
			break
		}
	}
	return token
}

func (p *Parser) parsePropertyAccess() *Token {
	i1 := p.parseIdentifier()
	list := []*Token{i1}

	for p.current == '.' && p.next() {
		list = append(list, p.parseIdentifier())
	}

	if len(list) == 1 {
		return i1
	}

	token := NewToken(PropertyAccessToken, "")
	for x, identifier := range list {
		token.addChild(identifier)

		if len(token.ChildToken) == 2 &&  x != len(list)-1 {
			nextToken := NewToken(PropertyAccessToken, "")
			nextToken.addChild(token)
			token = nextToken
		}
	}
	return token
}

func (p *Parser) nextIs(predicat func(r rune) bool) bool {
	isNext := false
	if p.next() {
		isNext = predicat(p.current)
		p.revert()
	}
	return isNext
}

func (p *Parser) parseBlock() *Token {
	token := NewToken(BlockToken, "")
	for p.nextChar() {
		if p.current == '}' {
			break
		}
		token.addChild(p.parseCommand())
		if p.Error != nil {
			break
		}
	}
	return token
}

func (p *Parser) parseIdentifier() *Token {
	text := ""
	for isAlphabetical(p.current) || isNumber(p.current) || p.current == '_' {
		text += string(p.current)
		if !p.next() {
			return NewToken(IdentifierToken, text)
		}
	}
	return NewToken(IdentifierToken, text)
}

func (p *Parser) praseArray() *Token {
	token := NewToken(ArrayToken, "")
	for p.nextChar() {
		if isNumber(p.current) {
			token.addChild(p.parseNumber())
		} else if p.current == '"' {
			token.addChild(p.parseString())
		} else if p.current == ',' {
			continue
		} else if p.current == ']' {
			p.next()
			break
		}

		if p.Error != nil {
			break
		}
	}
	return token
}

func (p *Parser) parseString() *Token {
	text := string(p.current)
	for p.next() {
		text += string(p.current)
		if p.current == '"' {
			break
		}
	}
	if p.current != '"' {
		p.Error = errors.New("string '" + text + "' not closed")
	}
	p.next()
	return NewToken(NumberToken, text)
}

func (p *Parser) parseNumber() *Token {
	text := ""
	for isNumber(p.current) {
		text += string(p.current)
		if !p.next() {
			break
		}
	}
	p.revert()
	return NewToken(NumberToken, text)
}

func isNewLine(r rune) bool {
	return r == 0x10 || r == 0x13
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isWhiteSpace(r rune) bool {
	return isNewLine(r) || r == ' ' || r == '\t'
}

func isAlphabetical(r rune) bool {
	return r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
}
