package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	players     map[*Player]bool
	zombies     map[*Zombie]bool
	team_points int
	team_level  int
	hub         *Hub
}

func newGame(hub *Hub) *Game {
	return &Game{hub: hub, players: make(map[*Player]bool), zombies: make(map[*Zombie]bool)}
}

func (g *Game) addPlayer(player *Player) {
	g.players[player] = true
}

func (g *Game) deletePlayer(player *Player) {
	delete(g.players, player)
}

func (g *Game) deleteZombie(zombie *Zombie) {
	delete(g.zombies, zombie)
}

func (g *Game) run() {
	g.generateZombie()
	for {
		startTime := time.Now()
		var s []byte
		s = append(s, g.calc_players()...)
		s = append(s, g.calc_zombies()...)
		s = append(s, g.calc_weapons()...)
		g.calc_team()

		g.hub.broadcast <- s

		endTime := time.Now()
		difference := endTime.Sub(startTime)

		spf := (time.Second / 30) - difference
		time.Sleep(spf)
	}
}

func (g *Game) calc_players() []byte {

	var s []byte
	for player, _ := range g.players {
		if s != nil {
			d := []byte("&")
			s = append(s, d...)
		}
		if player.isDeleted {
			c := []byte("d:" + string(player.id))
			s = append(s, c...)
			g.deletePlayer(player)
			continue
		}
		speed := player.maxSpeed
		if (player.left || player.right) && (player.up || player.down) {
			speed = (math.Sqrt(math.Pow(speed, 2)+math.Pow(speed, 2)) / 2)
		}
		if player.left {
			player.x -= speed
			player.direction_is_right = false
		}
		if player.right {
			player.x += speed
			player.direction_is_right = true
		}
		if player.up {
			player.y -= speed
		}
		if player.down {
			player.y += speed
		}

		var dirIsRight string
		if player.direction_is_right {
			dirIsRight = "1"
		} else {
			dirIsRight = "0"
		}

		c := []byte("c:" + string(player.id) + ":" + dirIsRight + ":" + strconv.Itoa(int(player.x)) + ":" + strconv.Itoa(int(player.y)))

		s = append(s, c...)
	}

	return s
}

func (g *Game) calc_weapons() []byte {

	var s []byte
	for player, _ := range g.players {
		for _, weapon := range player.weapons {
			ws := weapon.calc(player, g)
			s = append(s, ws...)
		}
	}

	return s
}

func (g *Game) calc_zombies() []byte {

	var s []byte
	for zombie, _ := range g.zombies {
		if zombie.isDeleted {
			c := []byte("&r:" + string(zombie.id))
			s = append(s, c...)
			g.team_points += zombie.points
			g.deleteZombie(zombie)
			continue
		}
		var closest *Player
		var closestRange float64 = 1000
		for player, _ := range g.players {
			d := distance(player.x, player.y, zombie.x, zombie.y)
			if d < closestRange {
				closest = player
				closestRange = d
			}
		}
		dir_is_right := "1"
		if closest != nil {
			maxSpeed := float64(1)
			col_distance := float64(25)

			dx := closest.x - zombie.x
			dy := closest.y - zombie.y
			dist := distance(zombie.x, zombie.y, closest.x, closest.y)

			if dist <= 10 {
				continue
			}

			newX := zombie.x + (dx/dist)*math.Min(maxSpeed, dist)
			newY := zombie.y + (dy/dist)*math.Min(maxSpeed, dist)

			col := false
			for zombieCol, _ := range g.zombies {
				if zombie == zombieCol {
					continue
				}
				if distance(newX, newY, zombieCol.x, zombieCol.y) < col_distance {
					// angle := math.Atan2(dx, dy)
					// newX = zombie.x - math.Cos(angle)
					// newY = zombie.y - math.Sin(angle)
					col = true
					break
				}
			}

			if !col {
				zombie.x = newX
				zombie.y = newY
			}

			zombie.direction_is_right = closest.x > zombie.x
			if !zombie.direction_is_right {
				dir_is_right = "0"
			}
		}

		c := []byte("&z:a:" + string(zombie.id) + ":" + strconv.Itoa(int(zombie.x)) + ":" + strconv.Itoa(int(zombie.y)) + ":" + dir_is_right)
		s = append(s, c...)
	}
	return s
}

func (g *Game) calc_team() {
	if g.team_points >= g.team_level*10 {
		g.team_points = 0
		g.team_level++

		for player, _ := range g.players {
			ups := getUpgrades(player)
			c := []byte("u:" + strconv.Itoa(ups[0]) + ":" + strconv.Itoa(ups[1]) + ":" + strconv.Itoa(ups[2]))
			player.client.send <- c
		}
	}
}

func (g *Game) generateZombie() {
	for i := 0; i < 500; i++ {
		g.addZombie()
	}

	// r := rand.Intn(100)
	// if r < 10 {
	// 	g.addZombie()
	// }
}

func (g *Game) addZombie() {
	col_distance := float64(40)
	zombie := newZombie(generateId(), float64(rand.Intn(1500)), float64(rand.Intn(800)))
	col := false
	for zombieCol, _ := range g.zombies {
		if distance(zombie.x, zombie.y, zombieCol.x, zombieCol.y) < col_distance {
			col = true
		}
	}
	if !col {
		g.zombies[zombie] = true
	}
}
