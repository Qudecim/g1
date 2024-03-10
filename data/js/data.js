class Data {

    upgrades = {
        1: {
            text: 'Скорость движения +10%',
            img: 'data/src/up_hero.png'
        },
        2: {
            text: 'Максимальное HP +10',
            img: 'data/src/up_hero.png'
        },
        3: {
            text: 'Уклонение +1%',
            img: 'data/src/up_evasion.png'
        },
        4: {
            text: 'Броня +1',
            img: 'data/src/up_armour.png'
        },

        11: {
            text: 'Урон +2',
            img: 'data/src/up1.png'
        },
        12: {
            text: 'Задержка перед использованием -1%',
            img: 'data/src/up1.png'
        },
        13: {
            text: 'Шанс крита +1%',
            img: 'data/src/up1.png'
        },
        14: {
            text: 'Урон крита +1%',
            img: 'data/src/up1.png'
        },
        15: {
            text: 'Количество тайлов +1',
            img: 'data/src/up1.png'
        },
    }


    getUpgrade(upgrade) {
        return this.upgrades[upgrade]
    }
}