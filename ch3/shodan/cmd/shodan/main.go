package main

import (
	shodan "github.com/niclaslind/black-hat-go/ch3/shodan/shodan"
	"log"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New()
}
