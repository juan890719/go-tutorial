package geo

import "fmt"

type geometry interface {
	area() float64
}

type Retangle struct {
	Width  float64
	Height float64
}

func (r *Retangle) area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) area() float64 {
	return c.Radius * c.Radius * 3.1416
}

func PrintInfo(g geometry) {
	fmt.Printf("%#v\n", g)
	fmt.Println("Its area:", g.area())
}
