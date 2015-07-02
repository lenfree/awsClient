package main

import(
  "github.com/aws/aws-sdk-go/service/route53"
  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/awsutil"
  "github.com/sirupsen/logrus"
)

func getHostedZones(ctx *cli.Context){
  svc := route53.New(nil)
  params := &route53.ListHostedZonesByNameInput{
  }
  resp, err := svc.ListHostedZonesByName(params)

  if err != nil {
      if awsErr, ok := err.(awserr.Error); ok {
          logger.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
          if reqErr, ok := err.(awserr.RequestFailure); ok {
              logger.Error(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
          }
      } else {
          logger.Error(err.Error())
      }
  }

  for _, hostedZone := range resp.HostedZones {
    logger.Debug(awsutil.StringValue(hostedZone.Name))
    logger.WithFields(logrus.Fields{
      "RR set count": awsutil.StringValue(hostedZone.ResourceRecordSetCount),
      "Zone ID": awsutil.StringValue(hostedZone.ID),
    }).Info("Name", awsutil.StringValue(hostedZone.Name))
  }
}
