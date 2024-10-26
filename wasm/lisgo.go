package main

import (
	"lisgo/pkg/lisgo"
	"syscall/js"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func lisgoJsEval(this js.Value, args []js.Value) interface{} {
	return lisgo.Eval(args[0].String()).ToString()
}

func main() {
	js.Global().Set("lisgo", js.FuncOf(lisgoJsEval))
	<-c
}
