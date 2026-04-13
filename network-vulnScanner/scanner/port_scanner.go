package scanner

import (
	"fmt"
	"net"
	"time"
)

// Public Function
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