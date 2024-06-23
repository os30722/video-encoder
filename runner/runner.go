package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloud/encoder/codecs"
	"github.com/cloud/encoder/mom"
	"github.com/cloud/encoder/repository/jobDb"
	"github.com/cloud/encoder/vo"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Start(ctx context.Context, jobDao *jobDb.JobDao) error {
	msgChan, err := mom.GetTaskMsg()
	if err != nil {
		return err
	}

	fmt.Println("Listening for messages")
	for msg := range msgChan {
		var taskMsg vo.TaskMsg
		err := json.Unmarshal(msg.Body, &taskMsg)
		fmt.Println(taskMsg)
		go func(msg amqp.Delivery) {
			if err != nil {
				log.Printf("Error in message %s => %s", string(msg.Body), err)
			}
			switch codec := taskMsg.Options.Codec; codec {
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

			completed, err := jobDao.UpdateAndReturnCompletion(ctx, taskMsg.JobId)
			if err != nil {
				log.Print(err)
			}

			if completed {
				log.Print("COmpleted")
				info, err := jobDao.GetJobFileAndOptions(ctx, taskMsg.JobId)
				if err != nil {
					log.Print(err)
				}

				err = codecs.Concat(info)
				if err != nil {
					log.Println(err)
				}
			}

			msg.Ack(false)
		}(msg)
	}

	return nil
}
