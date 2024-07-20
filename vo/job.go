package vo

type EncodeOutputStruct struct {
	Video []EncodeOutput
	Audio EncodeOutput
}

type JobTemplate struct {
	Streams []string
	Outputs EncodeOutputStruct
}

type EncodeOutput struct {
	Codec   string
	Options []interface{}
}

type VideoH264 struct {
	Width      string
	Height     string
	Fps        string
	Bitrate    string
	Profile    string
	MaxBitRate string
	MinBitRate string
	BuffSize   string
}

type VideoConcat struct {
	Height string
	Fps    string
}

type Process struct {
	JobId     int
	PartName  string
	TotalPart int
}
