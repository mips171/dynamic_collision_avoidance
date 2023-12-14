package main

import (
	"math"
)

type Entity struct {
	Position       Vector
	Velocity       Vector
	Acceleration   Vector
	Size           float64
	FuturePosition Vector
	Priority       int
}

// might combine obstacle and entity into one struct
type Obstacle struct {
	Position Vector
	Size     float64
}

func (e *Entity) PredictFuturePosition(timeStep float64) Vector {
	futureVelocity := e.Velocity.Add(e.Acceleration.Multiply(timeStep))
	futurePosition := e.Position.Add(futureVelocity.Multiply(timeStep))
	return futurePosition
}

func (e *Entity) LimitVelocity(maxSpeed float64) {
	speed := math.Sqrt(e.Velocity.X*e.Velocity.X + e.Velocity.Y*e.Velocity.Y)
	if speed > maxSpeed {
		e.Velocity.X = e.Velocity.X / speed * maxSpeed
		e.Velocity.Y = e.Velocity.Y / speed * maxSpeed
	}
}

func (e *Entity) InformOthers(entities []Entity) {
	for i := range entities {
		// later Entity can keep a table of entities near it, and use memoization to avoid checking all entities
		if &entities[i] != e {
			distance := e.Position.Distance(entities[i].Position)
			reactionDistance := e.Size * e.Size * 4 // Adjusted for dynamic distance
			if distance < reactionDistance {
				entities[i].ReactToInformation(e.FuturePosition, e.Priority)
			}
		}
	}
}

func (e *Entity) ReactToInformation(futurePosition Vector, priority int) {
	diff := e.Position.Subtract(futurePosition)
	distance := diff.Distance(futurePosition)

	priorityFactor := 1.0

	if priority > e.Priority {
		// TODO investigate if this should be dynamic
		priorityFactor = 1.5
	}

	// dynamic reaction distance
	reactionDistance := e.Size*e.Size*4 + e.Velocity.Magnitude()*2

	if distance < reactionDistance {
		steerStrength := math.Min(10/distance, 1)
		steer := diff.Normalize().Multiply(steerStrength)
		e.Velocity = e.Velocity.Add(steer)

		// Hit the brakes if collision imminent
		// TODO investigate if this should be dynamic
		if distance < e.Size*e.Size*4 {
			brakingForce := e.Velocity.Multiply(-1).Normalize().Multiply(priorityFactor)
			e.Velocity = e.Velocity.Add(brakingForce)
		}
	}

	// Slow down if the other entity is faster regardless of priority
	// TODO investigate if this should be tied to priority
	e.LimitVelocity(0.25)
}

func (e *Entity) MoveTowards(target Vector, envWidth, envHeight int) {
	direction := Vector{X: target.X - e.Position.X, Y: target.Y - e.Position.Y}
	mag := math.Sqrt(direction.X*direction.X + direction.Y*direction.Y)

	// apply direction to velocity only if not too close to the objective
	// this will form a ring around the objective
	if mag > e.Size*5 {
		direction.X /= mag
		direction.Y /= mag
		e.Velocity = e.Velocity.Add(direction)
	}

	e.Position = e.Position.Add(e.Velocity)
}

// Separation is a force that pushes entities away from each other
// It should increase with the number of vehciles.
// TODO Working on a ratio now.
func (e *Entity) Separation(entities []Entity, separationDistance float64) Vector {
	var steer Vector
	count := 0.0

	for _, other := range entities {
		distance := e.Position.Distance(other.Position)
		if distance > 0 && distance < separationDistance {
			diff := Vector{
				X: e.Position.X - other.Position.X,
				Y: e.Position.Y - other.Position.Y,
			}
			mag := math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
			diff.X /= mag
			diff.Y /= mag
			diff.X /= distance
			diff.Y /= distance
			steer = steer.Add(diff)
			count++
		}
	}

	// TODO refine this. It will also need to be changed for 3D
	SEPARATION_FORCE := float64(len(entities)) * 4

	if count > 0 {
		steer.X /= count
		steer.Y /= count
		steer.X *= SEPARATION_FORCE
		steer.Y *= SEPARATION_FORCE
	}

	return steer
}

// AvoidObstacles returns a force to avoid obstacles
func (e *Entity) AvoidObstacles(obstacles []Obstacle, safeDistance float64) Vector {
	var avoidanceForce Vector
	for _, obstacle := range obstacles {
		distance := e.Position.Distance(obstacle.Position)


		if distance < safeDistance {
			// Calculate a force vector that points away from the obstacle
			awayFromObstacle := e.Position.Subtract(obstacle.Position)
			// The closer the obstacle, the stronger the force
			forceStrength := 1 - (distance / safeDistance)
			avoidanceForce = avoidanceForce.Add(awayFromObstacle.Multiply(forceStrength))
		}
	}
	return avoidanceForce
}

func (e *Entity) GroupBehavior(entities []Entity) {
	var align, separate, cohesion Vector
	var count int

	for _, other := range entities {
		distance := e.Position.Distance(other.Position)
		if distance > 0 && distance < e.Size*e.Size*4 {
			align = align.Add(other.Velocity)
			separate = separate.Add(e.Position.Subtract(other.Position).Normalize().Divide(distance))
			cohesion = cohesion.Add(other.Position)
			count++
		}
	}

	// TODO refine these weights
	// TODO may need to be dynamic based on formation
	alignWeight := 1.0
	separateWeight := 3.0
	cohesionWeight := 0.3

	if count > 0 {
		align = align.Divide(float64(count)).Normalize()
		separate = separate.Divide(float64(count))
		cohesion = cohesion.Divide(float64(count)).Subtract(e.Position).Normalize()

		// Apply these forces to e.Velocity with some weights
		e.Velocity = e.Velocity.Add(align.Multiply(alignWeight))
		e.Velocity = e.Velocity.Add(separate.Multiply(separateWeight))
		e.Velocity = e.Velocity.Add(cohesion.Multiply(cohesionWeight))
	}
}
