package main

type GameRules struct {
	running  bool
	dieSides int
	lose     []bool
	partials chan Scores
}

func NewGameRules(dieSides int, lose ...int) *GameRules {
	slice := make([]bool, dieSides)
	for _, value := range lose {
		slice[value] = true
	}
	return &GameRules{dieSides: dieSides, lose: slice}
}

func (game *GameRules) Simulate(parallelism int) *Scores {
	game.partials = make(chan Scores, 10)
	for i := 0; i < parallelism; i++ {
		go game.keepGenerating()
	}
	tracker := &Scores{}
	go game.accumulate(tracker)
	game.running = true
	return tracker
}

func (game *GameRules) keepGenerating() {
	var play []int
	die := NewDie(game.dieSides)
	for game.running {
		// accumulate a number of scores before emitting them
		scores := Scores{}
		for i := 1000000; i > 0; i-- {
			roll := die.Roll()
			if game.lose[roll-1] {
				scores.addPlay(play)
				play = play[:0] // reuse the same array
			} else {
				play = append(play, roll)
			}
		}
		game.partials <- scores
	}
}

func (game *GameRules) accumulate(tracker *Scores) {
	for score := range game.partials {
		tracker.addScore(score)
	}
}

func (game *GameRules) Stop() {
	game.running = false
	close(game.partials)
}
