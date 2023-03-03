package module1

import "fmt"

type TryCatchBlock struct {
	Try     func()
	Catch   func()
	Finally func()
}

func (tc TryCatchBlock) Handle() {
	defer tc.Finally()
	defer func() {
		if r := recover(); r != nil {
			tc.Catch()
		} else {
			fmt.Println("no error occured", r)
		}

	}()
	tc.Try()
}

func errorHandling() {
	divideit := 5
	divideby := 0
	TryCatchBlock{
		Try: func() {
			ans, err := divideBy(divideit, 3)
			if err != nil {
				fmt.Println(ans, err.Error(), ans)
			}
			ans, err = divideBy(divideit, divideby)
			fmt.Println(ans, err)
			fmt.Println("This will not execute")
			//ans, err = divideBy(6, divideby)
			//fmt.Println(ans, err.Error())
		},
		Catch: func() {
			fmt.Println("error handled 1")
		},
		Finally: func() {
			fmt.Println("Finally also executed")
		},
	}.Handle()

	TryCatchBlock{
		Try: func() {
			ans, err := divideBy(divideit, 3)
			if err != nil {
				fmt.Println(ans, err.Error())
			}
			// ans, _ = divideBy(divideit, divideby)
			// fmt.Println(ans)
			fmt.Println("This will not execute")
			//ans, err = divideBy(6, divideby)
			//fmt.Println(ans, err.Error())
		},
		Catch: func() {
			fmt.Println("error handled 1")
		},
		Finally: func() {
			fmt.Println("Finally also executed")
		},
	}.Handle()

	fmt.Println("after error")
}

func divideBy(x, y int) (int, error) {
	return x / y, nil
}
