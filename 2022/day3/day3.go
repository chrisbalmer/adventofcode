package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"unicode"
)

const INPUT string = "./input"
const PRIORITIES string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Rucksack struct {
	All          string
	CompartmentA string
	CompartmentB string
	SharedItem   rune
}

func main() {
	rucksacks := readRucksacks(INPUT)

	priorityTotal := 0
	for i, rucksack := range rucksacks {
		log.Printf("=================================")
		log.Printf("Rucksack %d:", i)
		log.Printf("All: %s", rucksack.All)
		log.Printf("Compartment A: %s", rucksack.CompartmentA)
		log.Printf("Compartment B: %s", rucksack.CompartmentB)
		log.Printf("Shared Item: %q", rucksack.SharedItem)

		priority, err := GetPriority(rucksack.SharedItem)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Priority of Shared Item: %d\n\n", priority)
		priorityTotal += priority
	}

	log.Printf("=================================")
	log.Printf("Priority Total: %d\n\n", priorityTotal)
}

func readRucksacks(rucksackPackOutFile string) []Rucksack {
	readFile, err := os.Open(rucksackPackOutFile)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var rucksacks []Rucksack
	for fileScanner.Scan() {
		line := fileScanner.Text()
		itemCount := len(line)

		rucksack := Rucksack{}
		rucksack.All = line // Just for debugging
		rucksack.CompartmentA = line[0:(itemCount / 2)]
		rucksack.CompartmentB = line[(itemCount / 2):itemCount]
		rucksack.FindSharedItem()
		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}

func (r *Rucksack) FindSharedItem() {
	for _, letterA := range r.CompartmentA {
		for _, letterB := range r.CompartmentB {
			if letterA == letterB {
				r.SharedItem = letterA
				return
			}
		}
	}
}

func GetPriority(r rune) (int, error) {
	for i, letter := range PRIORITIES {
		if r == letter {
			return (i + 1) + 26, nil
		} else if r == unicode.ToLower(letter) {
			return (i + 1), nil
		}
	}

	return -1, errors.New("unable to match shared item to a priority")
}
