package service

import (
	"context"
	"play-with-audio/internal/port"
	"play-with-audio/internal/utils"
)

func AudioRec(){
	ctx := context.Background()	
	ctx, cancel := context.WithCancel(ctx)
	port.GetAvailableMics() //Testing
	go func() {
		utils.AskInput("Press enter to stop recording")
		cancel()
		port.ImproveAudio()
	}()

	port.RecordAudio(ctx)
}

