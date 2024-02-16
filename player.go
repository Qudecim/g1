package main

type Player struct {
	id    []byte
	left  bool
	right bool
	up    bool
	down  bool

	x float64
	y float64

	weapons []WeaponInterface
}

func newPlayer() *Player {
	player := &Player{id: generateId(), x: float64(250), y: float64(250)}
	weapon := newWeapon1(0)
	player.weapons = append(player.weapons, weapon)
	return player
}
