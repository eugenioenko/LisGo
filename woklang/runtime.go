package woklang

var RuntimeScope = map[string]WokData{
	"print":       WF("print", RuntimePrint),
	"cond":        WF("cond", RuntimeCond),
	"while":       WF("while", RuntimeWhile),
	"debug":       WF("debug", RuntimeDebug),
	"defun":       WF("defun", RuntimeDefun),
	"return-from": WF("return-from", RuntimeReturnFrom),
	"if":          WF("if", RuntimeIf),
	":=":          WF(":=", RuntimeAssignment),
	"==":          WF("==", RuntimeEquality),
	"!":           WF("!", RuntimeNegation),
	"+":           WF("+", RuntimeAddition),
	"*":           WF("*", RuntimeMultiplication),
	"-":           WF("-", RuntimeSubstraction),
}
