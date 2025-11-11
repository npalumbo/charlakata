package charlakata

//go:generate mockery
type Player interface {
	GetHealth() int
	GetLevel() int
	IsAlive() bool
	DealDamage(target Player, damage int)
	TakeDamage(damage int)
}

type rpgPlayer struct {
	health int
	level  int
}

func (p *rpgPlayer) TakeDamage(damage int) {
	p.health -= damage
}

func (p *rpgPlayer) DealDamage(target Player, damage int) {
	targetHealth := target.GetHealth()
	if damage > targetHealth {
		target.TakeDamage(targetHealth)
	} else {
		target.TakeDamage(damage)
	}
}

func (p *rpgPlayer) GetHealth() int {
	return p.health
}

func (p *rpgPlayer) GetLevel() int {
	return p.level
}

func (p *rpgPlayer) IsAlive() bool {
	return p.health > 0
}

func NewRpgPlayer(health, level int) Player {
	return &rpgPlayer{health: health, level: level}
}
