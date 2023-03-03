package module5

import (
	"fmt"
	"strings"
)

func findAndReplace() {

	str := "this is something to replace not only once but twise, replace with find, in turms of reset re is comman prefix with replace"

	replaceTo := "replace"
	replaceWith := "find"

	fmt.Println(strings.Contains(str, replaceTo))

	fmt.Println(strings.HasPrefix(str, "the"), strings.Count(str, replaceTo))

	fmt.Println(strings.Contains(str, replaceTo))

	newString := strings.Replace(str, replaceTo, replaceWith, 1)

	println(newString)

	newString = strings.Replace(newString, replaceTo, replaceWith, -1)

	println(newString)

	newString = strings.ReplaceAll(newString, replaceWith, replaceTo)

	println(newString)

}
