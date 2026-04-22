package scanner

import (
	"os"
	"os/exec"
)

func GeneralVuln(target string) error {
	cmd := exec.Command("nmap", "-Pn", "-sV", "--script", "vuln", target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CVEScan(target string) error {
	cmd := exec.Command("nmap", "-sV", "--script", "vulners", target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
