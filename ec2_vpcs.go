package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (c *InstanceCollection) FetchVpcs() {
	sess, err := session.NewSession()
	check(err)

	svc := ec2.New(sess)

	params := &ec2.DescribeVpcsInput{}
	resp, err := svc.DescribeVpcs(params)
	check(err)

	vpcs := make(map[string]string)

	for i := range resp.Vpcs {
		vpc := resp.Vpcs[i]
		vpcs[*vpc.VpcId] = getVpcTag(vpc, "Name")
	}

	c.Vpcs = vpcs
}

func getVpcTag(vpc *ec2.Vpc, tagKey string) string {
	for i := range vpc.Tags {
		tag := vpc.Tags[i]
		if *tag.Key == tagKey {
			return *tag.Value
		}
	}
	return ""
}
