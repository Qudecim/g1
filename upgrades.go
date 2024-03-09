package main

import "math/rand"

func upgrade(player *Player, upgrade int) {
	del := 10 // step beatwen weapon

	weapon := int(upgrade / del)
	ost := upgrade % del

	if weapon == 0 {
		playerUpgrade(player, ost)
	} else {
		player.getWeapon(getWeaponKindById(weapon)).upgrade(ost)
	}
}

func getUpgrades(player *Player) []int {
	del := 10
	upgrades := playerGetUpgrades()

	for _, weapon := range player.weapons {
		ups := weapon.getUpgrades()
		for _, upgrade := range ups {
			upgrades = append(upgrades, upgrade+(del*1))
		}
	}

	var result []int
	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(len(upgrades))
		result = append(result, upgrades[randomIndex])
	}

	return result

}

func getWeaponKindById(id int) string {
	var weaponKind string
	switch id {
	case 1:
		weaponKind = "a"
	}

	return weaponKind
}
