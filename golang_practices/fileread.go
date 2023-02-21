package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// )

// func fileread() {
// 	file, err := os.open("demo.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		var id int
// 		var name string

// 		n, err := fmt.Sscanf(scanner.Text(), "here %d is %s", &id, &name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if n == 2 {
// 			fmt.Printf("%d,%s\n", id, name)
// 		}
// 	}
