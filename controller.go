package main

import "net/http"

type fullAction struct {
	action            func(writer http.ResponseWriter, request *http.Request)
	requestStructFunc func() interface{}
}

type controller struct {
	actions map[string]fullAction
}

func NewController() *controller {
	return &controller{actions: make(map[string]fullAction)}
}

func (c *controller) SetAction(pattern string, action func(writer http.ResponseWriter, request *http.Request), fn func() interface{}) *controller {
	var act = c.actions[pattern]
	act.action = action
	act.requestStructFunc = fn
	c.actions[pattern] = act

	return c
}

func (c *controller) AddAction(pattern string, action func(writer http.ResponseWriter, request *http.Request)) *controller {
	var act = c.actions[pattern]
	act.action = action
	c.actions[pattern] = act

	return c
}

func (c *controller) AddStructFunc(pattern string, fn func() interface{}) *controller {
	var act = c.actions[pattern]
	act.requestStructFunc = fn
	c.actions[pattern] = act

	return c
}

func (c *controller) SetActions(actions map[string]fullAction) {
	c.actions = actions
}

func (c *controller) RemoveAction(pattern string) {
	delete(c.actions, pattern)
}

func (c *controller) UnsetActions() {
	c.actions = make(map[string]fullAction)
}

func (c controller) GetAction(pattern string) func(writer http.ResponseWriter, request *http.Request) {
	return c.actions[pattern].action
}

func (c controller) GetStructFunc(pattern string) func() interface{} {
	return c.actions[pattern].requestStructFunc
}
