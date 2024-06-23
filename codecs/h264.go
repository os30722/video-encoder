package codecs

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/vo"
)

const video_codec = "h264_qsv"

func RunH264(msg vo.TaskMsg) error {
	opts := msg.Options
	fps, _ := strconv.Atoi(opts.Fps)

	encoder := cmd.GetFfmpeg().Qsv().Input(msg.InputDir).
		Async().
		Codec(video_codec).
		Scale(opts.Width, opts.Height).
		Profile(opts.Profile).
		VRate(opts.Fps).
		Gop(strconv.Itoa(fps * 2)).
		MaxRate(opts.MaxBitRate).
		MinRate(opts.MinBitRate).
		BuffSize(opts.BuffSize).
		VBitRate(opts.Bitrate).
		NoAudio().Overwrite().Output(msg.OutputDir)

	fmt.Println(encoder.GetCmd())

	if err := encoder.Run(); err != nil {
		log.Printf("Error in running ffmepg (h264) => %s", err)
		panic(err)
		return err
	}

	return nil
}
