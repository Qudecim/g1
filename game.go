package main

import (
	"strconv"
	"time"
)

type Game struct {
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) run(hub *Hub) {
	for {
		time.Sleep(time.Second / 30)

		var s []byte
		for client, _ := range hub.clients {
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
			s = []byte("c:" + string(client.id) + ":" + strconv.Itoa(client.y) + ":" + strconv.Itoa(client.x))
		}

		//s := []byte("test")

		hub.broadcast <- s
	}
}
