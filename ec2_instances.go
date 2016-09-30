package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Instance struct {
	VpcId     string            `json:"vpc_id"`
	Vpc       string            `json:"vpc"`
	Id        string            `json:"id"`
	Type      string            `json:"type"`
	PrivateIp string            `json:"private_ip"`
	PublicIp  string            `json:"public_ip"`
	State     string            `json:"state"`
	Tags      map[string]string `json:"tags"`
}

func (c *InstanceCollection) FetchEc2Instances() {

	sess, err := session.NewSession()
	check(err)

	svc := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{}

	resp, err := svc.DescribeInstances(params)
	check(err)

	instances := []Instance{}

	for r := range resp.Reservations {
		for i := range resp.Reservations[r].Instances {
			instanceResponse := resp.Reservations[r].Instances[i]
			instance := Instance{}

			if instanceResponse.VpcId != nil {
				instance.VpcId = *instanceResponse.VpcId

				vpcName := c.Vpcs[instance.VpcId]

				if vpcName != "" {
					instance.Vpc = vpcName
				}
			}

			instance.Id = *instanceResponse.InstanceId
			instance.Type = *instanceResponse.InstanceType

			instance.State = *instanceResponse.State.Name

			if instanceResponse.PrivateIpAddress != nil {
				instance.PrivateIp = *instanceResponse.PrivateIpAddress
			}

			if len(instanceResponse.Tags) > 0 {
				instance.Tags = make(map[string]string)
				for t := range instanceResponse.Tags {
					tag := instanceResponse.Tags[t]
					instance.Tags[*tag.Key] = *tag.Value
				}
			}

			instances = append(instances, instance)
		}
	}

	c.Instances = instances
}

func (i *Instance) GetTag(tagKey string) string {
	for key, value := range i.Tags {
		if key == tagKey {
			return value
		}
	}

	return ""
}
