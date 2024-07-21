package packager

import (
	"context"
	"encoding/json"
	"log"
	"path/filepath"
	"strconv"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/repository/jobDb"
	"github.com/cloud/encoder/vo"
)

const (
	segmentTime = "4000"
	profile     = "live"
	extension   = ".mp4"
	inputDir    = "E:/test/output"
	outputDir   = "package/out.mpd"
	audioFile   = "encoded_audio.m4a"
)

func Package(ctx context.Context, jobId int, jobDao jobDb.JobRepo) error {
	folderDir := filepath.Join(inputDir, strconv.Itoa(jobId))

	template, err := jobDao.GetTemplate(ctx, jobId)
	if err != nil {
		return err
	}

	var vopts = make([]vo.VideoConcat, 0, len(template.Outputs.Video))

	for _, group := range template.Outputs.Video {
		var outputs []vo.VideoConcat
		if err = json.Unmarshal(group.Options, &outputs); err != nil {
			return err
		}

	}

	inputStr := make([]string, 0, len(vopts)+1) // including audio

	for _, opt := range vopts {
		outFile := opt.Height + "@" + opt.Fps + extension
		inputStr = append(inputStr, outFile+"#video")
	}

	// // Might change to support ocmmont fps
	fps := vopts[0].Fps

	inputStr = append(inputStr, audioFile+"#audio")

	packager := cmd.GetMp4box().GenerateDash(
		segmentTime,
		fps,
		inputStr,
		profile,
		outputDir,
	)

	if err := packager.RunInDir(folderDir); err != nil {
		log.Printf("Error in running package(dash) => %s", err)
		panic(err)
		return err
	}

	log.Print("Completed")

	return nil
}
