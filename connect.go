package main

import (
  "github.com/aws/aws-sdk-go/service/ec2"
)

func ec2connect(svc *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
  logger.Info("Attempting to connect...")
  resp, err := svc.DescribeInstances(nil)
  if err != nil {
    return nil, err
  }
  return resp, nil
}
