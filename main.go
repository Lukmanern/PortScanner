package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

// commonly used ports
var common = map[int]string{
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
}


func main() {
	start := time.Now()
	runner("facebook.com")
	finish := time.Since(start)
	fmt.Printf("Scan duration: %s", finish)
}

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// runner function takes in a hostname 
// as an argument and performs a port scan on it
func runner(host string) {
	// create a WaitGroup to manage the goroutines
	var wg sync.WaitGroup 
	// get the IP addresses of the host
	ip, err := checkingHost(host) 
	ErrorHandler(err)
	fmt.Println("total IP :", len(ip), "->", ip)

	for port, name := range common {
		// add 1 to the WaitGroup counter
		wg.Add(1)
		// launch a goroutine to scan the current port
		go scan(host, strconv.Itoa(port), name, &wg) 
	}
	// wait for all goroutines to finish
	wg.Wait() 
}

// checkingHost function takes in a hostname and returns its IP addresses
func checkingHost(host string) ([]net.IP, error) {
	// get the IP addresses of the host
	ip, err := net.LookupIP(host) 
	ErrorHandler(err)

	return ip, nil
}

// scan function performs a port scan on a specific host and port
func scan(host, port, name string, wg *sync.WaitGroup) {
	// decrement the WaitGroup counter
	defer wg.Done()
	// try to establish a connection to the host and port
	_, err := net.DialTimeout("tcp", host+":"+port, 1*time.Second)
	if err == nil {
		// if the connection was successful, 
		// print the port and name
		fmt.Println(port, name)
	}
}
