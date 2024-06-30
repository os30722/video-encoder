package packager

import (
	"context"
	"log"
	"path/filepath"
	"strconv"

	"github.com/cloud/encoder/cmd"
	"github.com/cloud/encoder/repository/jobDb"
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

	vopts := template.Outputs.Video

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
