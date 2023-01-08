package main

import (
	"PortScanner/ports"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	runner("facebook.com")
	finish := time.Since(start)
	fmt.Println("\nScan duration: " + finish.String())
}

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// runner func takes in a hostname 
// as an argument and performs a port scan on it
func runner(host string) {
	// create a WaitGroup to manage the goroutines
	// and get the IP addresses of the host
	var wg sync.WaitGroup
	ip, err := checkingHost(host) 
	ErrorHandler(err)
	fmt.Println("total IP :", len(ip), "->", ip)

	ports := ports.GetPort()
	for port, name := range ports {
		// add 1 to the WaitGroup counter and
		// launch a goroutine to scan the current port
		wg.Add(1)
		go scan(host, strconv.Itoa(port), name, &wg) 
	}
	// wait for all goroutines to finish
	wg.Wait() 
}

// checkingHost func takes in a hostname 
// and returns its IP addresses 
// (the ip addresses is can/ possible more than one)
func checkingHost(host string) ([]net.IP, error) {
	// get the IP addresses of the host
	ip, err := net.LookupIP(host) 
	ErrorHandler(err)
	return ip, nil
}

// scan func performs a port scan on a specific host and port
func scan(host, port, name string, wg *sync.WaitGroup) {
	// decrement the WaitGroup counter and
	// try to establish a connection 
	// to the host and port
	defer wg.Done()
	_, err := net.DialTimeout("tcp", host+":"+port, 1*time.Second)
	if err == nil {
		// if the connection was successful, 
		// print the port and name
		fmt.Println(port, name)
	}
}
