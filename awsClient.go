package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/sirupsen/logrus"
  "reflect"
)

var logger = logrus.New()

func init() {
  logger.Formatter = new(logrus.JSONFormatter)

  // Only log the warning severity or above.
  //log.SetLevel(log.WarnLevel)
  logger.Level = logrus.DebugLevel
}

func main() {
  // Create an EC2 service object in the "ap-southeast-2" region
  svc := ec2.New(&aws.Config{Region: "ap-southeast-2"})
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
