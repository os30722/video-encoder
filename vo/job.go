package vo

type EncodeOuputStruct struct {
	Video []EncodeOutput
	Audio EncodeOutput
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
