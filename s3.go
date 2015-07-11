package main

import(
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/awsutil"
  "github.com/sirupsen/logrus"
  "strings"
)

func s3List(ctx *cli.Context) {
  svc := s3.New(&aws.Config{Region: "ap-southeast-2"})
  resp, err := s3connect(svc)
  if err != nil {
    if awsErr, ok := err.(awserr.Error); ok {
     logger.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
     if reqErr, ok := err.(awserr.RequestFailure); ok {
         logger.Error(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
      }
    } else {
        logger.Debug(err.Error())
    }
  }
  buckets := s3Buckets(resp)
  for _, bucket := range buckets {
    logger.WithFields(logrus.Fields{
      "Creation Date": bucket.CreationDate,
    }).Info("Name:", bucket.Name)
  }
}

func s3Buckets(resp *s3.ListBucketsOutput) ([]Bucket) {
  var buckets []Bucket
  for _, bucket := range resp.Buckets {
    bucket := Bucket{
      Name: strings.Replace(awsutil.StringValue(bucket.Name), "\"", "", -1),
      CreationDate: bucket.CreationDate,
    }
    buckets = append(buckets, bucket)
  }
  return buckets
}
