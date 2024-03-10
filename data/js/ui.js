class UI {

    show_upgrades = false
 
    showUpdates() {
        if (game.upgrades.length && !this.show_upgrades) {
            let element = document.getElementById("updates_container");
            element.classList.add("show");
            element.innerHTML = '';

            for (let i = 0; i < game.upgrades[0].length; i++) {
                let upgrade_value = game.upgrades[0][i];
                let upgrade_data = data.getUpgrade(upgrade_value)

                let html = '<div class="update_item" onclick="ui.sendUpdate(' + i + ')"><div class="upgrade_img_container"><img class="upgrade_img" src="' + upgrade_data['img'] + '"></div><span class="upgrade_text">' + upgrade_data['text'] + '</span></div>';

                element.innerHTML = element.innerHTML + html;
            }

            this.show_upgrades = true
        }
    }
    
    sendUpdate(item_id) {
        let item = game.upgrades.shift();
        console.log([item_id, item]);
        this.show_upgrades = false
        transport.sendUpdate(item[item_id]);
    }

}