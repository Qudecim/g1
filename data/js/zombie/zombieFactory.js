class ZombieFactory {
    
    static make(type) {
        let zombie = null
        switch (type) {
            case 'a':
                zombie = ZombieA.create()
                break;
            case 4:
                alert( 'В точку!' );
                break;
            default:
                console.log('Wrong zombie: ' + type)
          }
        return zombie
    }

}