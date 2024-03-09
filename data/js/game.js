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

    _fps = 0
    step_fps = new Date().getTime()


    resources_loaded = false

    constructor(canvas, ctx) {
        this.canvas = canvas
        this.ctx = ctx
    }

    move() {
        transport.send();
    }

    draw(deltaTime) {

        let now = new Date().getTime()
        let fps = now - this.step_fps;
        if (fps < 1000) {
            this._fps++;
        } else {
            this.step_fps = now;
            let domFPS = document.getElementById('domFPS');
            domFPS.innerHTML = game._fps;
            this._fps = 0;
        }
        
        this.ctx.fillStyle = "#27ae60";
        this.ctx.fillRect(0,0, 1500, 800);

        if (!this.resources_loaded) {
            return
        }

        const entries = Object.entries(this.zombies);
        const entries2 = Object.entries(this.players);
        const array3 = entries.concat(entries2);
        array3.sort((a, b) => a[1].y - b[1].y);

        // for (let zombie_id in this.zombies) {
        //     let zombie = this.zombies[zombie_id]
        //     zombie.draw()
        // }
 
        for (let item of array3) {
            item[1].draw()
        }

        for (let player_id in this.players) {
            let player = this.players[player_id]
            player.calc()
            //player.draw()
            for(let weapon_id in player.weapons) {
                let weapon = player.weapons[weapon_id]
                weapon.draw(deltaTime)
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