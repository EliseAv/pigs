package main

import (
	"time"

	"github.com/buger/goterm"
)

type Scores struct {
	scores []int
	plays  int
}

func (sum *Scores) addScore(part Scores) {
	if sum.scores == nil || sum.plays == 0 {
		*sum = part
		return
	}
	if len(part.scores) > len(sum.scores) {
		sum.scores, part.scores = part.scores, sum.scores
	}
	for i, v := range part.scores {
		sum.scores[i] += v
	}
	sum.plays += part.plays
}

func (sum *Scores) addPlay(play []int) {
	// make sure we have room for the play
	places := 0
	for _, val := range play {
		places += val
	}
	if places > len(sum.scores) {
		newScores := make([]int, places)
		copy(newScores, sum.scores)
		sum.scores = newScores
	}
	
	// calculate scores and accumulate them at the same time
	part := 0
	index := 0
	for _, roll := range play {
		part += roll
		for ; index < part; index++ {
			sum.scores[index] += part
		}
	}
	sum.plays++
}

func (tracker *Scores) accumulate(channel chan Scores) {
	for score := range channel {
		tracker.addScore(score)
	}
}

func (tracker *Scores) keepPrinting() {
	goterm.Clear()
	for {
		time.Sleep(time.Second / 6)
		snapshot := tracker.duplicate()
		NewPodium(snapshot.duplicate()).print(snapshot.plays)
	}
}

func (source Scores) duplicate() Scores {
	destination := Scores{make([]int, len(source.scores)), source.plays}
	copy(destination.scores, source.scores)
	return destination
}
