package lisgo

import "fmt"

func RuntimePrint(interpreter *Interpreter, expressions []Expression) LisgoData {
	var result LisgoData
	for _, expr := range expressions {
		result = interpreter.Evaluate(expr)
		fmt.Println(result.ToString())
	}
	return result
}
