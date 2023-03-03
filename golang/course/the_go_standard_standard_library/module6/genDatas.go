package module6

import (
	"fmt"
	"reflect"
)

type typeA struct {
	likes int
}

type typeB struct {
	rates int
}

func genDatas() {
	aobj := typeA{likes: 5}
	bobj := typeB{rates: 10}

	fmt.Printf("The type of aobj is %v_ %T_ %v_ %v_ %v_",
		reflect.TypeOf(aobj),
		aobj,
		reflect.ValueOf(aobj),
		reflect.ValueOf(aobj).Kind(),
		reflect.TypeOf(aobj).Name())
	fmt.Printf("%v", bobj)

	lists := make([]typeA, 3)

	typea := reflect.MakeSlice(reflect.TypeOf(lists), 0, 0)

	fmt.Println('\n', typea)

	addPerson(typeA{5})
	addPerson(typeB{500})

}

func addPerson(p any) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		v := reflect.ValueOf(p)

		switch reflect.TypeOf(p).Name() {
		case "typeA":
			fmt.Printf("SQL values: %v\n", v.Field(0))
		case "typeB":
			fmt.Printf("SQL value: %v\n", v.Field(0))
		}
		return true
	} else {
		return false
	}
}
