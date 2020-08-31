package parser

import (
	"fmt"
	"testing"
)

func (token *Token) IsOfKind(kind TokenKind) bool {
	return token.Kind == kind
}

func (token *Token) IsModuleAndCommand() (bool, *Token) {
	if !token.IsOfKind(ModuleToken) {
		return false, nil
	}

	if !token.ChildToken[0].IsOfKind(CommandToken) {
		return false, nil
	}

	return true, token.ChildToken[0]
}

func TestParseNumberArray(t *testing.T) {
	p := NewParser("print [1,2,5]")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v", command.ChildToken[1])

	if len(command.ChildToken[1].ChildToken) != 3 {
		t.Error("wrong token length")
	}
}

func TestParseStringArray(t *testing.T) {
	p := NewParser("print [\"test\",\"test3\"]")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v", command.ChildToken[1])

	if len(command.ChildToken[1].ChildToken) != 2 {
		t.Error("wrong token length")
	}
}

func TestParseCommand(t *testing.T) {
	p := NewParser("set test")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken)

	if len(command.ChildToken) != 2 {
		t.Error("wrong token length")
	}
}

func TestParseCommands(t *testing.T) {
	p := NewParser("set\nprint")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, _ := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", token.ChildToken)

	if len(token.ChildToken) != 2 {
		t.Error("wrong token length")
	}
}

func TestParseComment(t *testing.T) {
	p := NewParser("set #test")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken)

	if len(command.ChildToken) != 1 {
		t.Error("wrong token length")
	}
}

func TestParseIdentifier(t *testing.T) {
	p := NewParser("set n3409_4")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken)

	if len(command.ChildToken) != 2 {
		t.Error("wrong token length")
	}
}

func TestParseString(t *testing.T) {
	p := NewParser("set \"hallo\"")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken)

	if len(command.ChildToken) != 2 {
		t.Error("wrong token length")
	}
}

func TestParsePropertyAccess2(t *testing.T) {
	p := NewParser("print hallo.mm")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken)

	if len(command.ChildToken) != 2 {
		t.Error("wrong token length")
	}

	prop := command.ChildToken[1]
	if len(prop.ChildToken) != 2 {
		t.Error("should have 2 child tokens", len(prop.ChildToken))
	}
}

func TestParsePropertyAccess3(t *testing.T) {
	p := NewParser("print hallo.test.k mm")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken)

	if len(command.ChildToken) != 3 {
		t.Error("wrong token length")
	}
}

func TestParseBlock(t *testing.T) {
	p := NewParser("set {http}")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken)

	if !command.ChildToken[1].IsOfKind(BlockToken) {
		t.Error("not a block")
	}

	if len(command.ChildToken) != 2 {
		t.Error("wrong token length")
	}
}

func TestParseBlockWithParamter(t *testing.T) {
	p := NewParser("set {http[1,2]}")
	token, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	ok, command := token.IsModuleAndCommand()
	if !ok {
		t.Error("not a correct module start")
	}

	fmt.Printf("%+v\n", command.ChildToken[1].ChildToken[0])

	if !command.ChildToken[1].IsOfKind(BlockToken) {
		t.Error("not a block")
	}

	if len(command.ChildToken) != 2 {
		t.Error("wrong token length")
	}
}
