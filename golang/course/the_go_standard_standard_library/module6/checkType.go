package module6

import (
	"fmt"
	"main/module6/dataclass"
	"reflect"
)

func checkType() {
	interFaceType(5)
	interFaceType("hii")
	interFaceType([]int{})
}

func interFaceType(t interface{}) {
	fmt.Println(reflect.TypeOf(t).Name())
}

func interfaceOpr() {
	// myFav := dataclass.GetMovie("abc", dataclass.D, 55.3)
	// fmt.Println(myFav.GetTitle())
	// myFav.SetTitle("xyz")
	// fmt.Println(myFav.GetTitle())

	//dataclass.Movie{}
	var myFav dataclass.MetaData = &dataclass.Movie{}
	myFav.GetMovie("defTitle", dataclass.A)
	fmt.Println(myFav.GetTitle())
	myFav.SetTitle("xyz")
	fmt.Println(myFav.GetTitle())

}
