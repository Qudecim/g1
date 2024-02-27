class Control {

    init() {
        document.addEventListener("keydown", (event) => {
            if ( event.keyCode === 65) {
                if (!game.movement.left) {
                    game.movement.left = true
                    game.move()
                }
            }
            if (event.isComposing || event.keyCode === 68) {
                if (!game.movement.right) {
                    game.movement.right = true
                    game.move()
                }
            }
            if (event.isComposing || event.keyCode === 87) {
                if (!game.movement.up) {
                    game.movement.up = true
                    game.move()
                }
                
            }
            if (event.isComposing || event.keyCode === 83) {
                if (!game.movement.down) {
                    game.movement.down = true
                    game.move()
                }
            }
        });
    
        document.addEventListener("keyup", (event) => {
            if (event.isComposing || event.keyCode === 65) {
                game.movement.left = false
                game.move()
            }
            if (event.isComposing || event.keyCode === 68) {
                game.movement.right = false
                game.move()
            }
            if (event.isComposing || event.keyCode === 87) {
                game.movement.up = false
                game.move()
            }
            if (event.isComposing || event.keyCode === 83) {
                game.movement.down = false
                game.move()
            }
        });
    
    }

}