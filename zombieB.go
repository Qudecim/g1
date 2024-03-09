package main

type ZombieB struct {
	id         []byte
	is_deleted bool
	kind       string

	x  float64
	y  float64
	hp float64

	speed float64

	points   int
	is_right bool
}

func newZombieB(id []byte, x float64, y float64) *ZombieB {
	return &ZombieB{kind: "b", id: id, x: x, y: y, hp: 10, points: 10, speed: 3}
}

func (z *ZombieB) damage(damage float64) {
	z.hp -= damage
	if z.hp < 0 { // check it in game
		z.delete()
	}
}

func (z *ZombieB) delete() {
	z.is_deleted = true
}

func (z *ZombieB) getId() []byte {
	return z.id
}

func (z *ZombieB) getX() float64 {
	return z.x
}

func (z *ZombieB) getY() float64 {
	return z.y
}

func (z *ZombieB) isDeleted() bool {
	return z.is_deleted
}

func (z *ZombieB) getPonts() int {
	return z.points
}

func (z *ZombieB) isRight() bool {
	return z.is_right
}

func (z *ZombieB) setRight(isRight bool) {
	z.is_right = isRight
}

func (z *ZombieB) setX(x float64) {
	z.x = x
}

func (z *ZombieB) setY(y float64) {
	z.y = y
}

func (z *ZombieB) getKind() string {
	return z.kind
}

func (z *ZombieB) getSpeed() float64 {
	return z.speed
}
