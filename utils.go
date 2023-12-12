package main

import (
	"bufio"
	"fmt"
	"os"
)

func AcceptUserInput(promptMessage string) string {
	fmt.Println(promptMessage)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
		return ""
	}
	result := scanner.Text()
	fmt.Println("You entered: ", result)

	return result
}
