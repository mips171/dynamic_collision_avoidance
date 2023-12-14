package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	env := Environment{
		Width:  800,
		Height: 600,
	}

	env.Initialize(24, 10)

	game := &Game{env: env}
	ebiten.SetWindowSize(env.Width, env.Height)
	ebiten.SetWindowTitle("Entities and Obstacles")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	env Environment
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// Set the objective to the mouse cursor's position
		g.env.Objective = NewVector(float64(x), float64(y))
	}
	g.env.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.env.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.env.Width, g.env.Height
}
