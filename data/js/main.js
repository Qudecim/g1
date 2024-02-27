class Main {

    init() {
        let canvas = document.getElementById("canvas")
        let ctx = canvas.getContext("2d")

        res = new Res()
        res.init()

        game = new Game(canvas, ctx)
        tansport = new Transport()
        transport.init()
        control = new Control()
        control.init()

    }

} 

res = null
game = null
transport = null
control = null

window.onload = function () {
    if (window["WebSocket"]) {
        let main = new Main()
        main.init()
    }
}