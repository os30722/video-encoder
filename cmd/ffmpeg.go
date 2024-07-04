package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

type ffmpeg struct {
	cmd     string
	options []string
}

func GetFfmpeg() *ffmpeg {
	return &ffmpeg{
		cmd:     "ffmpeg",
		options: make([]string, 0),
	}
}

func (f ffmpeg) GetCmd() string {
	return f.cmd + " " + strings.Join(f.options[:], " ")
}

func (f *ffmpeg) append(opts ...string) {
	f.options = append(f.options, opts...)
}

func (f *ffmpeg) AddOptions(key string, value string) *ffmpeg {
	if len(value) == 0 {
		return f
	}

	f.append("-"+key, value)
	return f
}

func (f *ffmpeg) Qsv() *ffmpeg {
	f.append("-init_hw_device", "qsv=qsv:hw", "-filter_hw_device", "qsv")
	return f
}

func (f *ffmpeg) Async() *ffmpeg {
	f.AddOptions("threads", "1")
	return f

}

func (f *ffmpeg) Overwrite() *ffmpeg {
	f.append("-y")
	return f
}

func (f *ffmpeg) Input(file string) *ffmpeg {
	f.AddOptions("i", file)
	return f
}

func (f *ffmpeg) VCodec(codec string) *ffmpeg {
	f.AddOptions("c:v", codec)
	return f
}

func (f *ffmpeg) ACodec(codec string) *ffmpeg {
	f.AddOptions("c:a", codec)
	return f
}

func (f *ffmpeg) Scale(width string, height string) *ffmpeg {
	f.AddOptions("vf", fmt.Sprintf("scale=%s:%s", width, height))
	return f
}

func (f *ffmpeg) Profile(profile string) *ffmpeg {
	f.AddOptions("profile:v", profile)
	return f
}

func (f *ffmpeg) VRate(rate string) *ffmpeg {
	f.AddOptions("r", rate)
	return f
}

func (f *ffmpeg) ARate(rate string) *ffmpeg {
	f.AddOptions("ar", rate)
	return f
}

func (f *ffmpeg) Gop(gop string) *ffmpeg {
	f.AddOptions("g", gop)
	return f
}

func (f *ffmpeg) Format(format string) *ffmpeg {
	f.AddOptions("vf format=", format)
	return f
}

func (f *ffmpeg) MinRate(rate string) *ffmpeg {
	f.AddOptions("minrate", rate)
	return f
}

func (f *ffmpeg) MaxRate(rate string) *ffmpeg {
	f.AddOptions("maxrate", rate)
	return f
}

func (f *ffmpeg) BuffSize(size string) *ffmpeg {
	f.AddOptions("bufsize", size)
	return f
}

func (f *ffmpeg) VBitRate(rate string) *ffmpeg {
	f.AddOptions("b:v", rate)
	return f
}

func (f *ffmpeg) ABitRate(rate string) *ffmpeg {
	f.AddOptions("b:a", rate)
	return f
}

func (f *ffmpeg) NoAudio() *ffmpeg {
	f.append("-an")
	return f
}

func (f *ffmpeg) SplitVideo(segmentTime string, concatFile string, segmentDir string) *ffmpeg {
	f.append("-c", "copy", "-map", "v:0", "-segment_time", segmentTime, "-reset_timestamps", "1", "-f",
		"segment", "-segment_list", filepath.Join(segmentDir, concatFile))
	return f
}

func (f *ffmpeg) SplitAudio() *ffmpeg {
	f.AddOptions("c:a", "copy")
	f.AddOptions("map", "a:0")
	return f
}

func (f *ffmpeg) Output(out string) *ffmpeg {
	f.append(out)
	return f
}

func (f ffmpeg) RunInDir(dir string) error {
	cmd := exec.Command(f.cmd, f.options...)
	if dir != "" {
		cmd.Dir = dir
	}
	info, err := cmd.CombinedOutput()
	log.Print(string(info))
	if err != nil {
		return err
	}

	return nil
}

func (f ffmpeg) Run() error {
	return f.RunInDir("")
}
