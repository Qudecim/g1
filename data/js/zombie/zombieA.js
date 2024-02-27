class ZombieA {

    x = 0
    y = 0
    directionIsRight = true

    name = 'a'

    static build() {
        return new ZombieA()
    }

    draw(res) {
        ctx.drawImage(res.getZombie(this.name, this.directionIsRight), 0, 0, 50, 50, this.x - 25, this.y - 25, 50, 50)
    }

}