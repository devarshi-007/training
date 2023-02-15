package module5

import (
	"fmt"
	"strings"
)

func stringOpr() {

	str := "\x52\x60\x4a\x20\x45"

	fmt.Println(str)

	// for _, i := range str {
	// 	fmt.Printf("%q %x\n\n", i, i)
	// }

	fmt.Printf("\nstr[2]: %x \nstring: %s", str[2], strings.ToLower(str[:]))

	str1 := "this is something"
	str2 := "this is something."
	str3 := str2

	fmt.Printf("\nstr1 == str2 : %t, str2 == str3: %t, str1 == str3: %t",
		str1 == str2,
		str2 == str3,
		str1 == str3)

	fmt.Printf("\nstr1 == str2 : %d %v %v, \nstr2 == str3: %d %v %v, \nstr1 == str3: %d %v %v",
		strings.Compare(str1, str2), &str1, &str2,
		strings.Compare(str2, str3), &str2, &str3,
		strings.Compare(str1, str3), &str1, &str3,
	)

	println("\n", strings.Index(str1, "is"))

	println(strings.Join(strings.Split(str1, " "), " + "))

}
