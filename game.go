package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	zombies  []*Zombie
	zombieId int
}

type Zombie struct {
	id []byte
	x  int
	y  int
}

func newGame() *Game {
	return &Game{}
}

func newZombie(id []byte, x int, y int) *Zombie {
	return &Zombie{id: id, x: x, y: y}
}

func (g *Game) run(hub *Hub) {
	for {
		time.Sleep(time.Second / 30)

		var s []byte
		for client, _ := range hub.clients {
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
			c := []byte("c:" + string(client.id) + ":" + strconv.Itoa(client.y) + ":" + strconv.Itoa(client.x))
			s = append(s, c...)
		}

		s = append(s, g.zombie()...)

		//s := []byte("test")

		hub.broadcast <- s
	}
}

func (g *Game) zombie() []byte {
	g.generateZombie()
	var s []byte
	for _, zombie := range g.zombies {
		zombie.x++
		zombie.y++
		c := []byte("&c:" + string(zombie.id) + ":" + strconv.Itoa(zombie.y) + ":" + strconv.Itoa(zombie.x))
		s = append(s, c...)
	}
	return s
}

func (g *Game) generateZombie() {
	r := rand.Intn(100)
	if r < 10 {
		zombi := newZombie(generateId(), rand.Intn(500), rand.Intn(500))
		g.zombies = append(g.zombies, zombi)
	}
}
