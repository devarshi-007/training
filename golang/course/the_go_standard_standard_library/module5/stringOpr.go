package module5

import (
	"fmt"
	"regexp"
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

	str4 := `
	this is something
	hello
		remove extras
	`

	fmt.Printf("%T \n%s \n%s", str4, strings.TrimSpace(str4), strings.TrimSuffix(str4, "remove extras"))

	reg, _ := regexp.Compile(`([ \r\n\t\s])+`)

	fmt.Println(reg.ReplaceAllString(str4, " "))

}
