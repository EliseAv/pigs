package main

type Podium struct {
	subset  []int
	idxBest int
}

func NewPodium(tracker *Scores) *Podium {
	var podium Podium
	podium.subset = podium.calculateSubset(tracker.duplicate())
	if len(podium.subset) == 0 {
		return nil
	}
	podium.idxBest = podium.findBest()
	return &podium
}

func (Podium) calculateSubset(tracker Scores) []int {
	if len(tracker.scores) == 0 {
		return nil
	}
	wimpScore := tracker.scores[0]
	for i, s := range tracker.scores {
		if s < wimpScore {
			return tracker.scores[:i]
		}
	}
	return tracker.scores
}

func (p *Podium) findBest() int {
	best := 0
	value := 0
	for i, current := range p.subset {
		if value < current {
			value = current
			best = i
		}
	}
	return best
}
