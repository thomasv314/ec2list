package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	stderr              *log.Logger
	flgRunSetup         = flag.Bool("setup", false, "setup workspace")
	flgUpdateCollection = flag.Bool("update", false, "update the current collection")
	flgFilterInstances  = flag.Bool("f", false, "use filters to select instances")
	flgShortOutput      = flag.Bool("s", false, "short output - only show private ips")
)

func main() {
	flag.Parse()
	stderr = log.New(os.Stderr, "", 0)

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

func check(err error) {
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}
