package main

type WeaponInterface interface {
	calc(*Player, *Game) []byte
	getUpgrades() []int
	upgrade(int)
}
