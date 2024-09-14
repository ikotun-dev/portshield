package main

import (
	"fmt"
	"log"
	"os/exec"

	proc "github.com/mitchellh/go-ps"
)

func main_3() {
	fmt.Println("Portshield V2")
	var i int

	proccessList, err := proc.Processes()
	if err != nil {
		panic(err)
	}

	fmt.Print(proccessList)
	for x := range proccessList {
		var process proc.Process
		process = proccessList[x]
		log.Printf("PID: %d, Name: %s\n", process.Pid(), process.Executable())
	}

	var pid int

	fmt.Print("Enter PID to check: ")
	fmt.Scan(&pid)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("lsof -i -P | grep -i LISTEN | grep %d", pid))

	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)

	}
	log.Printf("Output: %s\n", output)

	fmt.Println("No port like that", i)
}
