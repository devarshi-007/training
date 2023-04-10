package module5

import (
	"fmt"
	"regexp"
)

func findAndReplaceWithRe() {

	str := "This is a string. the string is not so long. String is data-type in go"

	r, _ := regexp.Compile(`s([a-z]+)g`)

	// fmt.Println(r.MatchString(str))

	fmt.Println(r.FindString(str), r.FindStringIndex(str))

	fmt.Println(r.FindAllString(str, -1), r.FindAllStringIndex(str, -1))

	newText := r.ReplaceAllString(str, "byte Array")

	fmt.Println(newText)
}
