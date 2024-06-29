package codecs

import (
	"fmt"
	"path/filepath"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/vo"
)

const (
	outputDir = "E:/test/output"
	extension = ".mp4"
)

func Concat(task vo.TaskMsg) error {
	folderDir := filepath.Dir(task.OutputDir)
	output := task.Output

	// Performing stiching of video files
	outFile := output.Height + "@" + output.Fps + extension
	encoder := cmd.GetFfmpeg().ConcatVideo(concatFile).Output("../" + outFile)
	fmt.Println(encoder.GetCmd())
	err := encoder.RunInDir(folderDir)
	if err != nil {
		return err
	}

	// for _, format := range jobInfo.OutputFormats {
	// 	var err error
	// 	switch format {
	// 	case "dash":
	// 		err = packager.OutputDash(jobInfo)
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
