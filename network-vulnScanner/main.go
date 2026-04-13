package main

import (
	"fmt"
	"network-vulnScanner/scanner"
)

func main() {
	target := "scanme.nmap.org"
	port := 80

	fmt.Printf("Pinging %s on port %d...\n", target, port)

	// Call the PingTarget function from our scanner package
	isAlive := scanner.PingTarget(target, port)

	// Check the results
	if isAlive {
		fmt.Println("[+] Success! The target is alive and the port is OPEN.")
	} else {
		fmt.Println("[-] Failed. The target is down or the port is CLOSED.")
	}
}