package main

import(
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/sirupsen/logrus"
)
func ec2Count (resp *ec2.DescribeInstancesOutput) int {
  return len(resp.Reservations)
}

func ec2Instance (resp *ec2.DescribeInstancesOutput) {
  // resp has all of the response data, pull out instance IDs:

  for idx, _ := range resp.Reservations {
    for _, inst := range resp.Reservations[idx].Instances {
      logger.WithFields(logrus.Fields{
        "Instance ID": *inst.InstanceID,
        "State": *inst.State.Name,
      }).Info("Instance ID: ", *inst.InstanceID)
    }
  }
}
