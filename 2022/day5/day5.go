package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	crateMover := flag.String("cratemover", "9000", "which cratemover to run steps with")
	flag.Parse()
	dock := readManifest(INPUT)

	fmt.Println("Starting dock:")
	dock.Print()
	fmt.Print("\n\n")

	fmt.Println("Steps:")
	for _, step := range dock.Procedure {
		if *crateMover == "9000" {
			dock.RunStep9000(step)
		} else if *crateMover == "9001" {
			dock.RunStep9001(step)
		} else {
			log.Fatal("invalid cratemover specified")
		}

	}

	fmt.Println("Resulting dock:")
	dock.Print()
	fmt.Print("\n\n")
}

func readManifest(manifestFile string) Dock {
	readFile, err := os.Open(manifestFile)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	dock := Dock{}
	var stackRows []string
	var steps []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "[") {
			stackRows = append(stackRows, line)
		} else if strings.Contains(line, "move") {
			steps = append(steps, line)
		}
	}

	dock.LoadStacks(stackRows)
	dock.LoadSteps(steps)
	return dock
}

func (d *Dock) LoadStacks(rows []string) {
	for i := (len(rows) - 1); i >= 0; i-- {
		row := rows[i]
		for j := 0; j <= len(row); j += 4 {

			// Check if this is the start of a stack
			if j%4 == 0 {
				stackNum := j / 4
				// Add a new empty stack if this stack doesn't exist
				if len(d.Stacks) <= stackNum {
					d.Stacks = append(d.Stacks, Stack{})
				}

				// If this row for this stack has a crate
				if row[j:j+1] == "[" {
					// Push the crate onto the stack
					d.Stacks[stackNum].Push(row[j+1 : j+2])
				}
			}

		}
	}
}

func (d *Dock) LoadSteps(steps []string) {
	for _, rawStep := range steps {
		match := stepExp.FindStringSubmatch(rawStep)
		step := Step{}
		if len(match) == 4 {
			step.Quantity = GetInt(match[1])
			step.Source = GetInt(match[2])
			step.Destination = GetInt(match[3])
		} else {
			log.Println(match)
		}
		d.Procedure = append(d.Procedure, step)
	}
}

func (d *Dock) RunStep9000(step Step) {
	for c := 0; c < step.Quantity; c++ {
		crate, err := d.Stacks[step.Source-1].Pop()
		if err != nil {
			log.Fatal(err)
		}

		d.Stacks[step.Destination-1].Push(crate)
	}
}

func (d *Dock) RunStep9001(step Step) {
	var crates []string
	for c := 0; c < step.Quantity; c++ {
		crate, err := d.Stacks[step.Source-1].Pop()
		if err != nil {
			log.Fatal(err)
		}

		crates = append(crates, crate)
	}

	for i := len(crates) - 1; i >= 0; i-- {
		d.Stacks[step.Destination-1].Push(crates[i])
	}
}

func (d *Dock) GetMaxStackHeight() int {
	max := 0
	for _, stack := range d.Stacks {
		if len(stack.Crates) > max {
			max = len(stack.Crates)
		}
	}
	return max
}

func (d *Dock) Print() {
	maxStackHeight := d.GetMaxStackHeight()
	for i := (maxStackHeight - 1); i >= 0; i-- {
		for j, stack := range d.Stacks {
			if (len(stack.Crates) - 1) >= i {
				fmt.Printf("[%s]", stack.Crates[i])
			} else {
				fmt.Print("   ")
			}

			if j < (len(d.Stacks) - 1) {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}
	}
}

func (s *Stack) Push(crate string) {
	s.Crates = append(s.Crates, crate)
}

func (s *Stack) Pop() (string, error) {
	if len(s.Crates) > 0 {
		lastCrateIndex := len(s.Crates) - 1
		crate := s.Crates[lastCrateIndex]
		s.Crates = s.Crates[:lastCrateIndex]
		return crate, nil
	} else {
		return "", errors.New("tried to pop from an empty stack of crates")
	}
}

func GetInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
