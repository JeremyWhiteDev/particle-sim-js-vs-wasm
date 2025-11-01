package simulator

import (
	"fmt"
	"math"
	"math/rand/v2"
	"particlesim/particle"
)

type Simulator struct {
	canvasWidth  int
	canvasHeight int
	particles    []*particle.Particle
}

const (
	_ int = iota // iota creates an auto incrementing set of values in this block, much like an ENUM in other languages. LOWEST will have the int value 1, EQUALS 2, etc.
	NEGATIVE
	POSITIVE
	RANDOM
)

func (s *Simulator) generateParticles(particleCount int) {
	// grid dimension reresents both x/y of the grid.
	gridDim := math.Ceil(math.Sqrt(float64(particleCount)))

	cellWidth := math.Floor((float64(s.canvasWidth) / gridDim))
	// particleSize is
	particleDiameter := (float64(s.canvasWidth) / gridDim) * 0.8

	// The distance to offest the particle so that it is "centered" in each cell
	particleCenterOffset := (float64(s.canvasWidth) / gridDim) * 0.5

	// for each cell, I need to "center" in cell.

	currentCellCoord := []float64{
		0,
		0,
	}

	for i := range particleCount {
		positionX := (currentCellCoord[0] * cellWidth) + particleCenterOffset
		positionY := (currentCellCoord[1] * cellWidth) + particleCenterOffset

		p := &particle.Particle{
			PositionX: positionX,
			PositionY: positionY,
			Diameter:  particleDiameter,
			VelocityX: randomVelocity(RANDOM),
			VelocityY: randomVelocity(RANDOM),
		}

		fmt.Println(i, p.VelocityX, p.VelocityY)

		// fmt.Println(i, positionX, "positionx")
		// fmt.Println(i, positionY, "positionY")
		// fmt.Println(currentCellCoord[0], "current x coord")
		// fmt.Println(cellWidth, "cellWidth")
		// fmt.Println(s.canvasWidth, "canvasWidth")

		// fmt.Println(currentCellCoord[0]*int(cellWidth) < s.canvasWidth, "")

		// progress current cell coord
		if (currentCellCoord[0]+1)*cellWidth < float64(s.canvasWidth) {
			currentCellCoord[0] = currentCellCoord[0] + 1
		} else {
			currentCellCoord[0] = 0
			currentCellCoord[1] = currentCellCoord[1] + 1
		}
		s.particles = append(s.particles, p)

	}

}

func randomVelocity(direction int) int {
	var min, max int

	switch direction {
	case NEGATIVE:
		min = -10
		max = 0
	case POSITIVE:
		min = 0
		max = 10
	case RANDOM:
		min = -10
		max = 10
	}

	rangeSize := max - min + 1
	return rand.IntN(rangeSize) + min
}

func New(canvasWidth, canvasHeight, particleCount int) *Simulator {
	sim := &Simulator{
		canvasWidth:  canvasWidth,
		canvasHeight: canvasHeight,
	}

	sim.generateParticles(particleCount)

	return sim
}

func (s *Simulator) updatePartices() {
	for _, p := range s.particles {
		s.handleParticleCollision(p)
	}
	// TODO consolidate to one loop
	// for _, p := range s.particles {
	// 	s.handleParticleCollision(p)
	// }
}

// TODO there's a couple naive items with this implementation:
// 1. if particles touch, but have velocities exactly adjacent to each other (should pass by each other) those particles will still collide.
// 2. there is no accounting for correct physics during these collisions. Particles just go off in random directions that are opposite from their previous direction

func (s *Simulator) handleParticleCollision(p *particle.Particle) {

	fmt.Println(p.RightEdge(), "right edge")
	fmt.Println(float64(s.canvasWidth), "right edge")
	// fmt.Println(currParticle.RightEdge(), "curr particle right edge")
	// fmt.Println(p.LeftEdge(), "p left edge")

	if p.RightEdge() >= float64(s.canvasWidth) ||
		p.LeftEdge() <= 0 ||
		p.TopEdge() <= 0 ||
		p.BotEdge() >= float64(s.canvasHeight) {
		p.VelocityX = randomVelocity(getOppositeSign(p.VelocityX))
		p.VelocityY = randomVelocity(getOppositeSign(p.VelocityY))
		return
	}

	for _, currParticle := range s.particles {

		// We want to ignore our current particle
		if currParticle == p {
			continue
		}
		// fmt.Println("---", i, "---")
		// fmt.Println(currParticle.RightEdge(), "curr particle right edge")
		// fmt.Println(p.LeftEdge(), "p left edge")

		xOverlap := !(currParticle.RightEdge() < p.LeftEdge() || p.RightEdge() < currParticle.LeftEdge())

		yOverlap := !(currParticle.BotEdge() < p.TopEdge() || p.BotEdge() < currParticle.TopEdge())

		fmt.Println(xOverlap, yOverlap, "xOverlap, yOverlap")

		if xOverlap && yOverlap {

			// fmt.Println(randomVelocity(getOppositeSign(p.VelocityX)), "random veloxity x")
			// fmt.Println(randomVelocity(getOppositeSign(p.VelocityY)), "random veloxity Y")
			p.VelocityX = randomVelocity(getOppositeSign(p.VelocityX))
			p.VelocityY = randomVelocity(getOppositeSign(p.VelocityY))
			break
		}
	}

}

func getOppositeSign(i int) int {
	if i >= 0 {
		return NEGATIVE
	} else {
		return POSITIVE
	}
}
