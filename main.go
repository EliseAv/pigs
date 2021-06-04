/*

Simulation of the game Pigs.

Inspired by a Numberphile video.
https://www.youtube.com/watch?v=ULhRLGzoXQ0

*/

package main

func main() {
	const numberOfGoroutines = 6
	scores := make(chan Scores, 10)
	for i := 0; i < numberOfGoroutines; i++ {
		go keepGenerating(scores)
	}
	var tracker Scores
	go tracker.accumulate(scores)
	tracker.keepPrinting()
}

func keepGenerating(channel chan Scores) {
	die := NewDie(6)
	play := []int{}
	for {
		// accumulate a number of scores before emitting them
		scores := Scores{}
		for i := 0; i < 1000000; i++ {
			roll := die.Roll()
			if roll == 1 {
				scores.addPlay(play)
				play = play[:0] // reuse the same array
			} else {
				play = append(play, roll)
			}
		}
		channel <- scores
	}
}
