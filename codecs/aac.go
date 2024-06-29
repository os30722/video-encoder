package codecs

import (
	"fmt"
	"log"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/vo"
)

const audio_codec = "aac"

var aac = struct {
	BitRate    string
	SampleRate string
}{
	BitRate:    "Bitrate",
	SampleRate: "SampleRate",
}

func RunAac(msg vo.TaskMsg) error {
	output := msg.Output
	opts := output.Options

	encoder := cmd.GetFfmpeg().Input(msg.InputDir).
		Async().
		Codec(audio_codec).
		ABitRate(opts[aac.BitRate]).
		ARate(opts[aac.SampleRate]).
		Output(msg.OutputDir)

	fmt.Println(encoder.GetCmd())

	if err := encoder.Run(); err != nil {
		log.Printf("Error in running ffmepg => %s", err)
		panic(err)
		return err
	}

	return nil
}
