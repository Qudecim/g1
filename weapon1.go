package main

type Weapon1 struct {
}

func newWeapon1(weapon_type int) *Weapon1 {
	return &Weapon1{}
}

func (w *Weapon1) calc(player *Player, game *Game) {
	var closest *Zombie
	var closestRange float64 = 1000
	for zombie, _ := range game.zombies {
		d := distance(zombie.x, zombie.y, player.x, player.y)
		if d < closestRange {
			closest = zombie
			closestRange = d
		}
	}

	if closest != nil {
		r := closest.damage(20)
		if r {
			delete(game.zombies, closest)
		}
	}
}
