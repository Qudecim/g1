class Player {

    x = 0
    y = 0
    to_x = 0
    to_y = 0
    directionIsRight = true
    isDied = false

    weapons = []

    static create() {
        return new Player()
    }

    calc() {
        let x = (this.to_x - this.x) / 4
        this.x = (this.x*1) + (x*1)
        let y = (this.to_y - this.y) / 4
        this.y = (this.y*1) + (y*1)
    }

    draw() {
        game.ctx.drawImage(res.getPlayer('a', this.directionIsRight), 0, 0, 100, 100, this.x - 25, this.y - 25, 50, 50)
    }

}