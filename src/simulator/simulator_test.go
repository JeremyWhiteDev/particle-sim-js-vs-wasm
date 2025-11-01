package simulator

import (
	"particlesim/particle"
	"testing"
)

// TODO generalize this test and increase test cases
func TestParticleGen(t *testing.T) {
	s := New(100, 100, 25)

	if len(s.particles) != 25 {
		t.Fatalf("particle count wrong. expected=%d, got=%d", 25, len(s.particles))
	}

	if s.particles[0].PositionX != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%f", 10, s.particles[0].PositionX)
	}

	if s.particles[0].PositionY != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%f", 10, s.particles[0].PositionY)
	}

	if s.particles[1].PositionX != 30 {
		t.Fatalf("particle position wrong. expected=%d, got=%f", 30, s.particles[1].PositionX)
	}

	if s.particles[1].PositionY != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%f", 10, s.particles[1].PositionY)
	}

	if s.particles[4].PositionX != 90 {
		t.Fatalf("particle position wrong. expected=%d, got=%f", 90, s.particles[4].PositionX)
	}

	if s.particles[4].PositionY != 10 {
		t.Fatalf("particle position wronng. expected=%d, got=%f", 10, s.particles[4].PositionY)
	}

	if s.particles[5].PositionX != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%f", 10, s.particles[5].PositionX)
	}

	if s.particles[5].PositionY != 30 {
		t.Fatalf("particle position wrong. expected=%d, got=%f", 10, s.particles[5].PositionY)
	}

}

func TestCollision(t *testing.T) {

	/*
		+---+---+---+---+---+---+---+---+---+---+
		|   |   |   |   |   |   |   |   |   |   |
		+---+---+---+---+---+---+---+---+---+---+
		|   |   | A |   |   |   |   |   |   |   |
		+---+---+---+---+---+---+---+---+---+---+
		|   |   |   | C |   |   | B |   |   |   |
		+---+---+---+---+---+---+---+---+---+---+
		|   |   |   |   |   |   |   |   |   | D |
		+---+---+---+---+---+---+---+---+---+---+
		|   |   |   |   |   |   |   |   |   |   |
		+---+---+---+---+---+---+---+---+---+---+
	*/

	s := Simulator{
		canvasWidth:  10,
		canvasHeight: 5,
		particles: []*particle.Particle{
			{ // particle A collides with particle C
				PositionX: 3,
				PositionY: 2,
				VelocityX: 1,
				VelocityY: -1,
				Diameter:  1,
			},
			{ // paticle B is by itself
				PositionX: 6,
				PositionY: 3,
				VelocityX: 5,
				VelocityY: 5,
				Diameter:  1,
			},
			{ // particle C collides with particle A
				PositionX: 4,
				PositionY: 3,
				VelocityX: 0,
				VelocityY: 1,
				Diameter:  1,
			},
			{ // particle D collides with canvas boundary
				PositionX: 10,
				PositionY: 4,
				VelocityX: 1,
				VelocityY: 1,
				Diameter:  1,
			},
		},
	}

	s.updatePartices()

	if s.particles[0].VelocityX >= 0 {
		t.Fatalf("particle A velocity wrong. expected= <0, got=%d", s.particles[0].VelocityX)
	}

	if s.particles[0].VelocityY <= 0 {
		t.Fatalf("particle A velocity wrong. expected= >0, got=%d", s.particles[0].VelocityY)
	}

	if s.particles[1].VelocityX != 5 {
		t.Fatalf("particle B velocity wrong. expected=5, got=%d", s.particles[1].VelocityX)
	}

	if s.particles[1].VelocityY != 5 {
		t.Fatalf("particle B velocity wrong. expected=5, got=%d", s.particles[1].VelocityY)
	}

	if s.particles[2].VelocityX >= 0 {
		t.Fatalf("particle C velocity wrong. expected= <0, got=%d", s.particles[2].VelocityX)
	}

	if s.particles[2].VelocityY >= 0 {
		t.Fatalf("particle C velocity wrong. expected= <0, got=%d", s.particles[2].VelocityY)
	}

	if s.particles[3].VelocityX >= 0 {
		t.Fatalf("particle D velocity wrong. expected= <0, got=%d", s.particles[3].VelocityX)
	}

	if s.particles[3].VelocityY >= 0 {
		t.Fatalf("particle D velocity wrong. expected= <0, got=%d", s.particles[3].VelocityY)
	}

}
