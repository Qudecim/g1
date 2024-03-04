package main

type Player struct {
	id        []byte
	isDeleted bool
	client    *Client

	left               bool
	right              bool
	up                 bool
	down               bool
	direction_is_right bool

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

func playerGetUpgrades() []int {
	return []int{1, 2, 3, 4}
}

func playerUpgrade(player *Player, upgrade int) {
	switch upgrade {
	case 1:
		player.maxSpeed = player.maxSpeed * 2
	case 2:
		player.maxHP = player.maxHP + 10
	case 3:
		player.evasion = player.evasion + 1
	case 4:
		player.armor = player.armor + 1
	}
}
