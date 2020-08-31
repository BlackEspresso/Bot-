package botscripting

import (
	"testing"
)

func TestBotRuntime(t *testing.T) {
	scriptRt := NewScriptRuntime()
	err := RunScript(scriptRt, "print \"4\"")
	if err != nil {
		t.Error(err)
	}
}

func TestAddState(t *testing.T) {
	scriptRt := NewScriptRuntime()
	err := RunScript(scriptRt, "add state hello")
	if err != nil {
		t.Error(err)
	}

	if len(scriptRt.BotRuntime.GetStates()) == 0 {
		t.Error("state not added")
	}
}

func TestSetTrainingData(t *testing.T) {
	scriptRt := NewScriptRuntime()
	err := RunScript(scriptRt, "add state hello")
	if err != nil {
		t.Error(err)
	}

	err = RunScript(scriptRt, "set hello.responses [\"a\"]")
	if err != nil {
		t.Error(err)
	}

	if len(scriptRt.BotRuntime.GetStates()) == 0 {
		t.Error("state not added")
	}

	state := scriptRt.BotRuntime.GetStates()[0]
	if len(state.Reponses) == 0 {
		t.Error("no responses added")
	}
}
