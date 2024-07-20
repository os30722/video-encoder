package codecs

const audio_codec = "aac"

var aac = struct {
	BitRate    string
	SampleRate string
}{
	BitRate:    "Bitrate",
	SampleRate: "SampleRate",
}

// func RunAac(msg vo.TaskMsg) error {
// 	output := msg.Output
// 	opts := output.Options

// 	encoder := cmd.GetFfmpeg().Async().Input(msg.InputDir).
// 		ACodec(audio_codec).
// 		ABitRate(opts[aac.BitRate]).
// 		ARate(opts[aac.SampleRate]).
// 		Overwrite().
// 		Async().
// 		Output(msg.OutputDir)

// 	fmt.Println(encoder.GetCmd())

// 	if err := encoder.Run(); err != nil {
// 		log.Printf("Error in running ffmepg => %s", err)
// 		panic(err)
// 		return err
// 	}

// 	return nil
// }
