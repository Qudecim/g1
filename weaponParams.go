package main

import (
	"time"
)

type WeaponParams struct {
	weaponA                bool
	weaponA_damage         float64
	weaponA_timeout        int64
	weaponA_criticalChance float64
	weaponA_criticalDamage float64
}

func newWeaponParams() *WeaponParams {
	return &WeaponParams{
		weaponA:                true,
		weaponA_damage:         100,
		weaponA_timeout:        int64(time.Second),
		weaponA_criticalChance: 10,
		weaponA_criticalDamage: 10,
	}
}
