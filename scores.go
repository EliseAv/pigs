package main

type Scores struct {
	scores []int
	plays  int
}

func (scores *Scores) addScore(part Scores) {
	if scores.scores == nil || scores.plays == 0 {
		*scores = part
		return
	}
	if len(part.scores) > len(scores.scores) {
		scores.scores, part.scores = part.scores, scores.scores
	}
	for i, v := range part.scores {
		scores.scores[i] += v
	}
	scores.plays += part.plays
}

func (scores *Scores) addPlay(play []int) {
	// make sure we have room for the play
	places := 0
	for _, val := range play {
		places += val
	}
	if places > len(scores.scores) {
		newScores := make([]int, places)
		copy(newScores, scores.scores)
		scores.scores = newScores
	}

	// calculate scores and accumulate them at the same time
	part := 0
	index := 0
	for _, roll := range play {
		part += roll
		for ; index < part; index++ {
			scores.scores[index] += part
		}
	}
	scores.plays++
}

func (scores Scores) duplicate() Scores {
	destination := Scores{make([]int, len(scores.scores)), scores.plays}
	copy(destination.scores, scores.scores)
	return destination
}
