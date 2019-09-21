# Go-scan

![go-scan illustration](.github/go-scan.png "Go-scan")

Go-scan is a host discovery written in go. It returns a structured list of devices connected to a specified network.

## Getting started
```sh
# Download dependencies
$ dep ensure

# Generate the binary
$ make

# Read the doc
$ ./go-scan --help

# Scan a specific CIDR
$ ./go-scan --cidr 72.5.1.1/24

# Scan local network (192.168.1.1/24)
$ ./go-scan
```

## How to use it
```go
import (
  fmt

  scan "github.com/fberrez/go-scan"
)

func main() {
  scan, err := scan.New("192.168.1.1/24")

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
```

## Example of results
```sh
[
 {
   "host": "192.168.1.1",
   "name": "box"
 },
 {
   "host": "192.168.1.10",
   "name": "device1"
 },
 # ... other devices
]
```
