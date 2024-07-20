package vo

import "encoding/json"

type TaskMsg struct {
	JobId     int
	InputDir  string
	File      string
	OutputDir string
	Type      string
	Codec     string
	Outputs   []interface{}
}

type TaskMsgHolder struct {
	JobId     int
	InputDir  string
	File      string
	OutputDir string
	Codec     string
	Type      string
	Outputs   json.RawMessage
}
