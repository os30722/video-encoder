package vo

type JobFileAndOpts struct {
	JobName       string
	FileName      string
	OutputFormats []string
	Opts          EncodeOptions
}

type EncodeOptions struct {
	Video []VideoEncodeOption
	Audio AudioEncodeOption
}

type VideoEncodeOption struct {
	Codec      string
	Width      string
	Height     string
	Bitrate    string
	Fps        string
	Profile    string
	MaxBitRate string
	MinBitRate string
	BuffSize   string
	SampleRate string
}

type AudioEncodeOption struct {
	Codec      string
	BitRate    string
	SampleRate string
}
