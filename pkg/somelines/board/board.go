package board

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

// Board encapsulates simulation logic
type Board struct {
	width  int
	height int

	shaded bool
}

// New is a Board constructor
func New(w, h int) *Board {
	b := new(Board)

	b.width = w
	b.height = h

	b.shaded = false

	return b
}

// Setup prepares board
func (b *Board) Setup(numberOfCells int) {
}

// ToggleShading toggles shading
func (b *Board) ToggleShading() {
	b.shaded = !b.shaded
}

// Update performs board updates
func (b *Board) Update() error {
	return nil
}

// Size returns board size
func (b *Board) Size() (w, h int) {
	return b.width, b.height
}

// Draw draws board
func (b *Board) Draw(pixels []byte, counter, focalLength int) {
	theta := float64((2*counter)%360) * math.Pi / 180.0

	b.DrawPolygon(shapes.NewTriangle(point.NewPoint2D(30, 30), point.NewPoint2D(60, 25), point.NewPoint2D(70, 50)).Rotate(theta), pixels)

	b.DrawPolygon(shapes.NewQuadrangle(point.NewPoint2D(30, 70), point.NewPoint2D(60, 65), point.NewPoint2D(70, 90), point.NewPoint2D(40, 90)).Rotate(-theta), pixels)

	b.DrawLine(point.NewPoint2D(30, 120), point.NewPoint2D(70, 130), pixels)
	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 140, 0), point.NewPoint3D(70, 150, 20), point.NewPoint2D(30, 140), focalLength, pixels)
	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 160, 10), point.NewPoint3D(70, 175, 20), point.NewPoint2D(55, 160), focalLength, pixels)

	dy := 40
	p1 := point.NewPoint2D(80, 20)
	p2 := point.NewPoint2D(100, 40)
	p3 := point.NewPoint2D(120, 20)
	b.DrawPolygon(shapes.NewTriangle(p1, p2, p3), pixels)
	b.DrawPolygon3D(shapes.NewTriangle(p1.AddToY(dy), p2.AddToY(dy), p3.AddToY(dy)).RotateAroundX(theta), focalLength, pixels)
	b.DrawPolygon3D(shapes.NewTriangle(p1.AddToY(dy*2), p2.AddToY(dy*2), p3.AddToY(dy*2)).RotateAroundY(theta), focalLength, pixels)
	b.DrawPolygon3D(shapes.NewTriangle(p1.AddToY(dy*3), p2.AddToY(dy*3), p3.AddToY(dy*3)).RotateAroundZ(theta), focalLength, pixels)

	vertices3d := []*point.Point3D{
		point.NewPoint3D(140, 20, 0),
		point.NewPoint3D(170, 20, 0),
		point.NewPoint3D(140, 50, 0),
		point.NewPoint3D(170, 50, 0),
		point.NewPoint3D(140, 20, 30),
		point.NewPoint3D(170, 20, 30),
		point.NewPoint3D(140, 50, 30),
		point.NewPoint3D(170, 50, 30),
	}
	edges3d := [][2]int{
		{0, 1}, {0, 2}, {0, 4},
		{3, 1}, {3, 2}, {3, 7},
		{5, 1}, {5, 4}, {5, 7},
		{6, 2}, {6, 4}, {6, 7},
	}

	// dy := 40
	vertices3d2 := []*point.Point3D{
		vertices3d[0].AddToY(dy),
		vertices3d[1].AddToY(dy),
		vertices3d[2].AddToY(dy),
		vertices3d[3].AddToY(dy),
		vertices3d[4].AddToY(dy),
		vertices3d[5].AddToY(dy),
		vertices3d[6].AddToY(dy),
		vertices3d[7].AddToY(dy),
	}
	vertices3d3 := []*point.Point3D{
		vertices3d2[0].AddToY(dy),
		vertices3d2[1].AddToY(dy),
		vertices3d2[2].AddToY(dy),
		vertices3d2[3].AddToY(dy),
		vertices3d2[4].AddToY(dy),
		vertices3d2[5].AddToY(dy),
		vertices3d2[6].AddToY(dy),
		vertices3d2[7].AddToY(dy),
	}
	vertices3d4 := []*point.Point3D{
		vertices3d3[0].AddToY(dy),
		vertices3d3[1].AddToY(dy),
		vertices3d3[2].AddToY(dy),
		vertices3d3[3].AddToY(dy),
		vertices3d3[4].AddToY(dy),
		vertices3d3[5].AddToY(dy),
		vertices3d3[6].AddToY(dy),
		vertices3d3[7].AddToY(dy),
	}

	b.DrawPolygon3D(shapes.NewPolygon3D(vertices3d, edges3d), focalLength, pixels)
	b.DrawPolygon3D(shapes.NewPolygon3D(vertices3d2, edges3d).RotateAroundX(-theta), focalLength, pixels)
	b.DrawPolygon3D(shapes.NewPolygon3D(vertices3d3, edges3d).RotateAroundY(-theta), focalLength, pixels)
	b.DrawPolygon3D(shapes.NewPolygon3D(vertices3d4, edges3d).RotateAroundZ(-theta), focalLength, pixels)
}

