package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Service interface {
	Name() string
}

func initStub(service Service) []reflect.Value {
	var fns []reflect.Value

	if reflect.TypeOf(service).Kind() != reflect.Ptr {
		panic("unsupported type")
	}

	val := reflect.ValueOf(service).Elem()
	typ := reflect.TypeOf(service).Elem()
	numField := val.NumField()

	for i := 0; i < numField; i++ {
		fieldVal := val.Field(i)
		fieldTyp := typ.Field(i)

		if !fieldVal.CanSet() {
			log.Println("func is not exported")
			continue
		}
		if fieldTyp.Type.Kind() != reflect.Func {
			log.Println("field must be a func")
			continue
		}

		fn := reflect.MakeFunc(fieldTyp.Type, func(args []reflect.Value) (results []reflect.Value) {
			if _, ok := args[0].Interface().(context.Context); !ok {
				panic("first arg must be context.Context")
			}
			arg := args[1].Interface()
			outTyp := fieldTyp.Type.Out(0)

			// TODO
			// First step: encode arg
			// Second step: call client's Invoke func

			// we just print arg and outType
			fmt.Println(arg)
			fmt.Println(outTyp)

			first := reflect.New(outTyp.Elem()).Interface()
			second := reflect.ValueOf(errors.New("nil"))
			results = append(results, reflect.ValueOf(first), second)

			return results
		})
		fieldVal.Set(fn)

		fns = append(fns, fn)
	}

	return fns
}
