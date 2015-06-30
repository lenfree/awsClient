awsClient
=========

It's easy to spin up AWS resources, however, when it comes to clean up job,
it's always a challenge. There are instances where team members forgot to
decommission or stop resources after commissioning them and it just racks up
bill. This program/executable will help user/s automatically delete old EC2
instances based on tags. Let's say when a EC2 instance doesn't have production
tag, this means it could be a dev instance and supposed to be destroyed after
provisioning it.

Getting started
----------------
```shell
export AWS_ACCESS_KEY_ID=keyid
export AWS_SECRET_ACCESS_KEY=secretkey
go run *.go --help
go run *.go E list --region awsRegion
```

Build and execute
-----------------
```shell
go build
./aws E count --region "ap-southeast-2" --state "running"
./aws ec2 list --region "ap-southeast-2" --state "stopped"
./aws --help
```

