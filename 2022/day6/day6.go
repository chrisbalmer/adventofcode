package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const INPUT string = "./input"

func main() {
	readFile, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	counter := 0
	var packetMarker string
	var messageMarker string
	packetMarkerFound := false
	messageMarkerFound := false
	for fileScanner.Scan() {
		r := fileScanner.Text()
		counter++
		if !packetMarkerFound {
			packetMarkerFound = CheckPacketMarker(&packetMarker, r, counter)
		}
		if !messageMarkerFound {
			messageMarkerFound = CheckMessageMarker(&messageMarker, r, counter)
		}
	}
}

func CheckPacketMarker(marker *string, r string, counter int) bool {
	m := *marker
	if len(m) < 4 {
		m += r
	} else {
		m = m[1:] + r
	}

	if len(m) == 4 && AllUnique(m) {
		fmt.Println(m)
		fmt.Println(counter)
		return true
	}
	*marker = m
	return false
}

func CheckMessageMarker(marker *string, r string, counter int) bool {
	m := *marker
	if len(m) < 14 {
		m += r
	} else {
		m = m[1:] + r
	}

	if len(m) == 14 && AllUnique(m) {
		fmt.Println(m)
		fmt.Println(counter)
		return true
	}
	*marker = m
	return false
}

func AllUnique(s string) bool {
	for idx, char := range s {
		for i := idx + 1; i < len(s); i++ {
			if char == rune(s[i]) {
				return false
			}
		}
	}
	return true
}
