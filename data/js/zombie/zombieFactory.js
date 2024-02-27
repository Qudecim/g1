class ZombieFactory {
    
    static make(name) {
        let zombie = null
        switch (name) {
            case 'a':
                zombie = ZombieA.build()
                break;
            case 4:
                alert( 'В точку!' );
                break;
            default:
                console.log('Wrong zombie: ' + name)
          }
    }

}