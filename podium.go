package main

import (
	"sort"
	"time"

	"github.com/buger/goterm"
)

type Podium []struct {
	bet   int
	score int
}

var started = time.Now()

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

func (podium Podium) print(plays int) {
	elapsed := time.Since(started)
	rate := float64(plays) / float64(elapsed.Milliseconds())
	for line := goterm.Height(); line > 0; line-- {
		if line <= len(podium) {
			goterm.MoveCursor(1, line)
			item := podium[line-1]
			goterm.Printf("%4d: %d ", item.bet, item.score)
		}
	}
	goterm.MoveCursor(40, 5)
	goterm.Printf("%d/ms ", int(rate))
	goterm.Flush()
}
