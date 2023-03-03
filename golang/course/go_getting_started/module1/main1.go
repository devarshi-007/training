package module1

// all required imports
import (
	"fmt"
	"strconv"
)

/*
main is a driver of this file
*/
func variableDemo() {
	var i uint
	i = 50
	fmt.Println(i)

	var j uint8 = 11
	fmt.Print(uint8(i) + j)

	var k bool = true

	hello := "hello"

	fmt.Println("\nsomething " + hello)

	i32 := int32(i)
	j32 := int32(j)

	realAndImagenery := complex(float32(i32), float32(j32))
	fmt.Printf("%g %v %.2f as real and %.2f as imag, %t as %T\n",
		realAndImagenery, realAndImagenery, real(realAndImagenery), imag(realAndImagenery), k, k)

	var invalidInt uint8 = 254
	var half int8 = 127
	newone := int8(invalidInt)
	fmt.Println("\n", invalidInt, newone, int8(half), strconv.FormatInt(int64(newone), 2))
}
