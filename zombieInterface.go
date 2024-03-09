package main

type ZombieInterface interface {
	getId() []byte
	damage(float64)
	getX() float64
	getY() float64
	isDeleted() bool
	getPonts() int
	isRight() bool
	setRight(bool)
	setX(float64)
	setY(float64)
	getKind() string
	getSpeed() float64
}

func newZombie(zombie_type string, id []byte, x float64, y float64) ZombieInterface {
	var zombie ZombieInterface
	if zombie_type == "a" {
		return newZombieA(id, x, y)
	} else if zombie_type == "b" {
		return newZombieB(id, x, y)
	}
	return zombie
}
