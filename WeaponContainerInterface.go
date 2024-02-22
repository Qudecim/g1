package main

type WeaponContainerInterface interface {
	calc(*Player, *Game) []byte
}
