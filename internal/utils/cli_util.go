package utils

import(
	"fmt"
	"bufio"
	"os"
)

func AskInput(inputMsg string) string{
	fmt.Println(inputMsg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
	userInput := scanner.Text()
	return userInput
}
