package botscripting

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BlackEspresso/Bot-/botruntime"
	"github.com/BlackEspresso/Bot-/botscripting/parser"
)

type ScriptRuntime struct {
	BotRuntime *botruntime.BotRuntime
	Vars       map[string]string
}

func NewScriptRuntime() *ScriptRuntime {
	bot := botruntime.NewBotRuntime()
	script := &ScriptRuntime{bot, map[string]string{}}
	return script
}

func RunScript(scriptRuntime *ScriptRuntime, text string) error {
	parser := parser.NewParser(text)
	ast, err := parser.Parse()
	if err != nil {
		return err
	}

	return evaluateAst(ast, scriptRuntime)
}

func evaluateAst(token *parser.Token, script *ScriptRuntime) error {
	if token == nil {
		return errors.New("token is nil")
	}
	module := token
	if module.Kind != parser.ModuleToken {
		return errors.New("Doesn't start with module token")
	}
	for _, command := range module.ChildToken {
		script.evalCommand(command)
	}
	return nil
}

func (runtime *ScriptRuntime) evalCommand(command *parser.Token) {
	commandName := command.ChildToken[0]
	if commandName.Kind != parser.IdentifierToken {
		panic("first command parameter is not an idientifer " + commandName.Text)
	}
	switch strings.TrimSpace(commandName.Text) {
	case "print":
		runtime.evalPrint(command.ChildToken[1:])
	case "add":
		runtime.evalAddCommand(command.ChildToken[1:])
	case "set":
		runtime.evalSetCommand(command.ChildToken[1:])
	default:
		panic("unkown command " + commandName.Text)
	}
}

func (runtime *ScriptRuntime) evalSetCommand(command []*parser.Token) {
	if len(command) == 0 {
		panic("set command needs a statename")
	}
	target := command[0]
	if target.Kind == parser.IdentifierToken {
		// todo: set variable
	} else if target.Kind == parser.PropertyAccessToken {
		name := target.ChildToken[0]
		prop := target.ChildToken[1]
		if !isIdentifier(name) || !isIdentifier(prop) {
			panic(target.Text + " is not a valid property")
		}

		state := runtime.BotRuntime.GetStateByName(strings.TrimSpace(name.Text))
		switch strings.TrimSpace(prop.Text) {
		case "responses":
			state.Reponses = evalStringArray(command[1])
		default:
			panic("unkown property " + prop.Text + " on '" + name.Text + "'")
		}
	}
}

func evalStringArray(token *parser.Token) []string {
	if token.Kind != parser.ArrayToken {
		panic("not an array " + token.Text)
	}
	arr := []string{}
	for _, x := range token.ChildToken {
		arr = append(arr, x.Text)
	}
	return arr
}

func (runtime *ScriptRuntime) evalAddCommand(command []*parser.Token) {
	if len(command) == 0 {
		panic("add command has no type")
	}
	commandType := command[0]
	if isKeyword(commandType, "state") {
		if len(command) < 2 {
			panic("missing state name")
		}
		state := botruntime.NewState(strings.TrimSpace(command[1].Text))
		runtime.BotRuntime.AddState(state)
	} else {
		panic("add-command didnt find keyword " + commandType.Text)
	}
}

func isIdentifier(token *parser.Token) bool {
	return token.Kind == parser.IdentifierToken
}

func isKeyword(token *parser.Token, exp string) bool {
	return isIdentifier(token) &&
		strings.ToLower(strings.TrimSpace(token.Text)) ==
			strings.ToLower(strings.TrimSpace(exp))
}

func (runtime *ScriptRuntime) evalPrint(command []*parser.Token) {
	fmt.Print(">>")
	for _, c := range command {
		switch c.Kind {
		case parser.StringToken:
			fmt.Print(c.Text)
		case parser.NumberToken:
			fmt.Print(c.Text)
		default:
			panic("cant print " + c.Text)
		}
	}
	fmt.Print("\n")
}
