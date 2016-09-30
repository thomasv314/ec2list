package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flgRunSetup         = flag.Bool("setup", false, "setup workspace")
	flgUpdateCollection = flag.Bool("update", false, "update the current collection")
	flgFilterInstances  = flag.Bool("f", false, "use filters to select instances")
	flgShortOutput      = flag.Bool("short", false, "short output - only show private ips")
)

func main() {
	fmt.Println("ec2list", os.Args)

	flag.Parse()

	if *flgRunSetup {
		SetupWorkspace()
		exit("Workspace successfully setup.")
	}

	if *flgUpdateCollection {
		collection := InstanceCollection{}
		collection.LoadFromRemote(true)
		exit("Updated instance collection for " + getActiveProfile())
	}

	collection := InstanceCollection{}
	collection.Load()
	collection.ShowList()
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
