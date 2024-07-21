package codecs

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/vo"
)

const (
	outputDir = "E:/test/output"
	extension = ".mp4"
)

func Concat(task vo.TaskMsg) error {
	folderDir := filepath.Join(outputDir, strconv.Itoa(task.JobId))

	var outputs []vo.VideoConcat
	if err := json.Unmarshal(task.Outputs, &outputs); err != nil {
		return err
	}

	for _, output := range outputs {
		// Performing stiching of video files
		outFile := output.Height + "@" + output.Fps + extension
		encoder := cmd.GetFfmpeg().ConcatVideo(concatFile).Output("../" + outFile)
		fmt.Println(encoder.GetCmd())
		err := encoder.RunInDir(folderDir)
		if err != nil {
			return err
		}
	}

	return nil
}
