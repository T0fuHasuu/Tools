Network Vulnerability Scanner: Create a custom scanner (in Python or C++) that uses libraries like
Nmap or raw sockets to probe a network for open ports and known flaws. A vulnerability scanner
gives “a prioritized list of cybersecurity flaws” and helps managers “implement the appropriate
preventative measures” . You could start with a simple host/port scanner and extend it to check
for specific CVEs (for example by integrating a vulnerability database). To make it unique,
containerize your scanner with Docker for easy deployment and add a web dashboard showing
results. (Helps you learn about CVEs, networking, and Linux scripting.)


Using `subprocess` is an excellent choice. It is built directly into Python (so no installation is required) and is incredibly powerful. It effectively allows your Python script to open an invisible terminal, type a command, hit "Enter," and read exactly what comes back.



Before Python 3.5, there were many confusing ways to do this (like `os.system` or `subprocess.call`). Today, there is only **one** function you need to memorize for 95% of your tasks: `subprocess.run()`.

Here are the absolute basics and the most critical "must-know" concepts for using it safely and effectively.

---

### 1. The Anatomy of `subprocess.run()`
When you use `run()`, your Python script will pause, wait for the terminal command to finish completely, and then hand you a `CompletedProcess` object containing the results.

Here is the most standard, useful way to call it:

```python
import subprocess

command = ["ping", "-c", "4", "google.com"] # Linux/Mac ping format

result = subprocess.run(
    command, 
    capture_output=True, 
    text=True, 
    check=False
)

print(result.stdout)
```

### 2. The Must-Know Arguments
Let's break down exactly what those arguments in the code above are doing. These are your primary levers for controlling the process:

* **The Command List:** Notice that `["ping", "-c", "4", "google.com"]` is a list of strings, not one long string like `"ping -c 4 google.com"`. This is intentional and safer (more on this below).
* **`capture_output=True`:** If you do not include this, the output of your command will just print directly to your screen, and Python won't be able to save it to a variable. Setting this to `True` tells Python to grab the text so you can use it.
* **`text=True`:** By default, `subprocess` returns raw bytes (e.g., `b'Hello\n'`). Setting `text=True` automatically decodes those bytes into a normal, readable Python string (e.g., `'Hello\n'`).
* **`check=True` (or False):** If a command fails in the terminal (like trying to read a file that doesn't exist), it throws an error. If `check=True`, Python will crash and raise a `CalledProcessError`. If `check=False` (the default), Python will keep running and simply note the failure in the result object.

### 3. Reading the Results (stdout, stderr, returncode)
When `run()` finishes, the `result` object holds three critical pieces of information:

```python
import subprocess

result = subprocess.run(["ls", "/fake_directory"], capture_output=True, text=True)

# 1. returncode: 0 means success. Anything else (1, 2, 255, etc.) means failure.
print(f"Exit Code: {result.returncode}") 

# 2. stdout: The standard output (what you normally see when things go right)
print(f"Standard Output: {result.stdout}")

# 3. stderr: The standard error (the error messages the terminal spits out)
print(f"Error Message: {result.stderr}")
```

### 4. CRITICAL SECURITY WARNING: `shell=True`
You might see tutorials online telling you to write commands as a single string and use `shell=True`. 

```python
# DANGEROUS WAY
user_input = "127.0.0.1"
subprocess.run(f"ping {user_input}", shell=True)
```

**Never do this if you are accepting user input.** If you use `shell=True`, Python passes the entire string directly to the operating system's shell (like Bash or Command Prompt). If a malicious user types `127.0.0.1; rm -rf /` instead of an IP address, the shell will ping the IP, and then immediately delete every file on your computer. This is called **Shell Injection**, and it is one of the most common vulnerabilities in cybersecurity tools.

Always use a list of strings (like `["ping", "127.0.0.1"]`) and leave `shell=False` (which is the default). This forces the OS to treat `127.0.0.1; rm -rf /` as a very long, strange IP address rather than a secondary command, safely preventing the attack.

### 5. Timeouts (Preventing Infinite Freezes)
Sometimes tools hang. If you are running an Nmap scan or a brute-force tool and it gets stuck, your Python script will freeze forever waiting for it. You can prevent this using the `timeout` parameter.

```python
import subprocess

try:
    # If this takes longer than 10 seconds, kill it!
    subprocess.run(["sleep", "100"], timeout=10)
except subprocess.TimeoutExpired:
    print("The command took too long and was terminated.")
```

Are you planning to run commands where you just wait for the final output, or do you need to run long-lasting background tasks (like a continuous Wi-Fi monitor) where you need to read the output line-by-line while it's still running?