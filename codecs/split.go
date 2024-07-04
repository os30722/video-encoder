package codecs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cloud/encoder/cmd"
)

const (
	concatFile        = "input.ffconcat"
	videoOutputFormat = "out_%d.mp4"
	AudioOutputFormat = "audio.m4a"
	segmentTime       = "00:04:00"
)

func SplitVideo(inputDir string, outputDir string) error {
	os.Mkdir(outputDir, 0777)

	encoder := cmd.GetFfmpeg().Input(inputDir).Overwrite().
		SplitVideo(segmentTime, concatFile, outputDir).
		Output(filepath.Join(outputDir, videoOutputFormat)).
		SplitAudio().
		Output(filepath.Join(outputDir, AudioOutputFormat))

	fmt.Println(encoder.GetCmd())

	if err := encoder.Run(); err != nil {
		log.Printf("Error in running ffmepg => %s", err)
		return err
	}

	return nil
}
