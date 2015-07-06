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
      rrset := getRR(hostedZone.ID)
      for _, rrecord := range rrset {
        logger.WithFields(logrus.Fields{
          "type":  rrecord["type"],
          "value": rrecord["value"],
          "TTL": rrecord["ttl"],
        }).Info("Name", rrecord["name"])
      }
    } else {
      logger.WithFields(logrus.Fields{
        "RR set count": awsutil.StringValue(hostedZone.ResourceRecordSetCount),
        "Zone ID": awsutil.StringValue(hostedZone.ID),
      }).Info("Name", awsutil.StringValue(hostedZone.Name))
    }
  }
}

func getRR(hostedZoneID *string) ([]map[string]string) {
  svc := route53.New(nil)
  rrsresp, err := svc.ListResourceRecordSets (& route53.ListResourceRecordSetsInput {HostedZoneID: hostedZoneID})
  if err != nil {
     logger.Fatal(err)
  }
  rrSetResults := []map[string]string{}
  for i := range rrsresp.ResourceRecordSets {
    for j := range rrsresp.ResourceRecordSets[i].ResourceRecords{
      rrset := map[string]string{
        "name":  awsutil.StringValue(*rrsresp.ResourceRecordSets[i].Name),
        "type":  awsutil.StringValue(*rrsresp.ResourceRecordSets[i].Type),
        "ttl":  awsutil.StringValue(*rrsresp.ResourceRecordSets[i].TTL),
        "value": awsutil.StringValue(rrsresp.ResourceRecordSets[i].ResourceRecords[j].Value),
      }
    rrSetResults = append(rrSetResults, rrset)
    }
  }
  return rrSetResults
}
