package main

import (
	"flag"
	"fmt"
	"github.com/gosuri/uitable"
	"sort"
	"strings"
)

func (c *InstanceCollection) ShowList() {

	fmt.Println(len(c.Instances))
	c.FilterInstances()

	instances := Instances(c.Instances)
	sort.Sort(instances)

	printUiTable(instances)

}

func printUiTable(instances Instances) {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("vpcId", "vpc", "name", "type", "privateIp", "publicIp", "role", "environment")

	for i := range instances {
		instance := instances[i]
		table.AddRow(
			instance.VpcId,
			instance.Vpc,
			instance.GetTag("Name"),
			instance.Type,
			instance.PrivateIp,
			instance.PublicIp,
			instance.GetTag("Role"),
			instance.GetTag("Environment"),
		)
	}

	fmt.Print(table)
}

type Instances []Instance

func (slice Instances) Len() int {
	return len(slice)
}

func (slice Instances) Less(i, j int) bool {
	return slice[i].GetTag("Name") < slice[j].GetTag("Name")
}

func (slice Instances) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (c *InstanceCollection) FilterInstances() {
	if *flgFilterInstances {
		filters := flag.Args()

		var (
			keyval   []string
			key, val string
		)

		for i := range filters {
			keyval = strings.Split(filters[i], ":")
			key = keyval[0]
			val = keyval[1]

			switch key {
			case "vpc":
				c.FilterByVpc(val)
			case "role":
				c.FilterByRole(val)
			}
		}
	}
}

func (c *InstanceCollection) FilterByVpc(vpcName string) {
	filtered := []Instance{}

	for i := 0; i < len(c.Instances); i++ {
		instance := c.Instances[i]
		if instance.Vpc != "" {
			if strings.Contains(instance.Vpc, vpcName) {
				filtered = append(filtered, instance)
			}
		}
	}

	c.Instances = filtered
}

func (c *InstanceCollection) FilterByRole(role string) {
	filtered := []Instance{}

	for i := 0; i < len(c.Instances); i++ {
		instance := c.Instances[i]
		if instance.GetTag("Role") != "" {
			if strings.Contains(instance.GetTag("Role"), role) {
				filtered = append(filtered, instance)
			}
		}
	}

	c.Instances = filtered
}
