package particle

type Particle struct {
	PositionX int
	PositionY int
	VelocityX int
	VelocityY int
	Diameter  int
}

func (p *Particle) Position() []int {
	return []int{
		p.PositionX,
		p.PositionY}
}

func (p *Particle) Velocity() []int {
	return []int{
		p.VelocityX,
		p.VelocityY}
}
