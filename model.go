package main

type RootData struct {
	Ff int `json:"ff" validate:"required,numeric,min=10,max=100"`
}
