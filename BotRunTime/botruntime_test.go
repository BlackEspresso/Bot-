package botruntime

import (
	"fmt"
	"testing"
)

func TestBotRuntime(t *testing.T) {
	bt := NewBotRuntime()
	state := NewState("hello")
	state.Reponses = []string{
		"Hi, how are you ?",
		"hello",
		"welcome",
	}

	bt.AddState(state)

	fmt.Println(bt.Talk())
	bt.TalkToBot("fine thanks")
}

func TestBotRuntime2(t *testing.T) {
	bt := NewBotRuntime()

	state := NewState("hello")
	state.Reponses = []string{
		"Hi, how are you ?",
		"hello",
		"welcome",
	}
	bt.AddState(state)

	state2 := NewState("feeling")
	state2.TrainingData = []string{
		"who are you?",
		"who do you do?",
	}

	state2.Reponses = []string{
		"i'm fine thanks",
		"everything fine",
	}

	bt.AddState(state2)

	bt.TalkToBot("who you doing?")
	fmt.Println(bt.Talk())
}
