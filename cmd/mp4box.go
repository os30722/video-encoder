package cmd

import (
	"os/exec"
	"strings"
)

type mp4box struct {
	cmd     string
	options []string
}

func GetMp4box() *mp4box {
	return &mp4box{
		cmd:     "mp4box",
		options: make([]string, 0),
	}
}

func (m mp4box) GetCmd() string {
	return m.cmd + " " + strings.Join(m.options[:], " ")
}

func (m *mp4box) append(opts ...string) {
	m.options = append(m.options, opts...)
}

func (m *mp4box) AddOptions(key string, value string) *mp4box {
	if len(value) == 0 {
		return m
	}

	m.append("-"+key, value)
	return m
}

func (m *mp4box) GenerateDash(time string, fps string, inputStr []string, profile string, output string) *mp4box {
	m.AddOptions("dash", time)
	m.append("-rap")
	m.AddOptions("fps", fps)
	for _, inp := range inputStr {
		m.append(inp)
	}
	m.AddOptions("profile", profile)
	m.AddOptions("out", output)
	return m
}

func (m mp4box) RunInDir(dir string) error {
	cmd := exec.Command(m.cmd, m.options...)
	if dir != "" {
		cmd.Dir = dir
	}
	_, err := cmd.CombinedOutput()
	// log.Print(string(info))
	if err != nil {
		return err
	}

	return nil
}

func (m mp4box) Run() error {
	return m.RunInDir("")
}
