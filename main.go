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

var HOST string = "facebook.com"

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

func runner(host string) {
	var wg sync.WaitGroup
	var err error
	_, err = checkingHost(host)
	ErrorHandler(err)

	for port, name := range common {
		wg.Add(1)
		go scan(strconv.Itoa(port), name, &wg)
	}
	wg.Wait()
}

func checkingHost(host string) ([]net.IP, error) {
	var ip []net.IP
	ip, err := net.LookupIP(host)
	ErrorHandler(err)

	return ip, nil
}

func scan(port, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := net.DialTimeout("tcp", HOST+":"+port, 1*time.Second)
	if err == nil {
		fmt.Println(port, name)
	}
}