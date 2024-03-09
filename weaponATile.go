package main

import (
	"math"
)

type WeaponATile struct {
	damage         float64
	criticalChance float64
	criticalDamage float64

	angle float64
	x     float64
	y     float64
}

func newWeaponATile(x float64, y float64, damage float64, criticalChance float64, criticalDamage float64, angle float64) *WeaponATile {
	return &WeaponATile{x: x, y: y, damage: damage, criticalChance: criticalChance, criticalDamage: criticalDamage, angle: angle}
}

func (w *WeaponATile) calc(player *Player, game *Game) {
	speed := float64(300)
	atack_range := float64(20)

	speedX := speed * math.Sin(w.angle)
	speedY := speed * math.Cos(w.angle)

	w.x += speedX * game.deltaTime
	w.y += speedY * game.deltaTime

	for zombie, _ := range game.zombies {
		d := distance(zombie.getX(), zombie.getY(), w.x, w.y)
		if d < atack_range {
			zombie.damage(w.damage)
		}
	}
}
