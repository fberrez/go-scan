package scan

import (
	"fmt"
	"net"

	"github.com/juju/errors"
)

// Scan is the structure containing the CIDR and the result
type Scan struct {
	CIDR   *net.IPNet `json:"CIDR"`
	Result []*Result  `json:"Result"`
}

// Result contains the structured output of the nmap command
type Result struct {
	Host string `json:"host"`
	Name string `json:"name"`
}

// New initiliazes a new Scan struct.
// If the cidr is not valid, it returns an error.
func New(cidr string) (*Scan, error) {
	_, ipNet, err := net.ParseCIDR(cidr)

	if err != nil {
		return nil, errors.NewNotValid(err, "")
	}

	return &Scan{
		CIDR:   ipNet,
		Result: []*Result{},
	}, nil
}

// Scan executes the algorithm which scans the given CIDR.
// It updates the result field contained in the Scan struct.
func (s *Scan) Scan() error {
	address, err := s.getHosts()
	if err != nil {
		return err
	}

	results := []*Result{}
	for _, a := range address {
		host, err := net.LookupAddr(a)
		// If an error occurs, it assumes that the address is not used
		// in the given CIDR.
		if err != nil {
			continue
		}

		result := &Result{
			Host: a,
			Name: host[0],
		}

		results = append(results, result)
	}

	s.Result = results
	return nil
}

// getHosts returns a slice containing all addresses from a given CIDR.
func (s *Scan) getHosts() ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(s.CIDR.String())
	if err != nil {
		return nil, err
	}

	var ips []string
	for ipAddr := ip.Mask(ipnet.Mask); ipnet.Contains(ipAddr); inc(ipAddr) {
		ips = append(ips, ipAddr.String())
	}

	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

// inc increments the given IP address.
// http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func (r *Result) String() string {
	return fmt.Sprintf("host: %s - name: %s\n", r.Host, r.Name)
}
