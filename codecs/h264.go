package codecs

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/vo"
)

const video_codec = "h264_qsv"

var h264 = struct {
	Bitrate    string
	Profile    string
	MaxBitRate string
	MinBitRate string
	BuffSize   string
}{
	Bitrate:    "Bitrate",
	Profile:    "Profile",
	MaxBitRate: "MaxBitRate",
	MinBitRate: "MinBitRate",
	BuffSize:   "BuffSize",
}

func RunH264(msg vo.TaskMsg) error {
	output := msg.Output
	opts := output.Options

	fps, _ := strconv.Atoi(output.Fps)

	encoder := cmd.GetFfmpeg().Qsv().Async().Input(msg.InputDir).
		VCodec(video_codec).
		Scale(output.Width, output.Height).
		Profile(opts[h264.Profile]).
		VRate(output.Fps).
		Gop(strconv.Itoa(fps * 2)).
		MaxRate(opts[h264.MaxBitRate]).
		MinRate(opts[h264.MinBitRate]).
		BuffSize(opts[h264.BuffSize]).
		VBitRate(opts[h264.Bitrate]).
		NoAudio().Overwrite().Async().Output(msg.OutputDir)

	fmt.Println(encoder.GetCmd())

	if err := encoder.Run(); err != nil {
		log.Printf("Error in running ffmepg (h264) => %s", err)
		panic(err)
		return err
	}

	return nil
}
