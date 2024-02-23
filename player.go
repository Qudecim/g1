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

	weapons      map[int]WeaponContainerInterface
	weaponParams *WeaponParams
}

func newPlayer() *Player {
	player := &Player{id: generateId(), x: float64(250), y: float64(250), weapons: make(map[int]WeaponContainerInterface), weaponParams: newWeaponParams()}
	weapon := newWeaponContainer()
	player.weapons[1] = weapon
	return player
}

func (p *Player) delete() {
	p.isDeleted = true
}

func (p *Player) upgrade(weapon int, upgrade int) {
	if upgrade == 0 {
		p.upgrade_player(upgrade)
	} else {
		p.weaponParams.upgrade(p, weapon, upgrade)
	}
}

func (p *Player) upgrade_player(upgrade int) {

}
