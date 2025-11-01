package particle

type Particle struct {
	PositionX float64
	PositionY float64
	VelocityX int
	VelocityY int
	Diameter  float64
}

func (p *Particle) Position() []float64 {
	return []float64{
		p.PositionX,
		p.PositionY}
}

func (p *Particle) Velocity() []int {
	return []int{
		p.VelocityX,
		p.VelocityY}
}

func (p *Particle) TopEdge() float64 {
	return p.PositionY - (p.Diameter / 2)
}

func (p *Particle) BotEdge() float64 {
	return p.PositionY + (p.Diameter / 2)
}

func (p *Particle) RightEdge() float64 {
	return p.PositionX + (p.Diameter / 2)
}

func (p *Particle) LeftEdge() float64 {
	return p.PositionX - (p.Diameter / 2)
}
