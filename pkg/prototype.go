package main

import (
	"fmt"

	proc "github.com/mitchellh/go-ps"
)

func main() {
	fmt.Println("Portshield V2")
	var i int
	// fmt.Print("Enter port number : ")
	fmt.Scan(&i)

	//
	proccessList, err := proc.Processes()
	if err != nil {
		panic(err)
	}

	// linux command
	// lsof -i -P | grep -i LISTEN | grep {PID}
	//

	//check if there a port running
	// fmt.Print(proccessList)
	for x := range proccessList {
		var process proc.Process
		process = proccessList[x]
		// log.Printf("PID: %d, Name: %s\n", process.Pid(), process.Executable())
	}

	fmt.Println("No port like that", i)
}
