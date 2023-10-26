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
	initialNoise     int
	wrapAround       bool
}

const live = true
const dead = false

func New() *Game {
	g := &Game{}
	g.updatesPerSecond = 8
	g.scale = 5
	g.width = 1200
	g.height = 800
	g.columns = g.width / uint16(g.scale)
	g.rows = g.height / uint16(g.scale)
	g.spielfeld = g.newSpielfeld()
	g.wrapAround = true

	// lower value = higher noise = more initial pixels
	g.initialNoise = 2

	g.randomlyFillSpielfeld()

	return g
}

func (g *Game) newSpielfeld() [][]bool {
	spielfeld := make([][]bool, g.columns)
	for i := range spielfeld {
		spielfeld[i] = make([]bool, g.rows)
	}

	return spielfeld
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

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		g.spielfeld[x/int(g.scale)][y/int(g.scale)] = live
	} else {
		newSpielfeld := g.newSpielfeld()
		for i := 0; i < int(g.columns); i++ {
			for j := 0; j < int(g.rows); j++ {
				c := g.countNeighbors(i, j)
				newSpielfeld[i][j] = g.getNewCellState(g.spielfeld[i][j] == live, c)
			}
		}

		g.spielfeld = newSpielfeld
	}

	return nil
}

/*
Any live cell with fewer than two live neighbours dies, as if by underpopulation.
Any live cell with two or three live neighbours lives on to the next generation.
Any live cell with more than three live neighbours dies, as if by overpopulation.
Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/
func (g *Game) getNewCellState(liveCell bool, c uint8) bool {
	if liveCell && c < 2 {
		return dead
	} else if liveCell && (c == 2 || c == 3) {
		return live
	} else if liveCell && c > 3 {
		return dead
	} else if !liveCell && c == 3 {
		return live
	} else {
		return dead
	}
}

func (g *Game) countNeighbors(i int, j int) uint8 {
	var count uint8
	count = 0

	if g.wrapAround {
		if j == 0 {
			j = int(g.rows - 2)
		} else if j == int(g.rows)-1 {
			j = 1
		}
		if i == 0 {
			i = int(g.columns - 2)
		} else if i == int(g.columns)-1 {
			i = 1
		}
	} else if j == 0 || i == 0 || i == int(g.columns)-1 || j == int(g.rows)-1 {
		return count
	}

	if g.spielfeld[i][j+1] {
		count++
	}
	if g.spielfeld[i][j-1] {
		count++
	}
	if g.spielfeld[i-1][j] {
		count++
	}
	if g.spielfeld[i+1][j] {
		count++
	}
	if g.spielfeld[i+1][j-1] {
		count++
	}
	if g.spielfeld[i+1][j+1] {
		count++
	}
	if g.spielfeld[i-1][j+1] {
		count++
	}
	if g.spielfeld[i-1][j-1] {
		count++
	}

	return count
}

func (g *Game) randomlyFillSpielfeld() {
	for i := uint16(0); i < g.columns; i++ {
		for j := uint16(0); j < g.rows; j++ {
			g.spielfeld[i][j] = rand.Int()%g.initialNoise == 0
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
	g := New()
	g.Start()
}
