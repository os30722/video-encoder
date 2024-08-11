package vo

import "encoding/json"

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
	Options json.RawMessage
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
