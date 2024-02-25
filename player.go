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

	hp       float64
	maxSpeed float64
	maxHP    float64
	evasion  float64
	armor    float64
}

func newPlayer() *Player {
	player := &Player{
		id:           generateId(),
		x:            float64(250),
		y:            float64(250),
		weapons:      make(map[int]WeaponContainerInterface),
		weaponParams: newWeaponParams(),
		hp:           10,
		maxSpeed:     2,
		maxHP:        10,
		evasion:      1,
		armor:        1,
	}
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
	switch upgrade {
	case 1:
		p.maxSpeed = p.maxSpeed * 2
	case 2:
		p.maxHP = p.maxHP + 10
	case 3:
		p.evasion = p.evasion + 1
	case 4:
		p.armor = p.armor + 1
	}
}
