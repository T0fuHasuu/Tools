package scanner

import (
	"errors"
	"os/exec"
	"runtime"
)

func PingGoogle() error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "google.com")
	} else {
		cmd = exec.Command("ping", "-c", "1", "google.com")
	}

	return cmd.Run()
}

func HasNmap() bool {
	_, err := exec.LookPath("nmap")
	return err == nil
}

func PrerequisitesOK() error {
	if err := PingGoogle(); err != nil {
		return errors.New("No Internet Connection !!!")
	}

	if !HasNmap() {
		return errors.New("Nmap Is Not Installed !!!")
	}

	return nil
}