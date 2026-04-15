package main

import (
	"fmt"
	"os"

	"network-vulnScanner/scanner"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func menu() {
	fmt.Println(Cyan + "[+] Loading menu..." + Reset)
	// your menu code here
}

func main() {
	fmt.Println(Red + "[-] Starting..." + Reset)

	if err := scanner.PrerequisitesOK(); err != nil {
		fmt.Println(Red+"[-] "+err.Error()+Reset)
		os.Exit(1)
	}

	fmt.Println(Green + "[+] Connection and Nmap check passed" + Reset)
	menu()
}