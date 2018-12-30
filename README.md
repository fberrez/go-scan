# Go-scan

Go-scan is a tool based on nmap. It executes a host discovery and returns a structured list of devices connected to a specified network.

## Getting started
```sh
# Download dependencies
$ dep ensure

# Generate the binary
$ make

# Read the doc
$ ./go-scan --help
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
