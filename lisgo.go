package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eugenioenko/lisgo/pkg/lisgo"
)

func main() {

	args := os.Args
	if len(args) <= 2 || args[1] == "help" {
		help()
		return
	}
	if args[1] == "exec" {
		lisgo.Exec(args[2])
		return
	}
	if args[1] == "eval" {
		lisgo.Eval(strings.Join(args[2:], " "))
		return
	}

}

func help() {
	fmt.Print(`
Lisgo is the interpreter of LisGo programming language.

Usage:

  lisgo command [arguments]

  The commands are:

      exec [filename]     executes the script with filename
      eval [code]         executes the code passed as argument
      help                prints this message

`)
}
