class ZombieFactory {
    
    static make(type) {
        let zombie = null
        switch (type) {
            case 'a':
                zombie = ZombieA.create()
                break;
            case 'b':
                zombie = ZombieB.create()
                break;
            default:
                console.log('Wrong zombie: "' + type + '"')
          }
        return zombie
    }

}