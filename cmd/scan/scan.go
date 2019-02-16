package main

import (
	"encoding/json"
	"flag"
	"fmt"

	scan "github.com/fberrez/go-scan"
)

func main() {
	cidr := flag.String("cidr", "192.168.1.1/24", "Target of the scan")
	flag.Parse()

	scan, err := scan.New(*cidr)
	if err != nil {
		panic(err)
	}

	if err = scan.Scan(); err != nil {
		panic(err)
	}

	results, err := json.Marshal(scan.Result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", results)
}
