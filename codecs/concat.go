package codecs

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/packager"
	"github.com/cloud/encoder/vo"
)

const (
	concatFile = "input.ffconcat"
	outputDir  = "E:/test/output"
	extension  = ".mp4"
)

func Concat(jobInfo *vo.JobFileAndOpts) error {
	folderDir := filepath.Join(outputDir, jobInfo.JobName)
	log.Println(jobInfo.Opts)

	vopts := jobInfo.Opts.Video

	// Performing stiching of video files
	for _, opt := range vopts {
		outFile := opt.Height + "@" + opt.Fps + extension
		cmdDir := filepath.Join(folderDir, opt.Height+"@"+opt.Fps)
		encoder := cmd.GetFfmpeg().ConcatVideo(concatFile).Output("../" + outFile)
		fmt.Println(encoder.GetCmd())
		err := encoder.RunInDir(cmdDir)
		if err != nil {
			return err
		}
	}

	for _, format := range jobInfo.OutputFormats {
		var err error
		switch format {
		case "dash":
			err = packager.OutputDash(jobInfo)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
