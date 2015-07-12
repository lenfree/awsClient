package main

import(
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/sirupsen/logrus"
  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/aws/awsutil"
  "strings"
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
      instanceTags := getTags(inst)
      switch instanceState {
      case "stopped":
        if *inst.State.Name == "stopped" {
          logger.WithFields(logrus.Fields{
            "Launch date": inst.LaunchTime,
            "tags": instanceTags,
            "Instance ID": inst.InstanceID,
            "State": *inst.State.Name,
          }).Info("Instance ID: ", *inst.KeyName)
        }
      case "running":
        if *inst.State.Name == "running" {
          logger.WithFields(logrus.Fields{
            "Launch date": inst.LaunchTime,
            "tags": instanceTags,
            "Instance ID": inst.InstanceID,
            "State": *inst.State.Name,
          }).Info("Instance: ", *inst.KeyName)
        }
      default:
        logger.WithFields(logrus.Fields{
          "Launch date": inst.LaunchTime,
          "tags": instanceTags,
          "Instance ID": inst.InstanceID,
          "State": *inst.State.Name,
        }).Info("Instance: ", *inst.KeyName)
      }
    }
  }
}

func getTags(inst *ec2.Instance) ([]EC2InstanceTags) {
  var tags []EC2InstanceTags
  for _, tag := range inst.Tags {
    tag := EC2InstanceTags{
      Name: strings.Replace(awsutil.StringValue(tag.Key), "\"", "", -1),
      Value: strings.Replace(awsutil.StringValue(tag.Value), "\"", "", -1),
    }
    tags = append(tags, tag)
  }
  return tags
}
