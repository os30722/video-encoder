package vo

import "encoding/json"

type TaskMsgHolder struct {
	JobId     int
	InputDir  string
	File      string
	OutputDir string
	Codec     string
	Type      string
	Outputs   json.RawMessage
}
