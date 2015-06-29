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
    ec2Resource(),
  //  s3Resource(),
  }
  return app
}

func ec2Resource() cli.Command {
  command := cli.Command{
    Name:      "ec2",
    ShortName: "E",
    Usage:     "EC2 Resources",
    Subcommands: []cli.Command{
      {
        Name:  "list",
        Usage: "list EC2 instances",
        Action: ec2List,
        Flags: []cli.Flag{
          cli.StringFlag{
            Name:  "region,r",
            Value: "ap-southeast-2",
            Usage: "AWS Region, default to ap-southeast-2. For more info \ncheck http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html",
          },
          cli.StringFlag{
            Name:  "state, s",
            Value: "all",
            Usage: "state < running | stopped | all >. Default to all.",
          },
        },
      },
      {
        Name:  "count",
        Usage: "count EC2 instances",
        Action: ec2Total,
        Flags: []cli.Flag{
          cli.StringFlag{
            Name:  "region,r",
            Value: "ap-southeast-2",
            Usage: "AWS Region, default to ap-southeast-2. For more info \ncheck http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html",
          },
          cli.StringFlag{
            Name:  "state, s",
            Value: "all",
            Usage: "state < running | stopped | all >. Default to all.",
          },
        },
      },
    },
  }
  return command
}

/*func s3Resource() cli.Command {
  command := cli.Command{
    Name:      "s3",
    ShortName: "S",
    Usage:     "S3 Resources",
    Action:    s3Start,

    Flags: []cli.Flag{
      cli.StringFlag{
        Name:  "region,r",
        Value: "ap-southeast-2",
        Usage: "AWS Region",
      },
    },
  }
  return command
}*/
