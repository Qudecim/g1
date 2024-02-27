class UI {
 
    showUpdates() {
        if (updates.length) {
            let element = document.getElementById("updates_container");
            element.classList.add("show");
            for (let i = 1; i < 4; i++) {
                let element1 = document.getElementById("update_" + i);
                element1.innerHTML = 'weapon: 1 update:' + updates[0][i]
            }
            
        }
    }
    
    sendUpdate1() {
        let item = updates.shift();
        sendUpdate(item[0], item[1]);
    }

    sendUpdate2() {
        let item = updates.shift();
        sendUpdate(item[0], item[2]); 
    }

    sendUpdate3() {
        let item = updates.shift();
        sendUpdate(item[0], item[3]);
    }

}