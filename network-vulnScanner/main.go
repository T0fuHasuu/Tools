package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"network-vulnScanner/scanner"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Cyan  = "\033[36m"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/c", "cls").Run()
	} else {
		exec.Command("clear").Run()
	}
}

func menu(target string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		fmt.Println(Cyan + "[+] Menu" + Reset)
		fmt.Println("1. General Vuln Scan")
		fmt.Println("2. CVE Vuln Scan")
		fmt.Println("q. Quit")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			scanner.GeneralVuln(target)
		case "2":
			scanner.CVEScan(target)
		case "q", "Q":
			return
		default:
			fmt.Println("Invalid option")
		}

		fmt.Print("Press Enter to continue...")
		reader.ReadString('\n')
	}
}

func main() {
	target := scanner.GetInput()

	fmt.Println(Red + "[-] Starting..." + Reset)

	if err := scanner.PrerequisitesOK(target); err != nil {
		fmt.Println(Red+"[-] "+err.Error()+Reset)
		os.Exit(1)
	}

	fmt.Println(Green + "[+] All checks passed" + Reset)
	menu(target)
}