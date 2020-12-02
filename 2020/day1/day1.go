package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const desiredSum int = 2020

func main() {
	numbers := make([]int, 0)

	// Read input file
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	// Load values into list
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, i)
	}

	var expenseValues []int

	// Find sum of 2 values equal to 2020
	expenseValues = checkSum(numbers, 2, make([]int, 0))
	log.Println(expenseValues, product(expenseValues))

	// Find sum of 3 values equal to 2020
	expenseValues = checkSum(numbers, 3, make([]int, 0))
	log.Println(expenseValues, product(expenseValues))
}

func checkSum(values []int, depth int, currentValues []int) []int {
	var sum int

	for _, currentValue := range values {
		sum = 0
		newValues := append(currentValues, currentValue)
		if depth > 1 {
			var result []int = checkSum(values, depth-1, newValues)
			if result != nil {
				return result
			}
		} else {
			for _, checkValue := range newValues {
				sum = sum + checkValue
			}

			if sum == desiredSum {
				return newValues
			}
		}
	}

	return nil
}

func product(values []int) int {
	var product int = 1
	for _, value := range values {
		product *= value
	}

	return product
}
