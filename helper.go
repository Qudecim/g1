package main

import (
	"math"
	"math/rand"
)

func generateId() []byte {
	var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	n := 5
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func calculateAngle(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Atan2(dx, dy)
}