// Draw2 draws board
func (b *Board) Draw2(pixels []byte, counter, focalLength int) {
	theta := float64((int(1.5*float64(counter)))%360) * math.Pi / 180.0
	_ = theta

	vertices3d := []*point.Point3D{
		point.NewPoint3D(50, 40, 0),
		point.NewPoint3D(130, 40, 0),
		point.NewPoint3D(60, 100, 0),
		point.NewPoint3D(120, 100, 0),
		point.NewPoint3D(70, 40, 60),
		point.NewPoint3D(150, 40, 60),
		point.NewPoint3D(60, 100, 60),
		point.NewPoint3D(120, 100, 60),
	}
	edges3d := [][2]int{
		{0, 1}, {0, 2}, {0, 4},
		{3, 1}, {3, 2}, {3, 7},
		{5, 1}, {5, 4}, {5, 7},
		{6, 2}, {6, 4}, {6, 7},
	}

	b.DrawPolygon3D(shapes.NewPolygon3D(vertices3d, edges3d).RotateAroundX(-theta).RotateAroundY(-theta), focalLength, pixels)
}

// DrawPolygon draws a polygon based on vertices and edges matrices
func (b *Board) DrawPolygon(polygon *shapes.Polygon2D, pixels []byte) {
	for _, edge := range polygon.Edges {
		b.DrawLine(polygon.Vertices[edge[0]], polygon.Vertices[edge[1]], pixels)
	}
}

// DrawPolygon3D draws a 3D polygon based on vertices and edges matrices
func (b *Board) DrawPolygon3D(polygon *shapes.Polygon3D, focalLength int, pixels []byte) {
	var sumX, sumY int
	lv := len(polygon.Vertices)
	for _, v := range polygon.Vertices {
		sumX += v.X
		sumY += v.Y
	}

	pc := point.NewPoint2D(sumX/lv, sumY/lv)
	for _, edge := range polygon.Edges {
		b.DrawLine3D(polygon.Vertices[edge[0]], polygon.Vertices[edge[1]], pc, focalLength, pixels)
	}
}

// DrawLine3DRelativeToPoint draws a line
func (b *Board) DrawLine3DRelativeToPoint(p1, p2 *point.Point3D, pc *point.Point2D, focalLength int, pixels []byte) {
	b.DrawLine3D(p1, p2, pc, focalLength, pixels)
}

// DrawLine draws a line
func (b *Board) DrawLine(p1, p2 *point.Point2D, pixels []byte) {
	var (
		dx, sx int
		dy, sy int
	)

	if p1.X < p2.X {
		sx = 1
		dx = p2.X - p1.X
	} else {
		sx = -1
		dx = p1.X - p2.X
	}

	if p1.Y < p2.Y {
		sy = 1
		dy = p2.Y - p1.Y
	} else {
		sy = -1
		dy = p1.Y - p2.Y
	}

	p0 := point.NewPoint2D(p1.X, p1.Y)

	switch {
	case dx > dy:
		ai := (dy - dx) * 2
		bi := dy * 2
		d := bi - dx
		b.colorPixel(p0.X, p0.Y, pixels)
		for p0.X != p2.X {
			if d >= 0 {
				p0.Y += sy
				d += ai
			} else {
				d += bi
			}
			p0.X += sx
			b.colorPixel(p0.X, p0.Y, pixels)
		}
	case dx <= dy:
		ai := (dx - dy) * 2
		bi := dx * 2
		d := bi - dy
		b.colorPixel(p0.X, p0.Y, pixels)
		for p0.Y != p2.Y {
			if d >= 0 {
				p0.X += sx
				d += ai
			} else {
				d += bi
			}
			p0.Y += sy
			b.colorPixel(p0.X, p0.Y, pixels)
		}
	}
}

