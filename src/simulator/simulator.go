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

func (s *Simulator) generateParticles(particleCount int) {
	// grid dimension reresents both x/y of the grid.
	gridDim := math.Ceil(math.Sqrt(float64(particleCount)))

	cellWidth := math.Floor((float64(s.canvasWidth) / gridDim))
	// particleSize is
	particleDiameter := int(math.Floor((float64(s.canvasWidth) / gridDim) * 0.8))

	// The distance to offest the particle so that it is "centered" in each cell
	particleCenterOffset := int(math.Floor((float64(s.canvasWidth) / gridDim) * 0.5))

	// for each cell, I need to "center" in cell.

	currentCellCoord := []int{
		0,
		0,
	}

	for i := range particleCount {
		positionX := (currentCellCoord[0] * int(cellWidth)) + particleCenterOffset
		positionY := (currentCellCoord[1] * int(cellWidth)) + particleCenterOffset

		p := &particle.Particle{
			PositionX: positionX,
			PositionY: positionY,
			Diameter:  particleDiameter,
			VelocityX: randomVelocity(),
			VelocityY: randomVelocity(),
		}

		fmt.Println(i, positionX, "positionx")
		fmt.Println(i, positionY, "positionY")
		fmt.Println(currentCellCoord[0], "current x coord")
		fmt.Println(cellWidth, "cellWidth")
		fmt.Println(s.canvasWidth, "canvasWidth")

		fmt.Println(currentCellCoord[0]*int(cellWidth) < s.canvasWidth, "")

		// progress current cell coord
		if (currentCellCoord[0]+1)*int(cellWidth) < s.canvasWidth {
			currentCellCoord[0] = currentCellCoord[0] + 1
		} else {
			currentCellCoord[0] = 0
			currentCellCoord[1] = currentCellCoord[1] + 1
		}
		s.particles = append(s.particles, p)

	}

}

func randomVelocity() int {
	min, max := -10, 10

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
