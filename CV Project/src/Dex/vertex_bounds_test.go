package Dex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/vision/v1"
)

func TestBoundyBox(t *testing.T) {
	tests := []struct {
		poly vision.BoundingPoly
		minX int64
		maxX int64
		minY int64
		maxY int64
	}{
		{
			poly: vision.BoundingPoly{
				Vertices: []*vision.Vertex{
					{X: 0, Y: 0},
					{X: 0, Y: 10},
					{X: 10, Y: 10},
					{X: 10, Y: 00},
				},
			},
			minX: 0,
			minY: 0,
			maxX: 10,
			maxY: 10,
		},
		{
			poly: vision.BoundingPoly{
				Vertices: []*vision.Vertex{
					{X: 50, Y: 80},
					{X: 10, Y: 30},
					{X: 50, Y: 30},
					{X: 10, Y: 80},
				},
			},
			minX: 10,
			minY: 30,
			maxX: 50,
			maxY: 80,
		},
	}

	for _, test := range tests {
		box := NewBoundyBox(&test.poly)

		assert.Equal(t, test.minX, box.MinX)
		assert.Equal(t, test.minY, box.MinY)
		assert.Equal(t, test.maxX, box.MaxX)
		assert.Equal(t, test.maxY, box.MaxY)
	}

}
