//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/eugenioenko/lisgo/pkg/lisgo"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func lisgoJsEval(this js.Value, args []js.Value) any {
	return lisgo.Eval(args[0].String()).ToString()
}

func main() {
	js.Global().Set("lisgo", js.FuncOf(lisgoJsEval))
	<-c
}
