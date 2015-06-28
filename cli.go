package main

import (
  "github.com/codegangsta/cli"
)

func initApp() *cli.App {
  app := cli.NewApp()

  app.Name = "awsClient"
  app.Version = "0.0.1"
  app.Usage = "AWS Resource Inventory."
  app.Author = "lenfree"
  app.Email = "lenfree.yeung@gmail.com"

  app.Commands = []cli.Command{
    printCommand(),
  }

  return app
}
func printCommand() cli.Command {
  command := cli.Command{
    Name:      "ec2",
    ShortName: "E",
    Usage:     "EC2 Resources",
    Action:    ec2Start,

    Flags: []cli.Flag{
      cli.StringFlag{
        Name:  "region,r",
        Value: "ap-southeast-2",
        Usage: "AWS Region",
      },
    },
  }
  return command
}
