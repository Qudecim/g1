class Transport {

    init() {
        this.connection = new WebSocket("ws://" + document.location.host + "/ws");

        this.connection.onclose = function (evt) {
            // TODO: reconect
        };

        let current_instance = this

        this.connection.onmessage = function (evt) {
            let msg = evt.data;
            if (game.id == null) {
                let properties = evt.data.split(':');
                if (properties[0] == 'id') {
                    game.id = properties[1];
                }
                return;
            }
            
            let objects = msg.split('&');
            for (let index in objects) {
                
                let object = objects[index]
                if (object == '') {
                    continue;
                }

                let properties = object.split(':');

                switch (properties[0]) {
                    case 'c':
                        current_instance.getPlayer(properties[1], properties[2], properties[3])
                        break;
                    case 'd':
                        current_instance.diePlayer(properties[1])
                        break;
                    case 'z':
                        current_instance.getZombie(properties[1], properties[2], properties[3], properties[4], properties[5])
                        break;
                    case 'r':
                        current_instance.dieZombie(properties[1])
                        break;
                    case 'w':
                        current_instance.getWeapon(properties[1], properties[2], [properties[3]]) // TODO: make slice for props
                        break;
                    case 'u':
                        current_instance.getUpgrades(properties[1], properties[2], properties[3])
                        break;
                    default:
                        console.log('Wrong action: ' + properties[0])
                }
            }
        }
    }

    getPlayer(id,  x, y) {
        if (game.players[id] == undefined) {
            game.players[id] = Player.create()
        }
        game.players[id].to_x = x
        game.players[id].to_y = y
    }
    
    diePlayer(id) {
        delete players[id]
    }

    getZombie(type, id, x, y, directionIsRight) {
        if (game.zombies[id] == undefined) {
            game.zombies[id] = ZombieFactory.make(type)
        }
        game.zombies[id].x = x
        game.zombies[id].y = y
        game.zombies[id].directionIsRight = directionIsRight == '1'
    }

    dieZombie(id) {
        if (game.zombies[id] != undefined) {
            game.zombies[id].idDie = true
        }
    }

    getWeapon(type, player_id, properties) {
        if (game.players[player_id] != undefined) {
            let player = game.players[player_id] 
            let weapon = WeaponFactory.make(type, player, properties)
            player.weapons.push(weapon)
        }
    }

    getUpgrades(u1, u2, u3) {
        game.upgrades.push(u1, u2, u3)
        ui.showUpdates()
    }

    send() {
        let binaryString = '';
        let bytes = Array(5);
        bytes[0] = 0x0; // send move
        if (game.movement.left) {
            bytes[1] = 0x1;
        } else {
            bytes[1] = 0x0;
        }
        if (game.movement.right) {
            bytes[2] = 0x1;
        } else {
            bytes[2] = 0x0;
        }
        if (game.movement.up) {
            bytes[3] = 0x1;
        } else {
            bytes[3] = 0x0;
        }
        if (game.movement.down) {
            bytes[4] = 0x1;
        } else {
            bytes[4] = 0x0;
        }
        var length = bytes.length;
        for (var i = 0; i < length; i++) {
            binaryString += String.fromCharCode(bytes[i]);
        }
        this.connection.send(binaryString);
    }

    sendUpdate(update) {
        binaryString = '';
        var bytes = Array(2);
        bytes[0] = 0x1; // send update
        bytes[1] = update;
        var length = bytes.length;
        for (var i = 0; i < length; i++) {
            binaryString += String.fromCharCode(bytes[i]);
        }
        this.connection.send(binaryString);
        let element = document.getElementById("updates_container");
        element.classList.remove("show");
        
        if (game.updates.length) {
            ui.showUpdates()
        }
    }

    
}