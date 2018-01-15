package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	START_LINE = "Ethernet address:"
)

func main() {
	// First, read the file
	file, err := os.Open("test-logging")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), START_LINE) {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

/* Each stanza is a 

Ethernet address: ac3743490f76
	Incoming total 2 packets, 104 bytes; 0 IP packets
	Outgoing total 3 packets, 156 bytes; 0 IP packets
	Average rates: 0.00 kbits/s incoming, 0.10 kbits/s outgoing
	Last 5-second rates: 0.00 kbits/s incoming, 0.20 kbits/s outgoing

/*
func GetUsageStanzn(scanner *bufio.Scanner) {



}
