class Main {

    init() {
        let canvas = document.getElementById("canvas")
        let ctx = canvas.getContext("2d")

        res = new Res()
        res.init()

        game = new Game(canvas, ctx)

        transport = new Transport()
        transport.init()
        
        control = new Control()
        control.init()

        ui = new UI()

        main_tic()
    }

}

function main_tic() {
    game.draw()
    setTimeout(main_tic, 1000 / 60);
}

res = null
game = null
transport = null
control = null
ui = null

window.onload = function () {
    if (window["WebSocket"]) {
        let main = new Main()
        main.init()
    }
}