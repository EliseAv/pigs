package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"time"
)

type Gui struct {
	started time.Time
	tracker *Scores
	podium  *Podium
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

func (g *Gui) main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Pigs Simulator")
	g.started = time.Now()
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Gui) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Gui) Update() error {
	g.podium = NewPodium(g.tracker)
	return nil
}

func (g *Gui) Draw(screen *ebiten.Image) {
	podium := g.podium
	if podium == nil {
		return
	}

	gas := g.NewAutoScaler(screen, podium)
	for i, val := range podium.subset {
		gas.DrawBar(screen, i, val)
	}

	// Just a placeholder draw function
	elapsed := time.Now().Sub(g.started)
	rate := g.tracker.plays / int(elapsed.Milliseconds())
	label := fmt.Sprintf("%d plays\n%d/ms", g.tracker.plays, rate)
	text.Draw(screen, label, bigFont, 50, 50, colornames.Orange)
}

type GuiAutoScaler struct {
	scaleX, scaleY, bottom float64
}

func (Gui) NewAutoScaler(screen *ebiten.Image, podium *Podium) GuiAutoScaler {
	var gas GuiAutoScaler
	sw, sh := screen.Size()
	gas.bottom = float64(sh)
	pw := len(podium.subset)
	ph := podium.subset[podium.idxBest]
	gas.scaleX = float64(sw) / float64(pw)
	gas.scaleY = gas.bottom / float64(ph)
	return gas
}

func (s GuiAutoScaler) DrawBar(screen *ebiten.Image, i, v int) {
	width := s.scaleX
	x := width * float64(i)
	height := s.scaleY * float64(v)
	y := s.bottom - height
	ib := 0x80 + 0x8*(uint8(i)%2)
	c := color.RGBA{R: ib, G: ib, B: ib, A: 0xff}
	ebitenutil.DrawRect(screen, x, y, width, height, c)
}
