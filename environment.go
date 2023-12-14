package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Environment struct {
	Entities  []Entity
	Obstacles []Obstacle
	Width     int
	Height    int
	Objective Vector
}

func (env *Environment) Update() {

	for i := range env.Entities {
		separationForce := env.Entities[i].Separation(env.Entities, env.Entities[i].Size*env.Entities[i].Size*4)

		// Add obstacle avoidance
		obstacleAvoidanceForce := env.Entities[i].AvoidObstacles(env.Obstacles, env.Entities[i].Size)

		// Combine the separation force with obstacle avoidance
		totalForce := separationForce.Add(obstacleAvoidanceForce)

		// Apply the combined force
		env.Entities[i].Velocity = env.Entities[i].Velocity.Add(totalForce)

		// Move towards the objective
		env.Entities[i].MoveTowards(env.Objective, env.Width, env.Height)

		// Limit the velocity
		env.Entities[i].LimitVelocity(0.5)
	}

	// Check for collisions
	for i := 0; i < len(env.Entities); i++ {
		for j := i + 1; j < len(env.Entities); j++ {
			if env.Entities[i].Position.Distance(env.Entities[j].Position) < (env.Entities[i].Size + env.Entities[j].Size) {
				// Remove the collided entities - they are out of service now
				log.Printf("BANG! Collision between entity %d and entity %d\n", i, j)
				env.Entities = append(env.Entities[:i], env.Entities[i+1:]...)
				j--
			}
		}
	}
}

// Constants for priority calculation
const (
	meanPriority      = 10
	standardDeviation = 5
	minPriority       = 1
	maxPriority       = 20
)

func (env *Environment) Initialize(numEntities, numObstacles int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numEntities; i++ {
		priority := int(rand.NormFloat64()*standardDeviation + meanPriority)

		if priority < minPriority {
			priority = minPriority
		} else if priority > maxPriority {
			priority = maxPriority
		}

		env.Entities = append(env.Entities, Entity{
			Position: NewVector(rand.Float64()*float64(env.Width), rand.Float64()*float64(env.Height)),
			Velocity: NewVector(rand.Float64()*0.1-0.05, rand.Float64()*0.1-0.05),
			Size:     5,
			Priority: priority,
		})
	}

	// Add obstacles
	for i := 0; i < numObstacles; i++ {
		obSize := 5
		if i%2 == 0 {
			obSize = 10
		}
		env.Obstacles = append(env.Obstacles, Obstacle{
			Position: NewVector(rand.Float64()*float64(env.Width), rand.Float64()*float64(env.Height)),
			Size:     float64(obSize),
		})
	}
}

func (env *Environment) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 10, G: 10, B: 10, A: 255})

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	green := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	color := red

	for _, e := range env.Entities {
		if e.Priority > meanPriority+standardDeviation {
			color = blue
		}
		ebitenutil.DrawRect(screen, e.Position.X, e.Position.Y, e.Size, e.Size, color)
	}
	// Draw obstacles
	for _, e := range env.Obstacles {
		ebitenutil.DrawCircle(screen, e.Position.X, e.Position.Y, e.Size, green)
	}
}
