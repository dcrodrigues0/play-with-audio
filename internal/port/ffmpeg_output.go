package port

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func RecordAudio(ctx context.Context){
	cmd := exec.CommandContext(ctx, "ffmpeg",
    "-f", "alsa",
    "-i", "default",
    "-ar", "48000",            // Set sample rate
    "-ac", "2",                // Stereo
    "-sample_fmt", "s16",      // Sample format: 16-bit signed integer
    "-acodec", "pcm_s16le",    // WAV codec
    "-y",                      // Overwrite output file
    "output.wav",
	)
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Run()
	if err != nil{
		fmt.Println("Audio record stopped:",err)
	}
}

func ImproveAudio(){
	cmd := exec.Command("ffmpeg",
    "-i", "output.wav",
    "-af", "highpass=f=100, lowpass=f=3000, afftdn=nf=-25, loudnorm",
 		"-y",
		"clean.wav",
	)
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Run()
	if err != nil{
		fmt.Println("Audio improvement err:",err)
	}
}

func GetAvailableMics(){
	//TODO Implement a way to list available mics
	cmd := exec.Command("arecord","-l")
	err := cmd.Run()
	if err != nil{
		fmt.Println("Audio improvement err:",err)
	}
}





