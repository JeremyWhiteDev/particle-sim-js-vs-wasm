package simulator

import (
	"testing"
)

// TODO generalize this test and increase test cases
func TestParticleGen(t *testing.T) {
	s := New(100, 100, 25)

	if len(s.particles) != 25 {
		t.Fatalf("particle count wrong. expected=%d, got=%d", 25, len(s.particles))
	}

	if s.particles[0].PositionX != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%d", 10, s.particles[0].PositionX)
	}

	if s.particles[0].PositionY != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%d", 10, s.particles[0].PositionY)

	}

	if s.particles[1].PositionX != 30 {
		t.Fatalf("particle position wrong. expected=%d, got=%d", 30, s.particles[1].PositionX)
	}

	if s.particles[1].PositionY != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%d", 10, s.particles[1].PositionY)
	}

	if s.particles[4].PositionX != 90 {
		t.Fatalf("particle position wrong. expected=%d, got=%d", 90, s.particles[4].PositionX)
	}

	if s.particles[4].PositionY != 10 {
		t.Fatalf("particle position wronng. expected=%d, got=%d", 10, s.particles[4].PositionY)
	}

	if s.particles[5].PositionX != 10 {
		t.Fatalf("particle position wrong. expected=%d, got=%d", 10, s.particles[5].PositionX)
	}

	if s.particles[5].PositionY != 30 {
		t.Fatalf("particle position wrong. expected=%d, got=%d", 10, s.particles[5].PositionY)
	}

	// TODO test particle random velocities

}
