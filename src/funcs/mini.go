package funcs

import (
	"reflect"
	"runtime"
	"strings"
)

func Indent(count int) string {
	return strings.Repeat(" ", count)
}

func FuncName(fn interface{}) string {
	funcAddr := reflect.ValueOf(fn).Pointer()
	fullPathFunc := runtime.FuncForPC(funcAddr)
	strSlice := strings.Split(fullPathFunc.Name(), ".")
	return strSlice[len(strSlice)-1]
}

func IsFunc(fn interface{}, funcName string) bool {
	return FuncName(fn) == funcName
}
