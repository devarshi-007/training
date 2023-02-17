package main

import (
	"errors"
	"fmt"
)

func func_with_error() {
	call, calle := func_error()
	fmt.Println(call, calle)
}

func func_error() (error, error) {
	return errors.New("Hi, I am your dear error!"), nil
}
