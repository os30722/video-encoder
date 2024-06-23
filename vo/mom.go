package vo

type TaskMsg struct {
	JobId     int
	InputDir  string
	OutputDir string
	Options   VideoEncodeOption
}

type ConcatMSg struct {
	JobId int
}
