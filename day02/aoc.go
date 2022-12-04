package main

import (
	"errors"
	"fmt"
	"os"
	"petalv/aoc_2022/file"
	"strings"
)

const (
	ROCK    int = 1
	PAPER       = 2
	SCISSOR     = 3
)

const (
	WIN  int = 6
	LOSS     = 0
	DRAW     = 3
)

type Round struct {
	opponent int
	me       int
	outcome  int
}

func main() {
	part := os.Getenv("part")
	input, _ := file.ReadStringArray("input.txt")
	if part == "part2" {
		rounds := mapRoundsWithOutcome(input)
		score := part2(rounds)
		fmt.Printf("%d\n", score)
	} else {
		rounds := mapRounds(input)
		score := part1(rounds)
		fmt.Printf("%d\n", score)
	}
}

func part1(rounds []Round) int {
	var score = 0
	for _, round := range rounds {
		roundScore := scoreRound(round)
		score += roundScore
	}
	return score
}

func part2(rounds []Round) int {
	var score = 0
	for _, round := range rounds {
		round.me = getMeRound(round)
		roundScore := scoreRound(round)
		score += roundScore
	}
	return score
}

func mapRounds(rawRounds []string) []Round {
	var rounds []Round
	for _, str := range rawRounds {
		s := strings.Split(str, " ")
		opponent, _ := getType(s[0])
		me, _ := getType(s[1])
		rounds = append(rounds, Round{opponent: opponent, me: me})
	}
	return rounds
}

func mapRoundsWithOutcome(rawRounds []string) []Round {
	var rounds []Round
	for _, str := range rawRounds {
		s := strings.Split(str, " ")
		opponent, _ := getType(s[0])
		outcome, _ := getOutcomeType(s[1])
		rounds = append(rounds, Round{opponent: opponent, outcome: outcome})
	}
	return rounds
}

func getOutcomeType(t string) (int, error) {
	switch t {
	case "X":
		return -1, nil
	case "Y":
		return 0, nil
	case "Z":
		return 1, nil
	}
	return -127, errors.New("unknown type")
}

func getType(t string) (int, error) {
	switch t {
	case "A":
		return ROCK, nil
	case "X":
		return ROCK, nil
	case "B":
		return PAPER, nil
	case "Y":
		return PAPER, nil
	case "C":
		return SCISSOR, nil
	case "Z":
		return SCISSOR, nil
	}
	return -127, errors.New("unknown type")
}

func scoreRound(round Round) int {
	if round.opponent == SCISSOR && round.me == PAPER {
		return round.me + LOSS
	}
	if round.opponent == SCISSOR && round.me == ROCK {
		return round.me + WIN
	}
	if round.opponent == ROCK && round.me == SCISSOR {
		return round.me + LOSS
	}
	if round.opponent == ROCK && round.me == PAPER {
		return round.me + WIN
	}
	if round.opponent == PAPER && round.me == ROCK {
		return round.me + LOSS
	}
	if round.opponent == PAPER && round.me == SCISSOR {
		return round.me + WIN
	}
	return round.me + DRAW
}

func getMeRound(round Round) int {
	if round.opponent == SCISSOR && round.outcome == -1 {
		return PAPER
	}
	if round.opponent == SCISSOR && round.outcome == 1 {
		return ROCK
	}
	if round.opponent == ROCK && round.outcome == -1 {
		return SCISSOR
	}
	if round.opponent == ROCK && round.outcome == 1 {
		return PAPER
	}
	if round.opponent == PAPER && round.outcome == -1 {
		return ROCK
	}
	if round.opponent == PAPER && round.outcome == 1 {
		return SCISSOR
	}
	return round.opponent
}
