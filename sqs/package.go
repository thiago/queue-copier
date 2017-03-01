package sqs

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsqs "github.com/aws/aws-sdk-go/service/sqs"
)

func Copy(queue1, queue2 string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	client := awsqs.New(sess)

	for {
		msgs, err := client.ReceiveMessage(&awsqs.ReceiveMessageInput{
			QueueUrl: aws.String(queue1),
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, msg := range msgs.Messages {
			_, err = client.SendMessage(&awsqs.SendMessageInput{
				QueueUrl:    aws.String(queue2),
				MessageBody: msg.Body,
			})
			if err != nil {
				log.Fatal(err)
			}

			_, err = client.DeleteMessage(&awsqs.DeleteMessageInput{
				QueueUrl:      aws.String(queue1),
				ReceiptHandle: msg.ReceiptHandle,
			})

			if err != nil {
				log.Fatal(err)
			}
			log.Println("Copied msg")
		}
	}
}
