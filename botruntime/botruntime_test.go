package botruntime

import (
	"fmt"
	"testing"
)

func TestBotRuntime(t *testing.T) {
	bot := NewBotRuntime()
	state := NewState("hello")
	state.Reponses = []string{
		"Hi, how are you ?",
		"hello",
		"welcome",
	}

	bot.AddState(state)

	fmt.Println(bot.ListenToBot())
	bot.SayToBot("fine thanks")
}

func TestBotRuntime2(t *testing.T) {
	bot := NewBotRuntime()

	state := NewState("hello")
	state.Reponses = []string{
		"Hi, how are you ?",
		"hello",
		"welcome",
	}
	bot.AddState(state)

	state2 := NewState("feeling")
	state2.TrainingData = []string{
		"who are you?",
		"who do you do?",
	}

	state2.Reponses = []string{
		"i'm fine thanks",
		"everything fine",
	}

	bot.AddState(state2)

	bot.SayToBot("who you doing?")
	fmt.Println(bot.ListenToBot())
}
