package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BlackEspresso/Bot-/botscripting"
)

func main() {
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

	fmt.Println("Entering REPL, type exit to quit")

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}
		scripting.BotRuntime.SayToBot(input)
		fmt.Println(">> ", scripting.BotRuntime.ListenToBot())
	}
}
