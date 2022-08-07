package main

import (
	"github.com/buger/goterm"
	"time"
)

func TuiKeepPrinting(scores *Scores) {
	started := time.Now()
	goterm.Clear()
	for {
		time.Sleep(time.Second / 6)
		snapshot := scores.duplicate()
		podium := NewPodium(snapshot.duplicate())
		elapsed := time.Since(started)
		podium.tuiPrint(snapshot.plays, elapsed)
	}
}

func (podium Podium) tuiPrint(plays int, elapsed time.Duration) {
	rate := float64(plays) / float64(elapsed.Milliseconds())
	for line := goterm.Height(); line > 0; line-- {
		if line <= len(podium) {
			goterm.MoveCursor(1, line)
			item := podium[line-1]
			_, _ = goterm.Printf("%4d: %d ", item.bet, item.score)
		}
	}
	goterm.MoveCursor(40, 5)
	_, _ = goterm.Printf("%d/ms ", int(rate))
	goterm.Flush()
}
