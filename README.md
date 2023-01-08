# PortScanner

This project is a simple port scanner written in Go. It takes in a hostname as an argument and performs a port scan on it. It also has a list of common port names and numbers that it uses to scan.

## Customizing the scan

You can customize the ports that are being scanned by modifying the ports.GetPort function. This function returns a map of port numbers to port names that the program uses to perform the scan. You can add or remove entries from this map to include or exclude specific ports in the scan.

## Functions

`runner()`: Takes in a hostname as an argument and performs a port scan on it. It uses the scan function to scan each port and a WaitGroup to manage the goroutines.

`ErrorHandler()`: Takes in an error and logs it if it is not nil.

`checkingHost()`: Takes in a hostname and returns its IP addresses.

`scan()`: Performs a port scan on a specific host and port. It tries to establish a connection to the host and port and if the connection is successful, it prints the port and name.
