package module2

import (
	"fmt"
	"main/webService"
)

func structDataType() {

	var u webService.User
	u.ID = 1
	u.FirstName = "abc"
	u.LastName = "xyz"
	u.Done.IsAllGood = true
	fmt.Print(u)

	u2 := webService.User{
		ID:        2,
		FirstName: "fn",
		LastName:  "ln",
		Done:      webService.OurCustomClass{IsAllGood: true},
	}
	fmt.Println(u2, u2.ID)

}
