package main

import (
	"time"
)

type WeaponA struct {
	last_run time.Time
	items    map[*WeaponATile]bool
}

func newWeaponContainer() *WeaponA {
	return &WeaponA{last_run: time.Now(), items: make(map[*WeaponATile]bool)}
}

func (w *WeaponA) calc(player *Player, game *Game) []byte {
	next_run := w.last_run.Add(time.Duration(player.weaponParams.weaponA_timeout))
	var s []byte
	if next_run.Before(time.Now()) {
		weapon := newWeaponATile(player.weaponParams.weaponA_damage, player.weaponParams.weaponA_criticalChance, player.weaponParams.weaponA_criticalDamage)
		s = append(s, weapon.calc(player, game)...)
		w.last_run = time.Now()
	}

	return s
}

func weaponAGetUpgrades() []int {
	return []int{1, 2, 3, 4}
}

func weaponAUpgrade(player *Player, upgrade int) {
	switch upgrade {
	case 1:
		player.weaponParams.weaponA_damage = player.weaponParams.weaponA_damage * 2
	case 2:
		player.weaponParams.weaponA_timeout = int64(float64(player.weaponParams.weaponA_timeout) * 0.2)
	case 3:
		player.weaponParams.weaponA_criticalChance = player.weaponParams.weaponA_criticalChance * 1.1
	case 4:
		player.weaponParams.weaponA_criticalDamage = player.weaponParams.weaponA_criticalDamage * 1.1
	}
}
