package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	zombies  []*Zombie
	zombieId int
	hub      *Hub
}

type Zombie struct {
	id []byte
	x  float64
	y  float64
}

func newGame(hub *Hub) *Game {
	return &Game{hub: hub}
}

func newZombie(id []byte, x float64, y float64) *Zombie {
	return &Zombie{id: id, x: x, y: y}
}

func (g *Game) run() {
	g.generateZombie()
	for {
		time.Sleep(time.Second / 30)

		var s []byte
		for client, _ := range g.hub.clients {
			if s != nil {
				d := []byte("&")
				s = append(s, d...)
			}
			if client.left {
				client.x--
			}
			if client.right {
				client.x++
			}
			if client.up {
				client.y--
			}
			if client.down {
				client.y++
			}
			c := []byte("c:" + string(client.id) + ":" + strconv.Itoa(int(client.y)) + ":" + strconv.Itoa(int(client.x)))
			s = append(s, c...)
		}

		s = append(s, g.zombie()...)

		g.hub.broadcast <- s
	}
}

func (g *Game) zombie() []byte {

	var s []byte
	for _, zombie := range g.zombies {
		var closest *Client
		var closestRange float64 = 1000
		for client, _ := range g.hub.clients {
			d := distance(client.x, client.y, zombie.x, zombie.y)
			if d < closestRange {
				closest = client
				closestRange = d
			}
		}
		if closest != nil {
			maxSpeed := float64(2)
			col_distance := float64(10)

			dx := closest.x - zombie.x
			dy := closest.y - zombie.y
			dist := distance(zombie.x, zombie.y, closest.x, closest.y)

			if dist <= 10 {
				continue
			}

			newX := zombie.x + (dx/dist)*math.Min(maxSpeed, dist)
			newY := zombie.y + (dy/dist)*math.Min(maxSpeed, dist)

			col := false
			for _, zombieCol := range g.zombies {
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

		}

		c := []byte("&z:" + string(zombie.id) + ":" + strconv.Itoa(int(zombie.y)) + ":" + strconv.Itoa(int(zombie.x)))
		s = append(s, c...)
	}
	return s
}

func (g *Game) generateZombie() {
	col_distance := float64(10)
	// r := rand.Intn(100)
	// if r < 10 {
	for i := 0; i < 200; i++ {
		zombie := newZombie(generateId(), float64(rand.Intn(500)), float64(rand.Intn(500)))

		col := false
		for _, zombieCol := range g.zombies {
			if distance(zombie.x, zombie.y, zombieCol.x, zombieCol.y) < col_distance {
				col = true
			}
		}
		if !col {
			g.zombies = append(g.zombies, zombie)
		}
	}

	//}
}

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
