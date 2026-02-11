package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// gradBanner attempts to connect to a port  and read the service's signature
func grabBanner(ip string, port int) {
	address := fmt.Sprintf("%s:%d", ip, port)

	// Establish a connection with a 5-second timeout
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Printf("[-] Port %d: Could not connect\n", port)
		return
	}
	defer conn.Close()

	// Set a deadline for reading the banner
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	// For some services (like HTTP), we might need to send a probe but many (SSH, FTP, SMTP) send the banner immediately upon connection.

	fmt.Printf("[+] Port %d: Connected! Waiting for banner...\n", port)

	banner, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("[-] Port %d: Connection closes or no banner received\n", port)
		return
	}

	fmt.Printf("[*] Banner from port %d: %s", port, banner)
}

func main() {
	// Example target: public DNS or a know server.
	// Be careful: always test on targets you have permission to scan.
	targetIP := "scanme.nmap.org" // Professional testing target provided by Nmap
	targetPort := 22              // SSH port usually gives a clear banner

	fmt.Printf("Starting Banner Grabbing os %s...\n\n", targetIP)
	grabBanner(targetIP, targetPort)
}
