package main

import (
	"sort"
)

type Podium []struct {
	bet   int
	score int
}

func NewPodium(tracker Scores) Podium {
	podium := make(Podium, len(tracker.scores))
	for i, v := range tracker.scores {
		podium[i].bet = i + 1
		podium[i].score = v
	}
	sort.Sort(podium)
	return podium
}

func (podium Podium) Len() int {
	return len(podium)
}

func (podium Podium) Less(i, j int) bool {
	if podium[i].score != podium[j].score {
		return podium[i].score > podium[j].score
	} else {
		return podium[i].bet < podium[j].bet
	}
}

func (podium Podium) Swap(i, j int) {
	podium[i], podium[j] = podium[j], podium[i]
}
