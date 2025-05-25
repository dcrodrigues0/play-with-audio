package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CaptureAudio() {
	userWantsAudioCap()
	//TODO Implement a way to capture audio bits, idk how but i'll discover
}

func userWantsAudioCap() {
	fmt.Println("Should i capture audio? yes/no")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	if strings.ToLower(strings.TrimSpace(userInput)) != "yes" {
		fmt.Println("Closing app...")
		os.Exit(0)
	}
}
