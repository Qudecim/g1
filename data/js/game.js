class Game {

    canvas = null
    ctx = null
    
    players = []
    weapons = []
    upgrades = []

    movment = {
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

    move() {
        transport.send();
    }

    draw() {
        
        ctx.fillStyle = "white";
        ctx.fillRect(0,0, 1500, 800);

        if (!loadImages) {
            setTimeout(draw, 1000 / 60);
            return
        }

        for (player_id in players) {
            player = players[player_id]
            if (player_id == id) {
                ctx.fillStyle = "blue";
            } else {
                ctx.fillStyle = "green";
            }

            let x = (player.to_x - player.x) / 4
            player.x = (player.x*1) + (x*1)
            let y = (player.to_y - player.y) / 4
            player.y = (player.y*1) + (y*1)

            ctx.beginPath();
            ctx.arc(player.x, player.y, 5, 0, 2 * Math.PI);
            ctx.fill();
        }

        for (zombie_id in zombies) {
            zombie = zombies[zombie_id]
            

            if (zombies[zombie_id].direction_is_right) {
                ctx.drawImage(tomato_r, 0, 0, 50, 50, zombie.x - 25, zombie.y - 25, 50, 50)
            } else {
                ctx.drawImage(tomato_l, 0, 0, 50, 50, zombie.x - 25, zombie.y - 25, 50, 50)
            }
            
        }

        for (weapon_id in weapons) {
            weapon = weapons[weapon_id]
            player = players[weapon.player]
            zombie = zombies[weapon.to]
            if (zombies[weapon.to] != undefined) {
                ctx.beginPath();
                ctx.moveTo(player.x, player.y); 
                ctx.lineTo(zombie.x, zombie.y);
                ctx.stroke()
            }

            weapons[weapon_id].frame++
            if (weapons[weapon_id].frame > 1) {
                weapons.splice(weapon_id, 1);
            }
        }
        
        for (zombie_id in zombies) {
            zombie = zombies[zombie_id]
            if (zombie.isDeleted) {
                delete zombies[zombie_id]
            }
        }

        setTimeout(draw, 1000 / 60);
    }
}