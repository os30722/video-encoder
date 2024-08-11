package runner

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cloud/encoder/codecs"
	"github.com/cloud/encoder/mom"
	"github.com/cloud/encoder/repository/jobDb"
	"github.com/cloud/encoder/vo"
)

const (
	jobOutputDir = "E:/test/output"
)

func SubmitJob(ctx context.Context, msg vo.TaskMsgHolder, jobDao jobDb.JobRepo) error {
	inputDir := msg.InputDir

	jobId, err := jobDao.CreateJob(ctx, msg.JobId)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(msg.OutputDir, strconv.Itoa(jobId))

	err = codecs.SplitVideo(inputDir, outputDir)
	if err != nil {
		return err
	}

	dir, err := os.Open(outputDir)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdirnames(-10)
	if err != nil {
		return err
	}

	outputs, err := jobDao.GetOutputs(ctx, msg.JobId)
	if err != nil {
		log.Print("deded")
		return err
	}

	numOfParts := len(files) - 2

	vopts := outputs.Video

	processes := make([]vo.Process, 0, len(vopts))

	for _, output := range vopts {

		for _, file := range files {
			if !(strings.HasPrefix(file, "out")) {
				continue
			}

			out, err := json.Marshal(output.Options)
			if err != nil {
				return err
			}

			task := vo.TaskMsgHolder{
				JobId:     jobId,
				InputDir:  outputDir,
				OutputDir: filepath.Join(jobOutputDir, strconv.Itoa(jobId)),
				File:      file,
				Type:      "video",
				Codec:     output.Codec,
				Outputs:   out,
			}

			if err = mom.PublishTask(ctx, task); err != nil {
				return err
			}
		}

		process := vo.Process{
			JobId:     jobId,
			PartName:  output.Codec,
			TotalPart: numOfParts,
		}

		processes = append(processes, process)

	}
	err = jobDao.UpdateProcesses(ctx, jobId, processes)
	if err != nil {
		return err
	}

	return nil
}
