package main

import "time"

type Weapon1Container struct {
	last_run time.Time
	timeout  int64
	items    map[*Weapon1]bool
}

func newWeaponContainer() *Weapon1Container {
	return &Weapon1Container{last_run: time.Now(), timeout: 1, items: make(map[*Weapon1]bool)}
}

func (w *Weapon1Container) calc(player *Player, game *Game) []byte {
	next_run := w.last_run.Add(time.Second)
	var s []byte
	if next_run.Before(time.Now()) {
		weapon := newWeapon1(0)
		s = append(s, weapon.calc(player, game)...)
		w.last_run = time.Now()
	}

	return s
}
