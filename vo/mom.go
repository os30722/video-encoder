package vo

type TaskMsg struct {
	JobId     int
	InputDir  string
	OutputDir string
	Codec     string
	Output    EncodeOutput
}
