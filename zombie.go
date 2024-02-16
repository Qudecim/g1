package main

type Zombie struct {
	id []byte
	x  float64
	y  float64
	hp float64
}

func newZombie(id []byte, x float64, y float64) *Zombie {
	return &Zombie{id: id, x: x, y: y, hp: 100}
}

func (z *Zombie) damage(damage float64) bool {
	z.hp -= damage
	if z.hp < 0 { // check it in game
		return true
	}

	return false
}
