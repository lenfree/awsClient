package main

import (
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/aws/aws-sdk-go/service/s3"
)

func ec2connect(svc *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
  logger.Info("Attempting to connect...")
  resp, err := svc.DescribeInstances(nil)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func s3connect(svc *s3.S3) (*s3.ListBucketsOutput, error) {
  logger.Info("Attempting to connect...")
  var params *s3.ListBucketsInput
  resp, err := svc.ListBuckets(params)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

