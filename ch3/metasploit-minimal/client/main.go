package main

import (
	"log"
	"os"
	"github.com/niclaslind/black-hat-go/ch3/metasploit-minimal/rpc"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("Missing required enviroment variable MSFHOST of MSFPASS")
		msf, err := rpc.New(host, user, pass)
	}
}
