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
  instanceState := ctx.String("state")
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
  totalEc2:= ec2Count(resp, instanceState)
  logger.Debug("Total no of EC2 instances: ", totalEc2)

  //Get instance tag
}

func ec2Count (resp *ec2.DescribeInstancesOutput, instanceState string) int {
  countTotal := len(resp.Reservations)
  count := 0
  countStopped := 0
  countRunning := 0
  for idx, _ := range resp.Reservations {
    for _, inst := range resp.Reservations[idx].Instances {
      if *inst.State.Name == "stopped" {
        countStopped += 1
      } else if *inst.State.Name == "running" {
        countRunning += 1
      }
    }
  }
  logger.Debug("Running instsances", countRunning)
  logger.Debug("Stopped", countRunning)
  logger.WithFields(logrus.Fields{
    "Total": countTotal,
    "Running": countRunning,
    "Stopped": countStopped,
  }).Info("EC2 Instances")
  switch instanceState {
  case "stopped":
    count = countStopped
  case "running":
    count = countRunning
  default:
    count = countTotal
  }
  return count
}

func ec2Instance (resp *ec2.DescribeInstancesOutput) {
  // resp has all of the response data, pull out instance IDs:
  for idx, _ := range resp.Reservations {
    for _, inst := range resp.Reservations[idx].Instances {
      instanceTags := getTags(inst)
      for idx, tag := range instanceTags {
        logger.WithFields(logrus.Fields{
          "tag name": tag,
          "tag #": idx,
          "Instance ID": inst.InstanceID,
          "State": *inst.State.Name,
        }).Info("Instance ID: ", *inst.InstanceID)
      }
    }
  }
}

func getTags(inst *ec2.Instance) []string{
  var tags []string
  for _, tag := range inst.Tags {
    tags = append(tags, *tag.Value)
  }
  return tags
}
