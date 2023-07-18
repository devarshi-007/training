package module6

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func createType() {
	emps := make([]typeA, 3)

	emps = append(emps, typeA{likes: 50})

	eType := reflect.TypeOf(emps)

	newEmps := reflect.MakeSlice(eType, 0, 0)

	newEmps = reflect.Append(newEmps, reflect.ValueOf(typeA{55}))

	fmt.Printf("First list of employees: %v _%v_\n", emps, reflect.TypeOf(emps).Name())

	fmt.Printf("First list of newEmps: %v _%v_\n", newEmps, reflect.TypeOf(newEmps).Name())

	str := "this is a string"

	timed := InterfaceFunc(properTitle).(func(string) string)
	newTitle := timed(str)
	fmt.Println(newTitle)

}

func properTitle(input string) string {
	return strings.ToTitle(input)
}

func InterfaceFunc(f interface{}) interface{} {
	rf := reflect.TypeOf(f)

	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}

	vf := reflect.ValueOf(f)

	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))

		return out
	})
	return wrapperF.Interface()
}
