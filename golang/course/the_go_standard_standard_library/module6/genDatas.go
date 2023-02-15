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

}
