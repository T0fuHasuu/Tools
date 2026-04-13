package scanner

import (
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// PingTarget performs a native TCP connection test
func PingTarget(target string, port int) bool {
	address := fmt.Sprintf("%s:%d", target, port)
	timeout := 2 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}

// RunPingSweep stores and executes the nmap subprocess for host discovery
func RunPingSweep(subnet string) (string, error) {
	cmd := exec.Command("nmap", "-sn", subnet)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// ParseActiveIPs reads raw nmap output and returns a clean list of IPs
func ParseActiveIPs(rawOutput string) []string {
	var ips []string
	lines := strings.Split(rawOutput, "\n")

	// Regex pattern to match an IPv4 address
	ipRegex := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)

	for _, line := range lines {
		// Only look at lines that declare a host is up
		if strings.HasPrefix(line, "Nmap scan report for") {
			// Find all IP addresses in this line
			matches := ipRegex.FindAllString(line, -1)
			if len(matches) > 0 {
				// We take the last match to handle both "report for 192.168.1.1" 
				// and "report for hostname (192.168.1.1)"
				ips = append(ips, matches[len(matches)-1])
			}
		}
	}
	return ips
}