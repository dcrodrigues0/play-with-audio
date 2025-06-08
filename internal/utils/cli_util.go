package utils

import(
	"fmt"
	"bufio"
	"os"
	"context"
)

type InputCallback func(string)

func AskInput(inputMsg string, cancel CancelCauseFunc) string{
	fmt.Println(inputMsg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
	userInput := scanner.Text()
	//TODO Fix bug audio at stop audio
	return userInput
}
