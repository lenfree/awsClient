package main

import(
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/sirupsen/logrus"
  "github.com/codegangsta/cli"
  "reflect"
)

func ec2List(ctx *cli.Context) {
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
  ec2Instance(resp, instanceState)
  totalEc2:= ec2Count(resp, instanceState)
  logger.Debug(instanceState, " instances ", totalEc2)
}

func ec2Total(ctx *cli.Context) {
  region := ctx.String("region")
  instanceState := ctx.String("state")
  logger.Debug("AWS region: ", region)
  // Create an EC2 service object in the "ap-southeast-2" region
  svc := ec2.New(&aws.Config{Region: region})
  resp, err := ec2connect(svc)
  if err != nil {
      panic(err)
  }
  totalEc2:= ec2Count(resp, instanceState)
  logger.Debug(instanceState, " instances ", totalEc2)
}

func ec2Count(resp *ec2.DescribeInstancesOutput, instanceState string) (count int) {
  count = 0
  for idx, _ := range resp.Reservations {
    for _, inst := range resp.Reservations[idx].Instances {
      switch instanceState {
      case "stopped":
        if *inst.State.Name == "stopped" {
          count += 1
        }
      case "running":
        if *inst.State.Name == "running" {
          count += 1
        }
      default:
        count = len(resp.Reservations)
      }
    }
  }
  return
}

func ec2Instance (resp *ec2.DescribeInstancesOutput, instanceState string) {
  // resp has all of the response data, pull out instance IDs:
  for idx, _ := range resp.Reservations {
    for _, inst := range resp.Reservations[idx].Instances {
      switch instanceState {
      case "stopped":
        if *inst.State.Name == "stopped" {
          instanceTags := getTags(inst)
          for idx, tag := range instanceTags {
            logger.WithFields(logrus.Fields{
              "Launch date": inst.LaunchTime,
              "tag name": tag,
              "tag #": idx,
              "Instance ID": inst.InstanceID,
              "State": *inst.State.Name,
            }).Info("Instance ID: ", *inst.InstanceID)
          }
        }
      case "running":
        if *inst.State.Name == "running" {
          instanceTags := getTags(inst)
          for idx, tag := range instanceTags {
            logger.WithFields(logrus.Fields{
              "Launch date": inst.LaunchTime,
              "tag name": tag,
              "tag #": idx,
              "Instance ID": inst.InstanceID,
              "State": *inst.State.Name,
            }).Info("Instance ID: ", *inst.InstanceID)
          }
        }
      default:
        instanceTags := getTags(inst)
        for idx, tag := range instanceTags {
          logger.WithFields(logrus.Fields{
            "Launch date": inst.LaunchTime,
            "tag name": tag,
            "tag #": idx,
            "Instance ID": inst.InstanceID,
            "State": *inst.State.Name,
          }).Info("Instance ID: ", *inst.InstanceID)
        }
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
