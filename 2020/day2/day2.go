package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	var masterList []map[string]string

	// Read input file
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}
	originalList := string(content)

	masterList = processOriginalList(originalList)

	var valid int = 0
	for _, passwordSet := range masterList {
		count := strings.Count(passwordSet["password"], passwordSet["letter"])
		min, _ := strconv.Atoi(passwordSet["min"])
		max, _ := strconv.Atoi(passwordSet["max"])
		if count >= min && count <= max {
			log.Println(passwordSet)
			valid++
		}
	}
	log.Println(valid)

	valid = 0
	for _, passwordSet := range masterList {
		var found bool = false
		min, _ := strconv.Atoi(passwordSet["min"])
		max, _ := strconv.Atoi(passwordSet["max"])
		characters := strings.Split(passwordSet["password"], "")

		if characters[min-1] == passwordSet["letter"] {
			found = true
		}

		if characters[max-1] == passwordSet["letter"] && found {
			found = false
		} else if characters[max-1] == passwordSet["letter"] {
			found = true
		}

		if found {
			valid++
			log.Println(passwordSet)
		}

	}
	log.Println(valid)
}

func processOriginalList(originalList string) []map[string]string {
	masterList := make([]map[string]string, 0)

	scanner := bufio.NewScanner(strings.NewReader(originalList))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		passwordSet := make(map[string]string)
		passwordSet["policy"] = strings.TrimSpace(parts[0])
		passwordSet["password"] = strings.TrimSpace(parts[1])

		parts = strings.Split(passwordSet["policy"], " ")
		passwordSet["letter"] = parts[1]

		parts = strings.Split(parts[0], "-")
		passwordSet["min"] = parts[0]
		passwordSet["max"] = parts[1]

		masterList = append(masterList, passwordSet)
	}

	return masterList
}
