package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Round struct {
	OpponentMove int
	S1           Strategy
	S2           Strategy
}

type Strategy struct {
	YourMove int
	Score    int
	Winner   string
}

const input string = "./input"

const SCORE_TIE = 3
const WINNER_TIE = "tie"

const SCORE_WIN = 6
const WINNER_YOU = "you"

const SCORE_LOSE = 0
const WINNER_OPPONENT = "opponent"

func main() {
	strategy := readStrategy(input)
	firstScore := 0
	secondScore := 0

	for round, roundData := range strategy {
		log.Printf("Round %d, first strategy: Winner was %s, your score was %d.", round, roundData.S1.Winner, roundData.S1.Score)
		firstScore += roundData.S1.Score

		log.Printf("Round %d, second strategy: Winner was %s, your score was %d.", round, roundData.S2.Winner, roundData.S2.Score)
		secondScore += roundData.S2.Score
	}

	log.Printf("Your total score for the first strategy was %d.", firstScore)
	log.Printf("Your total score for the second strategy was %d.", secondScore)
}

func readStrategy(strategyFile string) []Round {
	var strategy []Round
	readFile, err := os.Open(strategyFile)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		moves := strings.Split(line, " ")
		currentRound := Round{}
		for _, move := range moves {
			switch move {
			case "A":
				currentRound.OpponentMove = 1
			case "B":
				currentRound.OpponentMove = 2
			case "C":
				currentRound.OpponentMove = 3
			case "X":
				currentRound.S1.YourMove = 1
			case "Y":
				currentRound.S1.YourMove = 2
			case "Z":
				currentRound.S1.YourMove = 3
			}
		}
		currentRound.DetermineS1Winner()
		currentRound.DetermineS2Winner()
		currentRound.ScoreStrategies()
		strategy = append(strategy, currentRound)
	}

	return strategy
}

func (r *Round) DetermineS1Winner() {
	if r.OpponentMove == r.S1.YourMove {
		r.S1.Winner = WINNER_TIE
	} else if r.OpponentMove == 3 && r.S1.YourMove == 1 {
		r.S1.Winner = WINNER_YOU
	} else if r.OpponentMove == 1 && r.S1.YourMove == 3 {
		r.S1.Winner = WINNER_OPPONENT
	} else if r.OpponentMove < r.S1.YourMove {
		r.S1.Winner = WINNER_YOU
	} else {
		r.S1.Winner = WINNER_OPPONENT
	}
}

func (r *Round) DetermineS2Winner() {
	switch r.S1.YourMove {
	case 1:
		r.S2.Winner = WINNER_OPPONENT
		r.S2.YourMove = r.OpponentMove - 1
	case 2:
		r.S2.Winner = WINNER_TIE
		r.S2.YourMove = r.OpponentMove
	case 3:
		r.S2.Winner = WINNER_YOU
		r.S2.YourMove = r.OpponentMove + 1
	}

	if r.S2.YourMove == 0 {
		r.S2.YourMove = 3
	} else if r.S2.YourMove == 4 {
		r.S2.YourMove = 1
	}
}

func (r *Round) ScoreStrategies() {
	strategies := []*Strategy{&r.S1, &r.S2}

	for _, strategy := range strategies {
		switch strategy.Winner {
		case WINNER_TIE:
			strategy.Score = strategy.YourMove + SCORE_TIE
		case WINNER_YOU:
			strategy.Score = strategy.YourMove + SCORE_WIN
		case WINNER_OPPONENT:
			strategy.Score = strategy.YourMove + SCORE_LOSE
		}
	}
}
