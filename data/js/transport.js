class Transport {

    init() {
        this.connection = new WebSocket("ws://" + document.location.host + "/ws");

        conn.onclose = function (evt) {
            //appendLog('Connection closed');
        };

        conn.onmessage = function (evt) {
            var msg = evt.data;
            if (id == '') {
                properties = evt.data.split(':');
                if (properties[0] == 'id') {
                    id = properties[1];
                }
                return;
            }
            
            objects = msg.split('&');
            for (index in objects) {
                
                let object = objects[index]
                if (object == '') {
                    continue;
                }

                properties = object.split(':');
                if (properties[0] == 'c') {
                    player_id = properties[1]
                    player_y = properties[2]
                    player_x = properties[3]
                    if (players[player_id] == undefined) {
                        players[player_id] = {
                        x: player_x,
                        y: player_y
                    }
                    }
                    players[player_id]['to_x'] = player_x
                    players[player_id]['to_y'] = player_y
                }
                if (properties[0] == 'z') {
                    zobmie_id = properties[1]
                    zobmie_y = properties[2]
                    zobmie_x = properties[3]
                    zobmie_dir = true
                    if (properties[4] == '0') {
                        zobmie_dir = false
                    }
                    zombies[zobmie_id] = {
                        x: zobmie_x,
                        y: zobmie_y,
                        isDeleted: false,
                        direction_is_right: zobmie_dir
                    }
                }
                if (properties[0] == 'd') {
                    player_id = properties[1]
                    delete players[player_id]
                }

                if (properties[0] == 'r') {
                    zombie_id = properties[1]
                    if (zombies[zombie_id] != undefined) {
                        zombies[zombie_id].isDeleted = true
                    }
                    
                }

                if (properties[0] == 'w') {
                    if (properties[1] == 1) {
                        weapons.push({
                            'type' : 1,
                            'player' : properties[2],
                            'to': properties[3], 
                            'frame': 0
                        })
                    }
                }

                if (properties[0] == 'u') {
                    if (properties[1] == id) {
                        updates.push([properties[2], properties[3], properties[4], properties[5]])
                        showUpdates()
                    }
                }

            
            }
        };
    }

    send() {
        binaryString = '';
        var bytes = Array(5);
        bytes[0] = 0x0; // send move
        if (left) {
            bytes[1] = 0x1;
        } else {
            bytes[1] = 0x0;
        }
        if (right) {
            bytes[2] = 0x1;
        } else {
            bytes[2] = 0x0;
        }
        if (up) {
            bytes[3] = 0x1;
        } else {
            bytes[3] = 0x0;
        }
        if (down) {
            bytes[4] = 0x1;
        } else {
            bytes[4] = 0x0;
        }
        var length = bytes.length;
        for (var i = 0; i < length; i++) {
            binaryString += String.fromCharCode(bytes[i]);
        }
        conn.send(binaryString);
    }

    
    sendUpdate(weapon, update) {
        console.log([weapon, update])
        binaryString = '';
        var bytes = Array(3);
        bytes[0] = 0x1; // send update
        bytes[1] = weapon;
        bytes[2] = update;
        var length = bytes.length;
        for (var i = 0; i < length; i++) {
            binaryString += String.fromCharCode(bytes[i]);
        }
        conn.send(binaryString);
        let element = document.getElementById("updates_container");
        element.classList.remove("show");
        
        if (updates.length) {
            showUpdates()
        }
    }
}