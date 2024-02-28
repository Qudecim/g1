class ZombieA {

    x = 0
    y = 0
    directionIsRight = true

    isDie = false
    type = 'a'

    static create() {
        return new ZombieA()
    }

    draw() {
        game.ctx.drawImage(res.getZombie(this.type, this.directionIsRight), 0, 0, 50, 50, this.x - 25, this.y - 25, 50, 50)
    }

}