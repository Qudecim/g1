package main

type Player struct {
	id        []byte
	isDeleted bool

	left  bool
	right bool
	up    bool
	down  bool

	x float64
	y float64

	weapons []WeaponContainerInterface
}

func newPlayer() *Player {
	player := &Player{id: generateId(), x: float64(250), y: float64(250)}
	weapon := newWeaponContainer()
	player.weapons = append(player.weapons, weapon)
	return player
}

func (p *Player) delete() {
	p.isDeleted = true
}

func (p *Player) upgrade(weapon int, upgrade int) {

}
