package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math/rand"
)

const scale = 5
const width = 640
const height = 480
const columns = width / scale
const rows = height / scale

type Game struct {
	spielfeld [columns][rows]bool
}

func (g *Game) Update() error {
	//g.initializeSpielfeld()
	return nil
}

func New() *Game {
	g := &Game{}
	g.initializeSpielfeld()

	return g
}

func (g *Game) initializeSpielfeld() {
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			g.spielfeld[i][j] = rand.Int()%2 == 0
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for x := 0; x < width; x += scale {
		for y := 0; y < height; y += scale {
			if g.spielfeld[x/scale][y/scale] {
				vector.DrawFilledRect(screen, float32(x), float32(y), scale, scale,
					color.Black, false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Game Of Life")
	ebiten.SetTPS(1)

	game := New()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
