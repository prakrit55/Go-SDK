package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var testSession, _ = session.NewSession(&aws.Config{
	Region: aws.String("us-east-1")},
)
var svc = ec2.New(testSession)

func StartEc2Instance() (output *ec2.StartInstancesOutput, err error) {
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String("i-0d002e2182cfbeb57"),
		},
	}
	result, err := svc.StartInstances(input)

	return result, err
}

func StopEc2Instance() (output *ec2.StopInstancesOutput, err error) {
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String("i-0d002e2182cfbeb57"),
		},
	}
	result, err := svc.StopInstances(input)

	return result, err
}


func main() {
	Startoutput, err := StartEc2Instance()
	if err != nil {
		fmt.Println(err,"The error in stopping an intance: %s")
	}else {
		fmt.Println(Startoutput)
	}
	

	Stopoutput, err := StopEc2Instance()
	if err != nil {
		fmt.Println(err,"The error in stopping an intance: %s")
	}else {
		fmt.Println(Stopoutput)
	}

}
