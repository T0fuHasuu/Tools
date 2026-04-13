Network Vulnerability Scanner: Create a custom scanner (in Python or C++) that uses libraries like
Nmap or raw sockets to probe a network for open ports and known flaws. A vulnerability scanner
gives “a prioritized list of cybersecurity flaws” and helps managers “implement the appropriate
preventative measures” . You could start with a simple host/port scanner and extend it to check
for specific CVEs (for example by integrating a vulnerability database). To make it unique,
containerize your scanner with Docker for easy deployment and add a web dashboard showing
results. (Helps you learn about CVEs, networking, and Linux scripting.)


Setting up a clean and logical folder structure from the very beginning will save you a lot of headaches as your project grows. For a project combining a Go backend, a web dashboard, and Docker, there is a standard way to organize things.

Here is the blueprint for your Network Vulnerability Scanner workspace, along with the exact terminal commands to build it.

### 1. The Project Blueprint

By the end of this setup, your project folder will look exactly like this:

```text
network-vuln-scanner/
├── scanner/
│   └── port_scanner.go    # Holds your concurrent Nmap/raw socket logic
├── cve/
│   └── nvd_api.go         # Handles querying the CVE database
├── web/
│   └── index.html         # The front-end HTML for your dashboard
├── main.go                # The central brain that ties everything together
├── Dockerfile             # The blueprint for containerizing your app
└── go.mod                 # Go's dependency tracker (created automatically)
```

### 2. Step-by-Step Setup Commands

Open your terminal (or Command Prompt/PowerShell if you are on Windows) and run these commands one by one.

**Step 1: Create the main project folder and move into it**
```bash
mkdir network-vuln-scanner
cd network-vuln-scanner
```

**Step 2: Initialize your Go Module**
This tells Go that this folder is a self-contained project. It will create a `go.mod` file to track any external libraries you decide to use later.
```bash
go mod init network-vuln-scanner
```

**Step 3: Create the subdirectories**
These folders will keep your different components organized so your code doesn't become a tangled mess.
```bash
mkdir scanner cve web
```

**Step 4: Create the blank files**
*Note: If you are using Windows Command Prompt, replace `touch` with `type nul >` (e.g., `type nul > main.go`). If you are using PowerShell, Linux, or macOS, `touch` works perfectly.*
```bash
touch main.go
touch scanner/port_scanner.go
touch cve/nvd_api.go
touch web/index.html
touch Dockerfile
```

### 3. What Each Component Does

* **`main.go`:** This is your entry point. When you run your program, it starts here. It will spin up the web server and act as the middleman between your dashboard and your scanning engine.
* **`scanner/port_scanner.go`:** This is where you will write those fast Go routines we talked about earlier. Its only job is to take an IP address, scan it, and return a list of open ports and services.
* **`cve/nvd_api.go`:** This file will take the services found by the scanner, send them to a vulnerability database (like the NVD), and return a list of known flaws (CVEs).
* **`web/index.html`:** This is the face of your application. You will use HTML to display the results in a clean, prioritized list for the user.
* **`Dockerfile`:** We will leave this blank until the very end. Once your app works perfectly on your machine, we will write a few lines in this file to package it up for Docker.

Now that your workspace is fully structured and ready to go, would you like to start by writing the concurrent port scanner logic, or would you prefer to set up the basic web server to see your dashboard in the browser first?




Here is the code to get your scanner up and running! 

In Go, doing a traditional "ICMP Ping" (like typing `ping` in your terminal) requires administrator/root privileges because it uses low-level raw sockets. To keep things simple and avoid needing admin rights every time you run your tool, port scanners usually do a **"TCP Ping."** A TCP Ping simply tries to open a quick connection to a specific port. If the connection succeeds, we know the target is alive and the port is open.

Here is how to set up your files based on the folder structure we just created.

### 1. The Scanner Logic (`scanner/port_scanner.go`)

Open your `scanner/port_scanner.go` file and add this code. This file acts as a module that your main program can call whenever it needs to check a target.

```go
package scanner

import (
	"fmt"
	"net"
	"time"
)

// PingTarget tries to connect to a specific port to see if it is alive.
// Notice that the function name starts with a capital 'P'. 
// In Go, starting with a capital letter makes the function "public" so other files can use it!
func PingTarget(target string, port int) bool {
	// Combine the target and port into the format "IP:PORT" (e.g., "scanme.nmap.org:80")
	address := fmt.Sprintf("%s:%d", target, port)
	
	// Set a 2-second timeout so our script doesn't hang forever if a server drops the packet
	timeout := 2 * time.Second

	// net.DialTimeout acts as our "TCP Ping"
	conn, err := net.DialTimeout("tcp", address, timeout)
	
	if err != nil {
		// If there is an error, the connection failed (port closed or host dead)
		return false
	}

	// If we successfully connected, close the connection cleanly and report success
	conn.Close()
	return true
}
```

### 2. The Main Brain (`main.go`)

Now, open your `main.go` file. This is the script you will actually run. It will import the scanner code you just wrote and use it to test a target.

```go
package main

import (
	"fmt"
	// Import the custom scanner package we created. 
	// Make sure "network-vuln-scanner" matches the name you used in the 'go mod init' step!
	"network-vuln-scanner/scanner"
)

func main() {
	// We will test Nmap's officially authorized scanning target on port 80 (HTTP)
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
```

### 3. How to Run It

Now that your code is saved, make sure you are in the root directory of your project (`network-vuln-scanner`) in your terminal, and run this single command:

```bash
go run main.go
```

If everything is set up correctly, your terminal will reach out across the internet, attempt a TCP handshake with the Nmap test server, and print a `[+] Success!` message.