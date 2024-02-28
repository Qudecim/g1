class WeaponFactory {
    
    static make(type, player, props) {
        let weapon = null
        switch (type) {
            case 'a':
                weapon = WeaponA.create(player, props)
                break;
            case 4:
                alert( 'В точку!' );
                break;
            default:
                console.log('Wrong weapon: ' + type)
          }

          return weapon
    }

}