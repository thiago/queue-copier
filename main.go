package main

import (
	"os"

	"github.com/DispatchMe/queue-copier/rabbit"
	"github.com/DispatchMe/queue-copier/sqs"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "queue copier"
	app.Usage = "Copy Messages from one queue to another"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "mode",
		},
		cli.StringFlag{
			Name: "dead-queue",
		},
		cli.StringFlag{
			Name: "exchange",
		},
		cli.StringFlag{
			Name: "queue1",
		},
		cli.StringFlag{
			Name: "queue2",
		},
	}

	app.Action = func(c *cli.Context) {
		switch c.String("mode") {
		case "rabbit":
			rabbit.Republish(c.String("dead-queue"), c.String("exchange"))
		case "sqs":
			sqs.Copy(c.String("queue1"), c.String("queue2"))
		}
	}

	app.Run(os.Args)
}
