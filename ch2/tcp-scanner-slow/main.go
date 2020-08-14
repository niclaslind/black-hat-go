package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		fmt.Printf("trying to open %d\n", i)
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("runned into trouble %e\n", err)
			// port is closed or filtered
			continue
		}

		if err := conn.Close(); err != nil {
			fmt.Printf("Problem when trying to close connection %e", err)
		}

		fmt.Printf("%d open\n", i)
	}
}
