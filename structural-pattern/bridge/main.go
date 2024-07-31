package main

import "fmt"

func main() {
	vectorRenderer := &VectorRenderer{}
	circle := &Circle{0, 0, 5, vectorRenderer}
	square := &Square{5, vectorRenderer}

	circle.Draw()
	square.Draw()

	rasterRenderer := &RasterRenderer{}
	circle.renderer = rasterRenderer
	square.renderer = rasterRenderer

	circle.Draw()
	square.Draw()
}

type Renderer interface {
	RenderCircle(radius int)
	RenderSquare(side int)
}

type Shape interface {
	Draw()
}

type Circle struct {
	x, y, radius float32
	renderer     Renderer
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(int(c.radius))
}

type Square struct {
	side     int
	renderer Renderer
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.side)
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius int) {
	fmt.Println("Drawing a circle of radius", radius)
}

func (v *VectorRenderer) RenderSquare(side int) {
	fmt.Println("Drawing a square of side", side)
}

type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius int) {
	fmt.Println("Drawing pixels for a circle of radius", radius)
}

func (r *RasterRenderer) RenderSquare(side int) {
	fmt.Println("Drawing pixels for a square of side", side)
}
