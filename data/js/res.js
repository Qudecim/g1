class Res {

    zombies = {
        a: {
            left: {
                src: '/data/z_1_l.png',
                img: null
            },
            right: {
                src: '/data/z_1_l.png',
                img: null
            }
        }
    }

    init() {
        let loaded = 0
        let need = 0
        for (zombie_name in this.zombies) {
            for (action_name in this.zombies[zombie_name]) {
                need++
                let img = new Image();
                img.src = this.zombies[zombie_name][action_name].src
                img.onload = function() {
                    loaded++
                    if (loaded == need) {
                        game.resources_loaded = true
                    }
                }
                this.zombies[zombie_name][action_name].img = img
            }
        }
    }

    getZombie(name, directionIsRight) {
        let action = 'left'
        if (directionIsRight) {
            action = 'right'
        }
        return this.zombies[name][action].img
    }

}