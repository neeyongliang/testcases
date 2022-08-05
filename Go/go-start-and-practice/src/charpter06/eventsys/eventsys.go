package main

import (
	"fmt"
)

var eventByName = make(map[string][]func(interface{}))

func registerEvent(name string, callback func(interface{})) {
	list := eventByName[name]
	list = append(list, callback)
	eventByName[name] = list
}

func callEvent(name string, param interface{}) {
	list := eventByName[name]
	for _, callback := range list {
		callback(param)
	}
}

type Actor struct {
}

func (a *Actor) onEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

func globalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	a := new(Actor)
	registerEvent("On Skill", a.onEvent)
	registerEvent("On Skill", globalEvent)
	callEvent("On Skill", 100)
}
