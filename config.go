package main

import (
	"fmt"
	"os"
)

const (
	ActiveProfileKey     string = "ACTIVE_EP"
	DefaultActiveProfile string = "default"
)

// mkdir -p ~/.eps/
func SetupWorkspace() {
	cachePath := getCachePath()
	err := os.MkdirAll(cachePath, 0600)

	if err != nil {
		fmt.Println("Error setting up workspace:", err)
	} else {
		fmt.Println("Setup workspace:", cachePath)
	}
}

func getActiveProfile() string {
	profile := os.Getenv(ActiveProfileKey)
	if profile == "" {
		return DefaultActiveProfile
	} else {
		return profile
	}
}

func getCachePath() string {
	ec2CachePath := os.Getenv("EC2LIST_CACHE_PATH")

	if ec2CachePath == "" {
		home := os.Getenv("HOME")

		if home == "" {
			home = "/root"
		}

		return home + "/.ec2list"
	} else {
		return ec2CachePath
	}
}

func getProfileCachePath() string {
	return getCachePath() + "/" + getActiveProfile() + "-instances.json"
}

func directoryExists(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false
	} else {
		return true
	}

}
