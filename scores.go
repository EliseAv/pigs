package main

type Scores struct {
	scores []int
	plays  int
}

func (s *Scores) addScore(part Scores) {
	if s.scores == nil || s.plays == 0 {
		*s = part
		return
	}
	if len(part.scores) > len(s.scores) {
		s.scores, part.scores = part.scores, s.scores
	}
	for i, v := range part.scores {
		s.scores[i] += v
	}
	s.plays += part.plays
}

func (s *Scores) addPlay(play []int) {
	// make sure we have room for the play
	places := -len(s.scores)
	for _, val := range play {
		places += val
	}
	if places > 0 {
		s.scores = append(s.scores, make([]int, places)...)
	}

	// calculate scores and accumulate them at the same time
	part := 0
	index := 0
	for _, roll := range play {
		part += roll
		for ; index < part; index++ {
			s.scores[index] += part
		}
	}
	s.plays++
}

func (s Scores) duplicate() Scores {
	destination := Scores{make([]int, len(s.scores)), s.plays}
	copy(destination.scores, s.scores)
	return destination
}
