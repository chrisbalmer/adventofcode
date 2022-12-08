package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT string = "./input"

type Command struct {
	Executable string
	Argument   string
	Output     []string
}

type File struct {
	Name   string
	Parent *Directory
	Size   int
}

type Directory struct {
	Name        string
	Parent      *Directory
	Directories []*Directory
	Files       []File
}

func main() {
	cmds := loadCommands(INPUT)

	root := runCommands(cmds)

	log.Println(root.SolvePartOne())
	free := 70000000 - root.GetSize()
	goal := 30000000 - free
	log.Printf("Goal: %d", goal)
	log.Println(solvePartTwo(root, goal, 0))
}

func runCommands(cmds []Command) Directory {
	root := Directory{
		Name: "/",
	}

	var dir *Directory
	for _, cmd := range cmds {
		switch cmd.Executable {
		case "cd":
			if cmd.Argument == "/" {
				dir = &root
			} else if cmd.Argument == ".." {
				dir = dir.Parent
			} else {
				found, err := dir.FindSubDirectory(cmd.Argument)
				if err != nil {
					log.Fatal(err)
				}
				dir = found
			}
		case "ls":
			dir.LoadContents(cmd.Output)
		}
	}

	return root
}

func printCommands(cmds []Command) {
	for i, cmd := range cmds {
		log.Printf("Command: %d, Executable: %s, Argument: %s",
			i, cmd.Executable, cmd.Argument)
		for _, output := range cmd.Output {
			fmt.Println(output)
		}
	}
}

func loadCommands(historyFile string) []Command {
	readFile, err := os.Open(historyFile)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var cmds []Command
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if string(line[0]) == "$" {
			parseCommand(line, &cmds)
		} else {
			o := &cmds[len(cmds)-1].Output
			*o = append(*o, line)
		}
	}
	return cmds
}

func parseCommand(line string, cmds *[]Command) {
	curCmd := Command{}
	chunks := strings.Split(line, " ")
	curCmd.Executable = chunks[1]
	if len(chunks) == 3 {
		curCmd.Argument = chunks[2]
	}
	*cmds = append(*cmds, curCmd)
}

func (d *Directory) LoadContents(output []string) {
	for _, item := range output {
		chunks := strings.Split(item, " ")
		if chunks[0] == "dir" {
			child := Directory{
				Name:   chunks[1],
				Parent: d,
			}
			d.Directories = append(d.Directories, &child)
		} else {
			size, err := strconv.Atoi(chunks[0])
			if err != nil {
				log.Fatal(err)
			}
			file := File{
				Name:   chunks[1],
				Parent: d,
				Size:   size,
			}
			d.Files = append(d.Files, file)
		}
	}
}

func (d *Directory) FindSubDirectory(n string) (*Directory, error) {
	for _, dir := range d.Directories {
		if dir.Name == n {
			return dir, nil
		}
	}
	return nil, errors.New("couldn't find directory")
}

func (d *Directory) GetSize() int {
	size := 0

	for _, dir := range d.Directories {
		size += dir.GetSize()
	}

	for _, file := range d.Files {
		size += file.Size
	}

	return size
}

func (d *Directory) SolvePartOne() int {
	size := 0
	partial := 0

	for _, dir := range d.Directories {
		partial += dir.GetSize()
	}

	for _, file := range d.Files {
		partial += file.Size
	}

	if partial <= 100000 {
		size += partial
	}

	for _, dir := range d.Directories {
		size += dir.SolvePartOne()
	}

	return size
}

func solvePartTwo(start Directory, goal int, candidate int) int {
	for _, dir := range start.Directories {
		size := dir.GetSize()
		if size >= goal && (size < candidate || candidate == 0) {
			candidate = size
		}
		candidate = solvePartTwo(*dir, goal, candidate)
	}

	return candidate
}
