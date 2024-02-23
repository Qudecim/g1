package main

type Weapon1 struct {
	damage         float64
	criticalChance float64
	criticalDamage float64
}

func newWeapon1(damage float64, criticalChance float64, criticalDamage float64) *Weapon1 {
	return &Weapon1{damage: damage, criticalChance: criticalChance, criticalDamage: criticalDamage}
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
		closest.damage(w.damage)
		return []byte("&w:1:" + string(player.id) + ":" + string(closest.id))
	}

	return nil
}
