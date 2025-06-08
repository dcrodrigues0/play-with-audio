package port

import (
	"context"
	"fmt"
	"os/exec"
)

func RecordAudio(ctx context.Context){
	err := exec.CommandContext(ctx,"ffmpeg", "-f", "alsa", "-i", "default", "-y", "output.wav").Run()
	if err != nil{
		fmt.Println("Não foi possível capturar o audio:",err)
	}
}
