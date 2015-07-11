package main

import(
  "github.com/aws/aws-sdk-go/service/route53"
  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/aws/awserr"
  "github.com/aws/aws-sdk-go/aws/awsutil"
  "github.com/sirupsen/logrus"
  "strings"
)

func getHostedZones(ctx *cli.Context){
  svc := route53.New(nil)
  params := &route53.ListHostedZonesByNameInput{}
  rrset := ctx.Bool("recordset")
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
    if rrset == true {
      rrset, err := getRR(hostedZone.ID)
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
      for _, rrecord := range rrset {
        logger.WithFields(logrus.Fields{
          "type":  rrecord.Type,
          "value": rrecord.Value,
          "TTL": rrecord.Ttl,
        }).Info("Name: ", rrecord.Name)
      }
    } else {
      logger.WithFields(logrus.Fields{
        "RR set count": awsutil.StringValue(hostedZone.ResourceRecordSetCount),
        "Zone ID": strings.Replace(awsutil.StringValue(hostedZone.ID), "\"", "", -1),
      }).Info("Name: ", strings.Replace(awsutil.StringValue(hostedZone.Name), "\"", "", -1))
    }
  }
}

func getRR(hostedZoneID *string) ([]R53, error) {
  svc := route53.New(nil)
  rrsresp, err := svc.ListResourceRecordSets (& route53.ListResourceRecordSetsInput {HostedZoneID: hostedZoneID})
  if err != nil {
    return nil, err
  }
  var rrSetResults []R53
  for i := range rrsresp.ResourceRecordSets {
    for j := range rrsresp.ResourceRecordSets[i].ResourceRecords{
      rrset := R53{
        Name:  strings.Replace(awsutil.StringValue(*rrsresp.ResourceRecordSets[i].Name), "\"", "", -1),
        Type:  strings.Replace(awsutil.StringValue(*rrsresp.ResourceRecordSets[i].Type), "\"", "", -1),
        Ttl:  *rrsresp.ResourceRecordSets[i].TTL,
        Value: strings.Replace(awsutil.StringValue(rrsresp.ResourceRecordSets[i].ResourceRecords[j].Value), "\"", "", -1),
      }
    rrSetResults = append(rrSetResults, rrset)
    }
  }
  return rrSetResults, nil
}
