package cmd

func (f *ffmpeg) ConcatVideo(concatFile string) *ffmpeg {
	f.AddOptions("f", "concat")
	f.Input(concatFile)
	f.VCodec("copy")

	return f
}
