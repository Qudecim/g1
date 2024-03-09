package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

type WeaponA struct {
	last_run time.Time

	damage         float64
	timeout        int64
	criticalChance float64
	criticalDamage float64
	count          int

	live_time int64

	tiles map[*WeaponATile]time.Time
}

func newWeaponContainer() *WeaponA {

	damage := float64(100)
	timeout := int64(time.Second)
	criticalChance := float64(10)
	criticalDamage := float64(10)
	count := 3
	live_time := int64(10000)

	return &WeaponA{
		last_run:       time.Now(),
		tiles:          make(map[*WeaponATile]time.Time),
		damage:         damage,
		timeout:        timeout,
		criticalChance: criticalChance,
		criticalDamage: criticalDamage,
		count:          count,
		live_time:      live_time,
	}
}

func (w *WeaponA) calc(player *Player, game *Game) []byte {
	next_run := w.last_run.Add(time.Duration(w.timeout))
	var s []byte

	currentTime := time.Now()

	for tile, time_created := range w.tiles {
		current_live_time := time_created.Add(time.Duration(w.live_time))
		if current_live_time.Before(currentTime) {
			//delete(w.tiles, tile)
			//continue
		}
		tile.calc(player, game)
	}

	if next_run.Before(currentTime) {
		s = append(s, w.addTile(player, game)...)
		w.last_run = currentTime
	}

	return s
}

func (w *WeaponA) addTile(player *Player, game *Game) []byte {

	// var closest ZombieInterface
	// var closestRange float64 = 2000
	// for zombie, _ := range game.zombies {
	// 	d := distance(zombie.getX(), zombie.getY(), player.x, player.y)
	// 	if d < closestRange {
	// 		closest = zombie
	// 		closestRange = d
	// 	}
	// }

	// angle := calculateAngle(player.x, player.y, closest.getX(), closest.getY())

	var s []byte

	for i := 0; i < w.count; i++ {
		//rand.Seed(time.Now().UnixNano())
		randomRadian := rand.Float64() * 2 * math.Pi
		angle := math.Mod(randomRadian, 2*math.Pi)

		tile := newWeaponATile(player.x, player.y, w.damage, w.criticalChance, w.criticalDamage, angle)

		w.tiles[tile] = time.Now()

		c := []byte("&w:a:" + string(player.id) + ":" + strconv.FormatFloat(angle, 'f', -1, 32))
		s = append(s, c...)
	}

	return s
}

func (w *WeaponA) getUpgrades() []int {
	return []int{1, 2, 3, 4}
}

func (w *WeaponA) upgrade(upgrade int) {
	switch upgrade {
	case 1:
		w.damage = w.damage * 2
	case 2:
		w.timeout = int64(float64(w.timeout) * 0.2)
	case 3:
		w.criticalChance = w.criticalChance * 1.1
	case 4:
		w.criticalDamage = w.criticalDamage * 1.1
	}
}
