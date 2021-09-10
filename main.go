package main

func main() {
	controller := NewController()
	SetActions(controller)
	StartServe("", 1111, controller)
}
