package main

import(
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/sirupsen/logrus"
"github.com/codegangsta/cli"
  "reflect"
)

func ec2Start(ctx *cli.Context) {
  region := ctx.String("region")

  logger.Debug("AWS region: ", region)
  // Create an EC2 service object in the "ap-southeast-2" region
  svc := ec2.New(&aws.Config{Region: region})
  logger.Debug(reflect.TypeOf(svc))
  resp, err := ec2connect(svc)
  if err != nil {
      panic(err)
  }
  logger.Debug(reflect.TypeOf(resp))
  ec2Instance(resp)
  totalEc2:= ec2Count(resp)
  logger.Debug("Total no of EC2 instance: ", totalEc2)

  //Get instance tag
}

func ec2Count (resp *ec2.DescribeInstancesOutput) int {
  return len(resp.Reservations)
}

func ec2Instance (resp *ec2.DescribeInstancesOutput) {
  // resp has all of the response data, pull out instance IDs:

  for idx, _ := range resp.Reservations {
    for _, inst := range resp.Reservations[idx].Instances {
      logger.WithFields(logrus.Fields{
        "Instance ID": *inst.InstanceID,
        "State": *inst.State.Name,
      }).Info("Instance ID: ", *inst.InstanceID)
    }
  }
}
