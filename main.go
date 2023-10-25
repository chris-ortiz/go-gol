package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math/rand"
)

type Game struct {
	spielfeld        [][]bool
	scale            uint8
	width            uint16
	height           uint16
	columns          uint16
	rows             uint16
	updatesPerSecond uint8
}

func New() *Game {
	g := &Game{}
	g.updatesPerSecond = 10
	g.scale = 5
	g.width = 640
	g.height = 480
	g.columns = g.width / uint16(g.scale)
	g.rows = g.height / uint16(g.scale)
	g.spielfeld = make([][]bool, g.columns)
	for i := range g.spielfeld {
		g.spielfeld[i] = make([]bool, g.rows)
	}

	g.randomlyFillSpielfeld()

	return g
}

func (g *Game) Start() {
	ebiten.SetWindowSize(int(g.width), int(g.height))
	ebiten.SetWindowTitle("Game Of Life")
	ebiten.SetTPS(int(g.updatesPerSecond))

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	g.randomlyFillSpielfeld()
	return nil
}

func (g *Game) randomlyFillSpielfeld() {
	for i := uint16(0); i < g.columns; i++ {
		for j := uint16(0); j < g.rows; j++ {
			g.spielfeld[i][j] = rand.Int()%2 == 0
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for x := uint16(0); x < g.width; x += uint16(g.scale) {
		for y := uint16(0); y < g.height; y += uint16(g.scale) {
			if g.spielfeld[x/uint16(g.scale)][y/uint16(g.scale)] {
				vector.DrawFilledRect(screen, float32(x), float32(y), float32(g.scale), float32(g.scale),
					color.Black, false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.width), int(g.height)
}

func main() {
	game := New()
	game.Start()
}
