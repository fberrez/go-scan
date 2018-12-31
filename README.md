# Go-scan

Go-scan is a tool based on nmap, written in go. It executes a host discovery and returns a structured list of devices connected to a specified network.

## Getting started
```sh
# Download dependencies
$ dep ensure

# Generate the binary
$ make

# Read the doc
$ ./go-scan --help
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

  fmt.Printf("%v", scan.Result)
}
```

## Example of results
```sh
[
 {
   "host": "192.168.1.1",
   "name": "(box)"
   "status": "Up"
 },
 {
   "host": "192.168.1.10",
   "name": "(device1)"
   "status": "Up"
 },
 # ... other devices
]
```
