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
	Height  string
	Width   string
	Fps     string
	Options map[string]string
}

type Process struct {
	JobId     int
	PartName  string
	TotalPart int
}
