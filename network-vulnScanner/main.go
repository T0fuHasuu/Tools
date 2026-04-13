package main

import (
	"fmt"
	"log"
	"network-vulnScanner/scanner"
)

// ANSI Color Codes for a modern CLI look
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func avaIP() {
	var networkIP string
	
	// Stylish Input Prompt
	fmt.Print(Cyan + "╭─ Enter IP/Subnet (Ex : 192.168.1.0/24) ❯ " + Reset)
	fmt.Scan(&networkIP)

	fmt.Printf(Yellow + "│  Scanning %s . . .\n" + Reset, networkIP)

	// Execute the Nmap command
	rawResults, err := scanner.RunPingSweep(networkIP)
	if err != nil {
		log.Fatalf(Red+"\n[-] Nmap command failed: %v\nOutput: %s"+Reset, err, rawResults)
	}

	// Parse the raw results to get JUST the IPs
	activeIPs := scanner.ParseActiveIPs(rawResults)
	totalDevices := len(activeIPs)

	// Modern Output Box
	fmt.Println(Cyan + "╰──────────────────────────────────────────────────" + Reset)
	fmt.Println()
	
	if totalDevices == 0 {
		fmt.Println(Red + "[-] Scan Complete! No active devices found." + Reset)
	} else {
		fmt.Printf(Green+"[+] Scan Complete! Found %d active devices.\n"+Reset, totalDevices)
		fmt.Println(Blue + "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" + Reset)
		
		// Loop through the clean IPs and print them stylishly
		for i, ip := range activeIPs {
			fmt.Printf(Cyan+"  [%d] "+Green+"%s\n"+Reset, i+1, ip)
		}
		
		fmt.Println(Blue + "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" + Reset)
	}
	fmt.Println()
}

func main() {
	avaIP()

	// Testing single target (from your original code)
	target := "192.168.1.5"
	port := 80

	fmt.Printf(Yellow+"Pinging %s on port %d...\n"+Reset, target, port)

	isAlive := scanner.PingTarget(target, port)

	if isAlive {
		fmt.Println(Green + "[+] Success! The target is alive and the port is OPEN." + Reset)
	} else {
		fmt.Println(Red + "[-] Failed. The target is down or the port is CLOSED." + Reset)
	}
}