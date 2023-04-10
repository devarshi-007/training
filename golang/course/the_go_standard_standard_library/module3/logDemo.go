package module3

import (
	"log"
	"os"
)

func logDemo() {
	log.Println("This is log file")

	file, err := os.OpenFile("module3/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("INFO: This is a message")

	output := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	output.Println("this is a message which is root cause of error")

	output = log.New(file, "Fatal: ", log.Ldate|log.Ltime|log.Lshortfile)
	output.Fatal("this is a message which is root cause of error")

}
