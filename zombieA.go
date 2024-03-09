package main

type ZombieA struct {
	id         []byte
	is_deleted bool
	kind       string

	x     float64
	y     float64
	hp    float64
	speed float64

	points   int
	is_right bool
}

func newZombieA(id []byte, x float64, y float64) *ZombieA {
	return &ZombieA{kind: "a", id: id, x: x, y: y, hp: 10, points: 10, speed: 1}
}

func (z *ZombieA) damage(damage float64) {
	z.hp -= damage
	if z.hp < 0 { // check it in game
		z.delete()
	}
}

func (z *ZombieA) delete() {
	z.is_deleted = true
}

func (z *ZombieA) getId() []byte {
	return z.id
}

func (z *ZombieA) getX() float64 {
	return z.x
}

func (z *ZombieA) getY() float64 {
	return z.y
}

func (z *ZombieA) isDeleted() bool {
	return z.is_deleted
}

func (z *ZombieA) getPonts() int {
	return z.points
}

func (z *ZombieA) isRight() bool {
	return z.is_right
}

func (z *ZombieA) setRight(isRight bool) {
	z.is_right = isRight
}

func (z *ZombieA) setX(x float64) {
	z.x = x
}

func (z *ZombieA) setY(y float64) {
	z.y = y
}

func (z *ZombieA) getKind() string {
	return z.kind
}

func (z *ZombieA) getSpeed() float64 {
	return z.speed
}
