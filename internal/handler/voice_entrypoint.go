package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"context"
	"play-with-audio/internal/port"	
)

func CaptureAudio() {
	userWantsAudioRec()
	audioRec()
}

func userWantsAudioRec() {
	fmt.Println("Should i record audio? yes/no")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	
	if strings.ToLower(strings.TrimSpace(userInput)) != "yes" {
		fmt.Println("Closing app...")
		os.Exit(0)
	}
}

func audioRec(){
	ctx := context.Background()	
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("Press enter to stop recording...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
	port.RecordAudio(ctx)	
	defer cancel()
}