// DrawLine3D draws a 3D line
func (b *Board) DrawLine3D(p1, p2 *point.Point3D, pc *point.Point2D, focalLength int, pixels []byte) {
	var (
		dx, sx int
		dy, sy int
		dz     int
		z, sz  float64
		ps, pe *point.Point2D
	)

	if p1.Z > p2.Z {
		ps = p2.To2DRelativeToPoint(pc, focalLength)
		pe = p1.To2DRelativeToPoint(pc, focalLength)
		dz = p1.Z - p2.Z
	} else {
		ps = p1.To2DRelativeToPoint(pc, focalLength)
		pe = p2.To2DRelativeToPoint(pc, focalLength)
		dz = p2.Z - p1.Z
	}

	if ps.X < pe.X {
		sx = 1
		dx = pe.X - ps.X
	} else {
		sx = -1
		dx = ps.X - pe.X
	}

	if ps.Y < pe.Y {
		sy = 1
		dy = pe.Y - ps.Y
	} else {
		sy = -1
		dy = ps.Y - pe.Y
	}

	switch {
	case dx > dy:
		ai := (dy - dx) * 2
		bi := dy * 2
		d := bi - dx
		z = float64(p1.Z)
		if p1.Z == p2.Z {
			sz = 0
		} else if dy == 0 {
			sz = float64(dz)
		} else {
			sz = float64(dz) / float64(dy)
		}
		b.color3DPixel(ps.X, ps.Y, z, pixels)
		for ps.X != pe.X {
			if d >= 0 {
				ps.Y += sy
				d += ai
			} else {
				d += bi
			}
			ps.X += sx
			if (sz > 0 && z <= float64(p2.Z)) || (sz < 0 && z >= float64(p2.Z)) {
				z += sz
			}
			b.color3DPixel(ps.X, ps.Y, z, pixels)
		}
	case dx <= dy:
		ai := (dx - dy) * 2
		bi := dx * 2
		d := bi - dy
		z = float64(p1.Z)
		if p1.Z == p2.Z {
			sz = 0
		} else if dx == 0 {
			sz = float64(dz)
		} else {
			sz = float64(dz) / float64(dx)
		}
		b.color3DPixel(ps.X, ps.Y, z, pixels)
		for ps.Y != pe.Y {
			if d >= 0 {
				ps.X += sx
				d += ai
			} else {
				d += bi
			}
			ps.Y += sy
			if (sz > 0 && z <= float64(p2.Z)) || (sz < 0 && z >= float64(p2.Z)) {
				z += sz
			}
			b.color3DPixel(ps.X, ps.Y, z, pixels)
		}
	}
}

func (b *Board) colorPixel(x, y int, pixels []byte) {
	i := b.getPixelsIndex(x, y)
	pixels[4*i] = 0xf0
	pixels[4*i+1] = 0xf0
	pixels[4*i+2] = 0xf0
	pixels[4*i+3] = 0xff
}

func (b *Board) color3DPixel(x, y int, z float64, pixels []byte) {
	colorAt0 := byte(0xe0)
	color := colorAt0

	if b.shaded {
		modifier := 2 * z
		switch {
		case modifier > 200:
			modifier = 200
		case modifier < -30:
			modifier = -30
		}

		color = colorAt0 - byte(modifier)
	}

	i := b.getPixelsIndex(x, y)
	if pixels[4*i] < color {
		pixels[4*i] = color
		pixels[4*i+1] = color
		pixels[4*i+2] = color
		pixels[4*i+3] = 0xff
	}
}

func (b *Board) getPixelsIndex(x, y int) int {
	for x < 0 || x >= b.width {
		switch {
		case x < 0:
			x += b.width
		case x >= b.width:
			x -= b.width
		}
	}

	for y < 0 || y >= b.height {
		switch {
		case y < 0:
			y += b.height
		case y >= b.height:
			y -= b.height
		}
	}

	max := b.width * b.height
	i := y*b.width + x
	for i < 0 || i >= max {
		switch {
		case i < 0:
			i += max
		case i >= max:
			i -= max
		}
	}

	return i
}
