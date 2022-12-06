package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT string = "./input"

type AssignmentPair struct {
	ElfA Assignment
	ElfB Assignment
}

type Assignment struct {
	Start int
	End   int
}

func main() {
	assignmentPairs := readAssignmentPairs(INPUT)
	totalRedundantPairs := 0
	totalOverlappingPairs := 0
	for i, assignmentPair := range assignmentPairs {
		if assignmentPair.IsRedundant() {
			log.Printf("Pair %d have a redundancy. ElfA is %d-%d, ElfB is %d-%d.", i,
				assignmentPair.ElfA.Start, assignmentPair.ElfA.End,
				assignmentPair.ElfB.Start, assignmentPair.ElfB.End)
			totalRedundantPairs++
		}

		if assignmentPair.HasOverlap() {
			log.Printf("Pair %d have overlap. ElfA is %d-%d, ElfB is %d-%d.", i,
				assignmentPair.ElfA.Start, assignmentPair.ElfA.End,
				assignmentPair.ElfB.Start, assignmentPair.ElfB.End)
			totalOverlappingPairs++
		}
	}
	log.Printf("Total redundant pairs found: %d", totalRedundantPairs)
	log.Printf("Total overlapping pairs found: %d", totalOverlappingPairs)

}

func readAssignmentPairs(assignmentsFile string) []AssignmentPair {
	readFile, err := os.Open(assignmentsFile)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var assignmentPairs []AssignmentPair
	for fileScanner.Scan() {
		line := fileScanner.Text()
		assignmentPair := AssignmentPair{}
		pair := strings.Split(line, ",")
		assignmentPair.ElfA.AssignElf(pair[0])
		assignmentPair.ElfB.AssignElf(pair[1])

		assignmentPairs = append(assignmentPairs, assignmentPair)
	}

	return assignmentPairs
}

func (ap *AssignmentPair) IsRedundant() bool {
	if (ap.ElfA.Start >= ap.ElfB.Start && ap.ElfA.End <= ap.ElfB.End) ||
		(ap.ElfA.Start <= ap.ElfB.Start && ap.ElfA.End >= ap.ElfB.End) {
		return true
	} else {
		return false
	}
}

func (ap *AssignmentPair) HasOverlap() bool {
	if (ap.ElfA.Start >= ap.ElfB.Start && ap.ElfA.Start <= ap.ElfB.End) ||
		(ap.ElfA.Start <= ap.ElfB.Start && ap.ElfA.End >= ap.ElfB.Start) ||
		(ap.ElfA.End >= ap.ElfB.Start && ap.ElfA.End <= ap.ElfB.End) ||
		(ap.ElfA.Start <= ap.ElfB.End && ap.ElfA.End >= ap.ElfB.End) {
		return true
	} else {
		return false
	}
}

func (a *Assignment) AssignElf(assignment string) {
	ends := strings.Split(assignment, "-")

	start, err := strconv.Atoi(ends[0])
	if err != nil {
		log.Fatal(err)
	}
	a.Start = start

	end, err := strconv.Atoi(ends[1])
	if err != nil {
		log.Fatal(err)
	}
	a.End = end
}
