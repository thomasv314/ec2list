package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type InstanceCollection struct {
	Updated   int32             `json:"updated_at"`
	Vpcs      map[string]string `json:"vpcs"`
	Instances []Instance        `json:"instances"`
}

func (c *InstanceCollection) Load() {
	profileCachePath := getProfileCachePath()

	if directoryExists(profileCachePath) {
		c.LoadFromDisk(profileCachePath)
	} else {
		c.LoadFromRemote(true)
	}
}

func (c *InstanceCollection) LoadFromDisk(path string) {
	blob, err := ioutil.ReadFile(path)
	check(err)

	err = json.Unmarshal(blob, &c)
	check(err)
}

func (c *InstanceCollection) LoadFromRemote(saveAfterLoad bool) {
	c.FetchVpcs()
	c.FetchEc2Instances()
	c.Updated = int32(time.Now().Unix())

	if saveAfterLoad {
		c.WriteToDisk()
	}
}

func (c *InstanceCollection) WriteToDisk() {
	blob, err := json.Marshal(&c)
	check(err)

	err = ioutil.WriteFile(getProfileCachePath(), blob, 0644)
	check(err)
}

func (c *InstanceCollection) VpcName(vpcId string) string {
	return c.Vpcs[vpcId]
}
