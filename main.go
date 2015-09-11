package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "queue1",
		},
		cli.StringFlag{
			Name: "queue2",
		},
	}

	app.Action = func(c *cli.Context) {
		client := sqs.New(&aws.Config{
			Region: aws.String("us-east-1"),
		})

		for {
			msgs, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl: aws.String(c.String("queue1")),
			})
			if err != nil {
				log.Fatal(err)
			}

			for _, msg := range msgs.Messages {
				_, err = client.SendMessage(&sqs.SendMessageInput{
					QueueUrl:    aws.String(c.String("queue2")),
					MessageBody: msg.Body,
				})
				if err != nil {
					log.Fatal(err)
				}

				_, err = client.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      aws.String(c.String("queue1")),
					ReceiptHandle: msg.ReceiptHandle,
				})

				if err != nil {
					log.Fatal(err)
				}
				log.Println("Copied msg")
			}

		}
	}

	app.Run(os.Args)

}
