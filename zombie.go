package main

type Zombie struct {
	id        []byte
	isDeleted bool

	x  float64
	y  float64
	hp float64

	points             int
	direction_is_right bool
}

func newZombie(id []byte, x float64, y float64) *Zombie {
	return &Zombie{id: id, x: x, y: y, hp: 100, points: 10}
}

func (z *Zombie) delete() {
	z.isDeleted = true
}

func (z *Zombie) damage(damage float64) {
	z.hp -= damage
	if z.hp < 0 { // check it in game
		z.delete()
	}
}
