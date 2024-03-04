class UI {
 
    showUpdates() {
        if (game.upgrades.length) {
            let element = document.getElementById("updates_container");
            element.classList.add("show");
            for (let i = 0; i < 3; i++) {
                let element1 = document.getElementById("update_" + i);
                element1.innerHTML = 'weapon: 1 update:' + game.upgrades[0][i]
            }
            
        }
    }
    
    sendUpdate1(item_id) {
        let item = game.upgrades.shift();
        console.log([item_id, item]);
        transport.sendUpdate(item[item_id]);
    }

}