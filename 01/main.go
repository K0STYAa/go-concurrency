/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import (
	"fmt"
)

// https://go.dev/doc/effective_go#embedding

type human struct {
	name string
}

func (h *human) setName(str string) {
	h.name = str
}

type dummy interface {
	dummyFunc()
}

// Action композиция + встраивание
// Благодаря встраиванию, поля и методы human доступны прямо из action, без обращеия к human (a.human.name = "????")
type action struct {
	human
	someField int
	//*log.Logger

	// можно встроить интерфейс... но зачем???????????????
	// Only interfaces can be embedded within interfaces.
	// dummy
}

func (a *action) dummyFunc() {
	fmt.Println("Beautiful, realistic, uncannily convincing human replica.")
}

func (a *action) sayMeow() {
	fmt.Printf("%s said oaaoooaoaoaoaoao.\n", a.name)
}

func main() {
	var a action

	a.setName("Vasya")
	a.sayMeow()

	// panic: runtime error: invalid memory address or nil pointer dereference
	// a.dummy.dummyFunc()

	a.dummyFunc()
}
