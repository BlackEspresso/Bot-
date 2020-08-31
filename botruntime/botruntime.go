package botruntime

import (
	"math/rand"
	"strings"
)

type UserContext struct {
	Name            string
	LiveTime        int
	CurrentLiveTime int
}

type Context struct {
	LastState          *State
	UserContext        []*UserContext
	CurrentUserMessage string
}

type State struct {
	Name          string
	TrainingData  []string
	Reponses      []string
	Action        func(state *State, context *Context) bool
	InputContext  []*UserContext
	OutputContext []*UserContext
}

func (state *State) Run(context *Context) string {
	if state.Action != nil {
		state.Action(state, context)
	}
	return state.selectResponse(context)
}

func (s *State) selectResponse(context *Context) string {
	i := rand.Int31n(int32(len(s.Reponses)))
	return s.Reponses[i]
}

func (s *State) MatchAgainstTraining(text string) int {
	splitted := strings.Split(text, " ")
	best := -1
	for _, m := range s.TrainingData {
		splittedTraining := strings.Split(m, " ")
		score := match(splitted, splittedTraining)
		if score > best {
			best = score
		}
	}
	return best
}

func match(listA []string, listB []string) int {
	lookup := map[string]bool{}
	for _, x := range listA {
		lookup[strings.ToLower(strings.TrimSpace(x))] = true
	}

	count := 0
	for _, x := range listB {
		if _, ok := lookup[strings.ToLower(strings.TrimSpace(x))]; ok {
			count++
		}
	}
	return count
}

func NewState(name string) *State {
	state := &State{
		name,
		[]string{},
		[]string{},
		nil,
		[]*UserContext{},
		[]*UserContext{},
	}
	return state
}

type BotRuntime struct {
	states  []*State
	context *Context
}

func NewBotRuntime() *BotRuntime {
	return &BotRuntime{
		[]*State{},
		&Context{nil, []*UserContext{}, ""},
	}
}

func (bt *BotRuntime) GetStates() []*State {
	return bt.states
}

func (bt *BotRuntime) GetStateByName(name string) *State {
	for _, state := range bt.states {
		if state.Name == name {
			return state
		}
	}
	return nil
}

func (bt *BotRuntime) AddState(state *State) {
	bt.states = append(bt.states, state)
}

func (bt *BotRuntime) ListenToBot() string {
	var state *State
	if bt.context.LastState == nil && bt.context.CurrentUserMessage == "" {
		state = bt.findStartState()
	} else {
		state = bt.findState()
	}

	if state == nil {
		panic("state is null")
	}

	resp := state.Run(bt.context)
	bt.context.LastState = state
	return resp
}

func (bt *BotRuntime) findState() *State {
	state := bt.findStateByTrainingData()
	// extend logic here: make context aware
	return state
}

func (bt *BotRuntime) SayToBot(text string) {
	bt.context.CurrentUserMessage = text
}

func (bt *BotRuntime) findStateByTrainingData() *State {
	startStates := findStateByInputContext(bt.states, "")
	var bestState *State
	bestScore := -1
	for _, s := range startStates {
		score := s.MatchAgainstTraining(bt.context.CurrentUserMessage)
		if score > bestScore {
			bestScore = score
			bestState = s
		}
	}
	return bestState
}

func findStateByInputContext(states []*State, contextName string) []*State {
	possibleStarts := []*State{}
	for _, s := range states {
		if findContextByName(s.InputContext, contextName) {
			possibleStarts = append(possibleStarts, s)
		}
	}
	return possibleStarts
}

func findContextByName(list []*UserContext, text string) bool {
	if len(list) == 0 && len(text) == 0 {
		return true
	}
	for _, a := range list {
		if a.Name == text {
			return true
		}
	}
	return false
}

func (bt *BotRuntime) findStartState() *State {
	startStates := findStateByInputContext(bt.states, "")
	i := rand.Int31n(int32(len(startStates)))
	return startStates[i]
}
