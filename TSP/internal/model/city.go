package model

import (
	"math"
	"math/rand"
)

type City struct {
	X float64
	Y float64
}

func NewCity() City {
	return City{
		X: rand.Float64() * 100,
		Y: rand.Float64() * 100,
	}
}

func (c City) DistanceTo(other City) float64 {
	dx := c.X - other.X
	dy := c.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}
