package main

import "math/rand"

func upgrade(player *Player, upgrade int) {
	del := 10 // step beatwen weapon

	weapon := int(upgrade / del)
	ost := upgrade % del

	switch weapon {
	case 0:
		playerUpgrade(player, ost)
	case 1:
		weaponAUpgrade(player, ost)
	}

}

func getUpgrades(player *Player) []int {
	del := 10
	upgrades := playerGetUpgrades()

	if player.weaponParams.weaponA {
		ups := weaponAGetUpgrades()
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
