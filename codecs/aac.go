package codecs

import (
	"fmt"
	"log"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/vo"
)

const audio_codec = "aac"

func RunAac(msg vo.TaskMsg) error {
	opts := msg.Options

	encoder := cmd.GetFfmpeg().Input(msg.InputDir).
		Async().
		Codec(audio_codec).
		ABitRate(opts.Bitrate).
		ARate(opts.SampleRate).
		Output(msg.OutputDir)

	fmt.Println(encoder.GetCmd())

	if err := encoder.Run(); err != nil {
		log.Printf("Error in running ffmepg => %s", err)
		panic(err)
		return err
	}

	return nil
}
