package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

const input string = "./input"

func main() {
	allElves := readCalories(input)

	bestElf := 0
	mostCalories := 0

	for i, v := range allElves {
		log.Printf("Elf %d has %d calories.", i, v)

		if v > mostCalories {
			bestElf = i
			mostCalories = v
		}
	}

	sort.Ints(allElves)
	totalElves := len(allElves)
	topThreeCalories := allElves[totalElves-1] + allElves[totalElves-2] + allElves[totalElves-3]

	log.Printf("Best elf was %d, who carried %d calories.", bestElf, mostCalories)
	log.Printf("Total calories carried by the top 3 elves is: %d", topThreeCalories)
}

func readCalories(caloriesFile string) []int {
	var caloriesList []int
	readFile, err := os.Open(caloriesFile)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	elfCalories := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			caloriesList = append(caloriesList, elfCalories)
			elfCalories = 0
		} else {
			lineCalories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			elfCalories += lineCalories
		}
	}

	return caloriesList
}
