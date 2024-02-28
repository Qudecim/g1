class UI {
 
    showUpdates() {
        if (game.upgrades.length) {
            let element = document.getElementById("updates_container");
            element.classList.add("show");
            for (let i = 1; i < 4; i++) {
                let element1 = document.getElementById("update_" + i);
                element1.innerHTML = 'weapon: 1 update:' + game.upgrades[0][i]
            }
            
        }
    }
    
    sendUpdate1() {
        let item = game.upgrades.shift();
        sendUpdate(item[0]);
    }

    sendUpdate2() {
        let item = game.upgrades.shift();
        sendUpdate(item[1]); 
    }

    sendUpdate3() {
        let item = game.upgrades.shift();
        sendUpdate(item[2]);
    }

}