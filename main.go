package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func parseIPRange(cidr string) ([]string, error) {

	ip, ipnet, err := net.ParseCIDR(cidr)

	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

// Function to increment the IP for the for loop in the CIDR range parser

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func createPortRange(portRange []string) ([]int, error) {

	if len(portRange) < 2 {
		port, err := strconv.Atoi(portRange[0])

		if err != nil {
			fmt.Println(err)
		}
		return []int{port}, nil
	}
	// Convert string cli input to integers for building lists
	range_start, err := strconv.Atoi(portRange[0])
	if err != nil {
		fmt.Println(err)
	}
	range_end, err := strconv.Atoi(portRange[1])
	if err != nil {
		fmt.Println(err)
	}
	var portList []int
	for port := range_start; port <= range_end; port++ {

		portList = append(portList, port)

	}

	return portList, nil

}

func convertPortListToString(portList []int) ([]string, error) {
	var stringPortList []string
	for _, port := range portList {
		stringPortList = append(stringPortList, strconv.Itoa(port))
	}
	return stringPortList, nil
}

func doScan(target string, portList []int, wg *sync.WaitGroup) {

	defer wg.Done()

	var portOutPut []int

	for _, port := range portList {
		// Convert string ports Ints back to strings to handles connection and printing status
		conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(port), time.Duration(1)*time.Second)
		if err != nil {
			fmt.Println(err)
		}
		// Convert string ports Ints back to strings to handles connection and printing status
		if conn == nil {
			fmt.Println("Could not connect to " + target + " on port " + strconv.Itoa(port))
		} else {
			fmt.Println("Connected to " + target + " on port " + strconv.Itoa(port))
			portOutPut = append(portOutPut, port)
		}
	}
	stringPortOutput, _ := convertPortListToString(portOutPut)
	if portOutPut != nil {
		fmt.Println("Host: " + target + " Ports: " + strings.Join(stringPortOutput, "/TCP, ") + "/TCP")
	}
	portOutPut = nil

}

func main() {
	// Get IP/CIDR range from args
	if len(os.Args) < 3 {
		fmt.Println("Usage portscanner <IP or IP range> <port or port range>")
		fmt.Println("Example: portscanner 192.168.2.3 1-1024")
		os.Exit(0)
	}
	ip := string(os.Args[1])
	port := string(os.Args[2])

	portRange := strings.Split(port, "-")

	portList, err := createPortRange(portRange)
	if err != nil {
		fmt.Println(err)
	}

	ips, err := parseIPRange(ip)
	if err != nil {
		fmt.Println(err)
	}

	if len(ips) == 0 {
		ips = append(ips, ip)
	}

	// Spawn Goroutine per target and let the scans run async
	var wg sync.WaitGroup
	for _, target := range ips {
		wg.Add(1)
		go doScan(target, portList, &wg)
	}
	wg.Wait()

}
