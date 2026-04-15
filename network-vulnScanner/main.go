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
	Cyan   = "\033[36m"
)

func menu() {
	fmt.Println(Cyan + "[+] Loading menu..." + Reset)
}

func main() {
	fmt.Println(Red + "[-] Starting..." + Reset)

	target := "192.168.1.1" // replace with user input later

	if err := scanner.PrerequisitesOK(target); err != nil {
		fmt.Println(Red+"[-] "+err.Error()+Reset)
		os.Exit(1)
	}

	fmt.Println(Green + "[+] All checks passed" + Reset)
	menu()
}