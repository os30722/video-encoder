package codecs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func RunH264(msg vo.TaskMsg, options []vo.VideoH264) error {
	// output := msg.Output
	// opts := output.Options
	inputDir := msg.InputDir
	prevPart := ""

	for _, opts := range options {
		partName := opts.Height + "@" + opts.Fps

		path := filepath.Join(msg.OutputDir, partName)
		os.Mkdir(path, 0777)

		fps, _ := strconv.Atoi(opts.Fps)

		encoder := cmd.GetFfmpeg().Qsv().Async().Input(filepath.Join(inputDir, prevPart, msg.File)).
			VCodec(video_codec).
			Scale(opts.Width, opts.Height).
			Profile(opts.Profile).
			VRate(opts.Fps).
			Gop(strconv.Itoa(fps * 2)).
			MaxRate(opts.MaxBitRate).
			MinRate(opts.MinBitRate).
			BuffSize(opts.BuffSize).
			VBitRate(opts.Bitrate).
			NoAudio().Overwrite().Async().Output(filepath.Join(msg.OutputDir, partName, msg.File))

		fmt.Println(encoder.GetCmd())

		if err := encoder.Run(); err != nil {
			log.Printf("Error in running ffmepg (h264) => %s", err)
			panic(err)
			return err
		}

		prevPart = partName
	}

	return nil
}
