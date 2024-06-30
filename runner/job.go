package runner

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cloud/encoder/codecs"
	"github.com/cloud/encoder/mom"
	"github.com/cloud/encoder/repository/jobDb"
	"github.com/cloud/encoder/vo"
)

func SubmitJob(ctx context.Context, msg vo.TaskMsg, jobDao jobDb.JobRepo) error {
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
		return err
	}

	vopts := outputs.Video
	aopts := outputs.Audio

	numOfParts := len(files) - 2

	task := vo.TaskMsg{
		JobId:     jobId,
		InputDir:  filepath.Join(outputDir, codecs.AudioOutputFormat),
		OutputDir: filepath.Join(outputDir, "encoded_"+codecs.AudioOutputFormat),
		Codec:     aopts.Codec,
		Output:    aopts,
	}

	if err = mom.PublishTask(ctx, task); err != nil {
		return err
	}

	processes := make([]vo.Process, 0, len(vopts))

	for _, output := range vopts {
		partName := output.Height + "@" + output.Fps
		path := filepath.Join(outputDir, partName)
		os.Mkdir(path, 0777)

		for _, file := range files {
			if !(strings.HasPrefix(file, "out")) {
				continue
			}

			task := vo.TaskMsg{
				JobId:     jobId,
				InputDir:  filepath.Join(outputDir, file),
				OutputDir: filepath.Join(path, file),
				Codec:     output.Codec,
				Output:    output,
			}

			if err = mom.PublishTask(ctx, task); err != nil {
				return err
			}

			data, err := os.ReadFile(filepath.Join(outputDir, "input.ffconcat"))
			if err != nil {
				return err
			}

			err = os.WriteFile(filepath.Join(path, "input.ffconcat"), data, 0777)
			if err != nil {
				return err
			}
		}

		process := vo.Process{
			JobId:     jobId,
			PartName:  partName,
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
