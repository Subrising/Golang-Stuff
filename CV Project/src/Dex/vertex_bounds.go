package Dex

import (
	"image"
	"google.golang.org/api/vision/v1"
)

type CVBoundingPoly struct {
	Vertices []image.Point `json:"vertices"`
}

type BoundyBox struct {
	init bool
	MinX int64
	MinY int64
	MaxX int64
	MaxY int64
}

// NewBoundyBox returns a BoundyBox for the vertices of boundingPoly
// with x, y coordinates
func NewBoundyBox(poly *vision.BoundingPoly) BoundyBox {
	box := BoundyBox{}

	for _, vertex := range poly.Vertices {
		box.AddPoint(vertex)
	}

	return box
}

// AddPoint adds a point to BoundyBox and set minX, minY, ... you get it
func (b *BoundyBox) AddPoint(p *vision.Vertex) {
	if b.init == false {
		b.MinX = p.X
		b.MaxX = p.X
		b.MinY = p.Y
		b.MaxY = p.Y
		b.init = true
	}

	if b.MinY > p.Y {
		b.MinY = p.Y
	}
	if b.MinX > p.X {
		b.MinX = p.X
	}
	if b.MaxY < p.Y {
		b.MaxY = p.Y
	}
	if b.MaxX < p.X {
		b.MaxX = p.X
	}
}
