package main

import (
	"regexp"
)

const INPUT string = "./input"

var stepExp = regexp.MustCompile(`^move (?P<quantity>\d+) from (?P<source>\d+) to (?P<destination>\d+)$`)

type Step struct {
	Quantity    int
	Source      int
	Destination int
}

type Stack struct {
	Crates []string
}

type Dock struct {
	Stacks    []Stack
	Procedure []Step
}

func main() {
	dock := readManifest(INPUT)

}

func readManifest(manifestFile string) Dock {

}
