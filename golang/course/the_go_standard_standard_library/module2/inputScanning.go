package module2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func simpleInputScanning() {
	var firstName, lastName string

	fmt.Println("enter your [first_name last_name]: ")
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hii, %s %s\n", firstName, lastName)

	fmt.Scanln(&firstName, &lastName)
	fmt.Println(firstName, lastName, "*")

	file, err := os.Open("module2/fixedCsv.csv")
	if err != nil {
		log.Fatal(err)
	}

	scannner := bufio.NewScanner(file)

	for scannner.Scan() {
		var age int
		var name string

		n, err := fmt.Sscanf(scannner.Text(), "%s %d,\n", &name, &age)

		if err != nil {
			fmt.Print(err, " : ")
		}

		fmt.Printf("%s's age is %d... %d\n", name, age, n)
	}

}
