/*

Simulation of the game Pigs.

Inspired by a Numberphile video.
https://www.youtube.com/watch?v=ULhRLGzoXQ0

*/

package main

func main() {
	rules := NewGameRules(6, 1)
	tracker := rules.Simulate(6)
	TuiKeepPrinting(tracker)
}
