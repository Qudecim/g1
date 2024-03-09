class WeaponA {
    created_time = null
    live_time = 10000

    type = 'a'
    player = null

    angle = 0
    isFinished = false

    x = 0
    y = 0



    static create(player, props) {
        let weapon = new WeaponA()
        weapon.player = player
        weapon.angle = props[0]
        weapon.x = player.x
        weapon.y = player.y
        weapon.created_time = Date.now()
        return weapon
    }

    draw(deltaTime) {
        let speed = 300

        let speedX = speed * Math.sin(this.angle)
        let speedY = speed * Math.cos(this.angle)
    
        this.x += speedX * deltaTime
        this.y += speedY * deltaTime

        
        
        // game.ctx.fillStyle = "black";
        // game.ctx.beginPath();
        // game.ctx.arc(this.x, this.y, 10, 0, 2 * Math.PI);
        // game.ctx.fill();

        game.ctx.translate(this.x, this.y);
        game.ctx.rotate(this.angle * -1);

        game.ctx.drawImage(res.getWeapon(this.type), 0, 0, 1024, 1024, -25, -25, 50, 50)

        game.ctx.rotate(-(this.angle*-1));
        game.ctx.translate(-(this.x), -(this.y));
        //game.ctx.restore();

        if (this.created_time + this.live_time < Date.now()) {
            this.isFinished = true
        }

    }
}