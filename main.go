package main

import "fmt"

type TrafficState interface {
	State()
}

type RedState struct{}

func (rs *RedState) State() {
	fmt.Println("Current State is Red!")
}

type YellowState struct{}

func (rs *YellowState) State() {
	fmt.Println("Current State is Yellow!")
}

type GreenState struct{}

func (rs *GreenState) State() {
	fmt.Println("Current State is Green!")
}

type StateContext struct {
	currentState TrafficState
}

func (c *StateContext) SetContext(state TrafficState) {
	c.currentState = state
}

func (c *StateContext) GetState() {
	c.currentState.State()
}

func GetContext() *StateContext {
	return &StateContext{
		currentState: &RedState{},
	}
}

func main() {
	fmt.Println("Default State...")
	stateContext := GetContext()
	stateContext.GetState()
	fmt.Println()
	fmt.Println()

	stateContext.SetContext(&YellowState{})
	stateContext.GetState()

	stateContext.SetContext(&GreenState{})
	stateContext.GetState()

	stateContext.SetContext(&RedState{})
	stateContext.GetState()
}
