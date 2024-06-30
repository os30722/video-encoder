package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloud/encoder/codecs"
	"github.com/cloud/encoder/mom"
	"github.com/cloud/encoder/packager"
	"github.com/cloud/encoder/repository/jobDb"
	"github.com/cloud/encoder/vo"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Start(ctx context.Context, jobDao jobDb.JobRepo) error {
	msgChan, err := mom.GetTaskMsg()
	if err != nil {
		return err
	}

	fmt.Println("Listening for messages")
	for msg := range msgChan {
		go func(msg amqp.Delivery) {
			var taskMsg vo.TaskMsg

			err := json.Unmarshal(msg.Body, &taskMsg)
			if err != nil {
				log.Printf("Error in message %s => %s", string(msg.Body), err)
			}

			switch codec := taskMsg.Codec; codec {
			case "split":
				err = SubmitJob(ctx, taskMsg, jobDao)
			case "h264":
				err = codecs.RunH264(taskMsg)
			case "aac":
				err = codecs.RunAac(taskMsg)
			default:
				log.Println("Error Codec Not Found ")
			}

			if err != nil {
				log.Println(err)

			}

			if taskMsg.Output.Height != "" {
				partName := taskMsg.Output.Height + "@" + taskMsg.Output.Fps

				completed, jobCompleted, err := jobDao.UpdateAndReturnCompletion(ctx, taskMsg.JobId, partName)
				if err != nil {
					log.Print(err)
					return
				}

				if completed {
					log.Print("COmpleted")
					if err = codecs.Concat(taskMsg); err != nil {
						log.Print(err)
					}
				}

				if jobCompleted {
					if err = packager.Package(ctx, taskMsg.JobId, jobDao); err != nil {
						log.Print(err)
					}
				}
			}

			msg.Ack(false)
		}(msg)
	}

	return nil
}
