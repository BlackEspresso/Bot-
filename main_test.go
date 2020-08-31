package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/BlackEspresso/Bot-/botscripting"
)

func TestBotRuntime(t *testing.T) {
	filePath := flag.String("file", "main.b++", "path to b++ file")
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	scripting := botscripting.NewScriptRuntime()
	err = scripting.RunScript(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	scripting.BotRuntime.SayToBot("hello")
	fmt.Println(scripting.BotRuntime.ListenToBot())
}
