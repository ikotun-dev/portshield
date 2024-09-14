package main

import (
	"fmt"
	"os/exec"
)

func main_0() {
	fmt.Println("LSOF Command Test")

	var port int
	fmt.Print("Enter port number to check: ")
	fmt.Scan(&port)

	// Construct the lsof command
	cmd := exec.Command("bash", "-c", fmt.Sprintf("lsof -i -P -n | grep LISTEN | grep :%d", port))

	// Execute the command and capture output
	output, err := cmd.CombinedOutput()

	// Print the raw output and error (if any)
	fmt.Println("Command output:")
	fmt.Println(string(output))

	if err != nil {
		fmt.Println("Error:", err)
	}

}
