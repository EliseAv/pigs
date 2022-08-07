package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
	"time"
)

type Gui struct {
	numTicks int
	started  time.Time
	duration time.Duration
	tracker  *Scores
}

var bigFont font.Face

func init() {
	fontFamily, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	bigFont, err = opentype.NewFace(fontFamily, &opentype.FaceOptions{Size: 24, DPI: 72, Hinting: font.HintingFull})
	if err != nil {
		log.Fatal(err)
	}
}

func (gui *Gui) main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Pigs Simulator")
	err := ebiten.RunGame(gui)
	if err != nil {
		log.Fatal(err)
	}
}

func (gui *Gui) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (gui *Gui) Update() error {
	gui.numTicks++
	return nil
}

func (gui *Gui) Draw(screen *ebiten.Image) {
	// Just a placeholder draw function
	label := fmt.Sprint(gui.numTicks)
	text.Draw(screen, label, bigFont, 50, 50, colornames.Turquoise)
}
