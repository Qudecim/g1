class ZombieB {

    x = 0
    y = 0
    directionIsRight = true

    isDie = false
    type = 'b'

    static create() {
        return new ZombieB()
    }

    draw() {
        game.ctx.drawImage(res.getZombie(this.type, this.directionIsRight), 0, 0, 50, 50, this.x - 25, this.y - 25, 50, 50)
    }

}