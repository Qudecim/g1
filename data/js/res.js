class Res {

    zombies = {
        a: {
            left: {
                src: '/data/src/z_1_l.png',
                img: null
            },
            right: {
                src: '/data/src/z_1_r.png',
                img: null
            }
        }
    }

    players = {
        a: {
            left: {
                src: '/data/src/p_1_l.png',
                img: null
            },
            right: {
                src: '/data/src/p_1_r.png',
                img: null
            }
        }
    }

    init() {
        let loaded = 0
        let need = 0
        for (let zombie_type in this.zombies) {
            for (let action_name in this.zombies[zombie_type]) {
                need++
                let img = new Image();
                img.src = this.zombies[zombie_type][action_name].src
                img.onload = function() {
                    loaded++
                    if (loaded == need) {
                        game.resources_loaded = true
                    }
                }
                this.zombies[zombie_type][action_name].img = img
            }
        }
        for (let player_type in this.players) {
            for (let action_name in this.players[player_type]) {
                need++
                let img = new Image();
                img.src = this.players[player_type][action_name].src
                img.onload = function() {
                    loaded++
                    if (loaded == need) {
                        game.resources_loaded = true
                    }
                }
                this.players[player_type][action_name].img = img
            }
        }
    }

    getZombie(type, directionIsRight) {
        let action = 'left'
        if (directionIsRight) {
            action = 'right'
        }
        return this.zombies[type][action].img
    }

    getPlayer(type, directionIsRight) {
        let action = 'left'
        if (directionIsRight) {
            action = 'right'
        }
        return this.players[type][action].img
    }

}