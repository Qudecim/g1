class Game {

    canvas = null
    ctx = null

    id = null
    
    players = {}
    zombies = {}
    upgrades = []
    movement = {
        left: false,
        right: false,
        up: false,
        down: false
    }

    position = {
        x: 0,
        y: 0
    }

    resources_loaded = false

    constructor(canvas, ctx) {
        this.canvas = canvas
        this.ctx = ctx
    }

    move() {
        transport.send();
    }

    draw() {
        
        this.ctx.fillStyle = "white";
        this.ctx.fillRect(0,0, 1500, 800);

        if (!this.resources_loaded) {
            return
        }

        for (let zombie_id in this.zombies) {
            let zombie = this.zombies[zombie_id]
            zombie.draw()
        }

        for (let player_id in this.players) {
            let player = this.players[player_id]
            player.calc()
            player.draw()
            for(let weapon_id in player.weapons) {
                let weapon = player.weapons[weapon_id]
                weapon.draw()
                if (weapon.isFinished) {
                    player.weapons.splice(weapon_id, 1);
                }
            }
        }

        for (let zombie_id in this.zombies) {
            let zombie = this.zombies[zombie_id]
            if (zombie.idDie) {
                delete this.zombies[zombie_id]
            }
        }
    }
}