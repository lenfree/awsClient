package main

import(
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/awsutil"
  "github.com/sirupsen/logrus"
  "reflect"
)

func s3List(ctx *cli.Context) {
  svc := s3.New(&aws.Config{Region: "ap-southeast-2"})
  logger.Debug(reflect.TypeOf(svc))
  resp, err := s3connect(svc)
  if err != nil {
    if awsErr, ok := err.(awserr.Error); ok {
      // Generic AWS Error with Code, Message, and original error (if any)
     logger.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
     if reqErr, ok := err.(awserr.RequestFailure); ok {
         // A service error occurred
         logger.Error(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
      }
    } else {
        // This case should never be hit, the SDK should always return an
        // error which satisfies the awserr.Error interface.
        logger.Debug(err.Error())
    }
  }
  s3Buckets(resp)
}

func s3Buckets(resp *s3.ListBucketsOutput) {
  for _, bucket := range resp.Buckets {
    logger.WithFields(logrus.Fields{
      "Creation Date": bucket.CreationDate,
    }).Info("Name", awsutil.StringValue(bucket.Name))
  }
}
