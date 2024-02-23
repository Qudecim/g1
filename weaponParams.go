package main

import (
	"fmt"
	"time"
)

type WeaponParams struct {
	weapon1_damage         float64
	weapon1_timeout        int64
	weapon1_criticalChance float64
	weapon1_criticalDamage float64
}

func newWeaponParams() *WeaponParams {
	return &WeaponParams{
		weapon1_damage:         100,
		weapon1_timeout:        int64(time.Second),
		weapon1_criticalChance: 10,
		weapon1_criticalDamage: 10,
	}
}

func (w *WeaponParams) upgrade(player *Player, weapon int, upgrade int) {
	fmt.Println(weapon)
	fmt.Println(upgrade)
	if weapon == 1 {
		switch upgrade {
		case 1:
			w.weapon1_damage = w.weapon1_damage * 2
		case 2:
			w.weapon1_timeout = int64(float64(w.weapon1_timeout) * 0.2)
		case 3:
			w.weapon1_criticalChance = w.weapon1_criticalChance * 1.1
		case 4:
			w.weapon1_criticalDamage = w.weapon1_criticalDamage * 1.1
		}
	}

}
