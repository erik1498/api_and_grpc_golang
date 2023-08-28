package utils

import (
	"runtime"
	"strings"
)

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	functionName := strings.Split(name, ".")
	return functionName[len(functionName)-1]
}
