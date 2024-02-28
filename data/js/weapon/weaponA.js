class WeaponA {

    type = 'a'
    player = null

    to = ''
    isFinished = true // Нужно ли удалить после отрисовки

    static create(player, props) {
        let weapon = new WeaponA()
        weapon.player = player
        weapon.to = props[0]
        return weapon
    }

    draw() {
        if (game.zombies[this.to] != undefined) {
            let zombie = game.zombies[this.to]
            game.ctx.beginPath();
            game.ctx.moveTo(this.player.x, this.player.y); 
            game.ctx.lineTo(zombie.x, zombie.y);
            game.ctx.stroke()
        }
    }
}