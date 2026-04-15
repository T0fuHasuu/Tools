package scanner

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"fmt"
)

func CheckRoot() bool {
	return os.Geteuid() == 0
}

func HasNmap() bool {
	_, err := exec.LookPath("nmap")
	return err == nil
}

func PingHost(host string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", host)
	} else {
		cmd = exec.Command("ping", "-c", "1", host)
	}
	return cmd.Run()
}

func GetInput() string {
	var target string
	fmt.Print("Target (x.x.x.x): ")
	fmt.Scanln(&target)

	if target == "" {
		return ""
	}
	return target
}

func PrerequisitesOK(target string) error {
	if !CheckRoot() {
		return errors.New("need root/admin privileges")
	}
	if !HasNmap() {
		return errors.New("nmap is not installed")
	}
	if err := PingHost(target); err != nil {
		return errors.New("target is unreachable")
	}
	return nil
}