package main

type Weapon1 struct {
}

func newWeapon1(weapon_type int) *Weapon1 {
	return &Weapon1{}
}

func (w *Weapon1) calc(player *Player, game *Game) []byte {
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
		closest.damage(20)
		return []byte("&w:1:" + string(player.id) + ":" + string(closest.id))
	}

	return nil
}
