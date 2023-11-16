package cmd

import (
	"fmt"
	"reflect"
	"runtime/debug"
)

func Goroutine(method interface{}, args ...interface{}) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				_ = fmt.Sprintf("panic: %v\n%v", r, string(debug.Stack()))
				//log.Err("[Goroutine][panic] %s", msg)
			}
		}()

		if reflect.TypeOf(method).Kind() != reflect.Func {
			//log.Err("[Goroutine][error] method must function")

			return
		}

		vArgs := make([]reflect.Value, len(args))
		for i, val := range args {
			vArgs[i] = reflect.ValueOf(val)
		}

		reflect.ValueOf(method).Call(vArgs)
	}()
}
