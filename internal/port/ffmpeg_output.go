package port

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"bytes"
	"play-with-audio/internal/models"
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
	output, err := exec.Command("arecord","-l").Output()
	
	if err != nil{
		fmt.Println("Could not get available mic list: ",err)
	}

	rawList := bytes.Split(output, []byte("\n"))
	var micList []Arecord
	for _,line := range rawList{
		if bytes.Contains(line,[]byte("card")){
			//TODO Rescue card id, device name and type from arecord -l
			micList = append(models.NewArecord(), line)	
		}	
	}

	fmt.Println(string(micList[0]) + "\n" + string(micList[1]))
}







