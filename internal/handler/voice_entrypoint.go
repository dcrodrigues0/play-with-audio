package handler
import (
	"strings"
	"os"
	"fmt"
	"play-with-audio/internal/service"
	"play-with-audio/internal/utils"
)

func CaptureAudio() {
	userWantsAudioRec()
	service.AudioRec()
}

func userWantsAudioRec() {
	userInput := utils.AskInput("Start audio record? yes/no")

	if strings.ToLower(strings.TrimSpace(userInput)) != "yes" {
		fmt.Println("Closing app...")
		os.Exit(0)
	}
}

