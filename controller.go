package main

import "net/http"

type controller struct {
	actions map[string]func(writer http.ResponseWriter, request *http.Request)
}

func NewController() *controller {
	return &controller{actions: make(map[string]func(writer http.ResponseWriter, request *http.Request))}
}

func (c *controller) AddAction(pattern string, action func(writer http.ResponseWriter, request *http.Request))  {
	c.actions[pattern] = action
}

func (c *controller) SetActions(actions map[string]func(writer http.ResponseWriter, request *http.Request))  {
	c.actions = actions
}

func (c *controller) RemoveAction(pattern string) {
	delete(c.actions, pattern)
}

func (c *controller) UnsetActions() {
	c.actions = make(map[string]func(writer http.ResponseWriter, request *http.Request))
}

func (c controller) GetAction(pattern string) func(writer http.ResponseWriter, request *http.Request) {
	return c.actions[pattern]
}
