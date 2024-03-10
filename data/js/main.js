class Main {

    init() {
        let canvas = document.getElementById("canvas")
        let ctx = canvas.getContext("2d")

        data = new Data()

        res = new Res()
        res.init()

        game = new Game(canvas, ctx)

        transport = new Transport()
        transport.init()
        
        control = new Control()
        control.init()

        ui = new UI()

        requestAnimationFrame(function(currentTime) {
            lastFrameTime = currentTime;
            main_tic(currentTime);
        });
    }

}

function main_tic(currentTime) {
    var deltaTime = (currentTime - lastFrameTime) / 1000;

    game.draw(deltaTime)

    lastFrameTime = currentTime;
    requestAnimationFrame(main_tic);
    //setTimeout(main_tic, 1000 / 60);
}

res = null
game = null
transport = null
control = null
ui = null
data = null

lastFrameTime = 0

window.onload = function () {
    if (window["WebSocket"]) {
        let main = new Main()
        main.init()
    }
}