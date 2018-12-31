package main

import (
	"flag"
	"fmt"

	scan "github.com/fberrez/go-scan"
	log "github.com/sirupsen/logrus"
)

func main() {
	cidr := flag.String("cird", "192.168.1.1/24", "Target of the scan")
	flag.Parse()

	scan, err := scan.New(*cidr)

	if err != nil {
		panic(err)
	}

	log.Info("go-scan uses nmap. You must have root access to use this command.")
	if err = scan.Scan(); err != nil {
		panic(err)
	}

	fmt.Printf("%v", scan.Result)
}
