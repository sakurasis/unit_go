package structs

import (
	"fmt"
	"math"
)

// Sphere a object
type Sphere struct {
	radius float64
}

// Volume is a function of calcuating the volume
func (s *Sphere) Volume() float64 {
	return float64(4) / float64(3) * math.Pi * math.Pow(s.radius, 3)
}

// SurfaceArea  a function of calcuating the surface-area
func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * math.Pow(s.radius, 2)
}

func main() {
	r := Sphere{
		radius: 2,
	}
	fmt.Println("radius:", r.radius)
	fmt.Println(r.SurfaceArea())
}
