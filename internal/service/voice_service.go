package service
import (
	"context"
	"play-with-audio/internal/port"
	"play-with-audio/internal/utils"
	"fmt"
	"os"
	"strconv"
)

func AudioRec(){
	ctx := context.Background()	
	ctx, cancel := context.WithCancel(ctx)
	recordOptions := port.GetAvailableMics()

	input := utils.AskInput(fmt.Sprintf("Select prefered card option to record (ex: 0 or 2): \n %s", recordOptions))
	
	s, err := strconv.Atoi(input)
	if err != nil{
		fmt.Println("Not valid option...")
		os.Exit(0)
	}

	go func() {
		utils.AskInput("Press enter to stop recording")
		cancel()
		port.ImproveAudio()
	}()

	port.RecordAudio(ctx,s)
}

