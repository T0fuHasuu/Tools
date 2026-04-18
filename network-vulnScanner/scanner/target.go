package scanner

import (
	"os/exec"
)

func GeneralVuln(target string) error {
	cmd := exec.Command("nmap", "-Pn", "-sV", "--script", "vuln", target)
	return cmd.Run()
}