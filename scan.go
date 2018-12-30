package scan

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/juju/errors"
)

// Scan is the structure containing the CIRD and the result
type Scan struct {
	CIRD   *net.IPNet `json:"CIRD"`
	Result []*Result  `json:"Result"`
}

// Result contains the structured output of the nmap command
type Result struct {
	Host   string `json:"host"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// New initiliazes a new Scan struct.
// If the cird is not valid, it returns an error.
func New(cird string) (*Scan, error) {
	_, ipNet, err := net.ParseCIDR(cird)

	if err != nil {
		return nil, errors.NewNotValid(err, "")
	}

	return &Scan{
		CIRD:   ipNet,
		Result: []*Result{},
	}, nil
}

// Scan executes the algorithm which scans the given CIRD.
// It updates the result field contained in the Scan struct.
func (s *Scan) Scan() error {
	command := fmt.Sprintf("sudo nmap -sn -oG - %v", s.CIRD)
	output, err := exec.Command("/bin/sh", "-c", command).Output()

	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")
	results, err := parseOutput(lines)

	if err != nil {
		return err
	}

	s.Result = results

	return nil
}

func parseOutput(lines []string) ([]*Result, error) {
	results := []*Result{}

	// The counter starts on on the second line (1): the first line
	// does not have any usefull information.
	// The last line is only a resume of all informations.
	for i := 1; i < len(lines)-2; i++ {
		// Example of a line : `Host: 192.168.1.1 (livebox.home)\t  Status: Up\n`
		// A line contains 2 fields. They are separated by a \t.
		line := lines[i]
		fields := strings.Split(line, "\t")
		hostField := fields[0]
		statusField := fields[1]

		host := strings.Split(hostField, " ")[1]
		name := strings.Split(hostField, " ")[2]
		status := strings.Split(statusField, " ")[1]

		result := &Result{
			Host:   host,
			Name:   name,
			Status: status,
		}

		results = append(results, result)
	}

	if len(results) == 0 {
		return results, errors.New(fmt.Sprintf("output cannot be parsed: %s", lines))
	}

	return results, nil
}

func (r *Result) String() string {
	return fmt.Sprintf("host: %s - name: %s - status: %s\n", r.Host, r.Name, r.Status)
}
