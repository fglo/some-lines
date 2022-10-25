package board

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/point"
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
	// b.RotateTriangle(point.NewPoint2D(30, 30), point.NewPoint2D(60, 25), point.NewPoint2D(40, 50), theta, pixels)
	b.RotateQuadrangle(point.NewPoint2D(30, 30), point.NewPoint2D(60, 25), point.NewPoint2D(70, 50), point.NewPoint2D(40, 50), theta, pixels)

	vertices := []*point.Point2D{point.NewPoint2D(30, 70), point.NewPoint2D(60, 65), point.NewPoint2D(70, 90), point.NewPoint2D(40, 90)}
	edges := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}
	b.RotatePolygon(vertices, edges, -theta, pixels)

	b.DrawLine(point.NewPoint2D(30, 120), point.NewPoint2D(70, 130), pixels)
	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 140, 0), point.NewPoint3D(70, 150, 20), point.NewPoint2D(30, 140), focalLength, pixels)
	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 160, 10), point.NewPoint3D(70, 175, 20), point.NewPoint2D(55, 160), focalLength, pixels)

	p1 := point.NewPoint3D(80, 20, 0)
	p2 := point.NewPoint3D(100, 40, 0)
	p3 := point.NewPoint3D(120, 20, 0)
	b.DrawTriangle(point.NewPoint2D(80, 20), point.NewPoint2D(100, 40), point.NewPoint2D(120, 20), pixels)
	dy := 40
	b.RotateTriangle3DAroundX(p1.AddToY(dy), p2.AddToY(dy), p3.AddToY(dy), theta, focalLength, pixels)
	b.RotateTriangle3DAroundY(p1.AddToY(dy*2), p2.AddToY(dy*2), p3.AddToY(dy*2), theta, focalLength, pixels)
	b.RotateTriangle3DAroundZ(p1.AddToY(dy*3), p2.AddToY(dy*3), p3.AddToY(dy*3), theta, focalLength, pixels)

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

	b.DrawPolygon3D(vertices3d, edges3d, focalLength, pixels)
	b.RotatePolygon3DAroundX(vertices3d2, edges3d, -theta, focalLength, pixels)
	b.RotatePolygon3DAroundY(vertices3d3, edges3d, -theta, focalLength, pixels)
	b.RotatePolygon3DAroundZ(vertices3d4, edges3d, -theta, focalLength, pixels)
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

	// b.DrawPolygon3D(vertices3d, edges3d, focalLength, pixels)
	// b.RotatePolygon3DAroundY(vertices3d, edges3d, 90.0*math.Pi/180.0, focalLength, pixels)
	// b.RotatePolygon3DAroundX(vertices3d, edges3d, -theta, focalLength, pixels)
	// b.RotatePolygon3DAroundY(vertices3d, edges3d, -theta, focalLength, pixels)
	// b.RotatePolygon3DAroundZ(vertices3d, edges3d, -theta, focalLength, pixels)
	b.RotatePolygon3DAroundXAndY(vertices3d, edges3d, -theta, focalLength, pixels)
}

// DrawTriangle draws a triangle
func (b *Board) DrawTriangle(p1, p2, p3 *point.Point2D, pixels []byte) {
	b.DrawLine(p1, p2, pixels)
	b.DrawLine(p2, p3, pixels)
	b.DrawLine(p3, p1, pixels)
}

// RotateTriangle rotates a triangle
func (b *Board) RotateTriangle(p1, p2, p3 *point.Point2D, theta float64, pixels []byte) {
	pc := point.NewPoint2D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3)
	b.DrawTriangle(p1.RotatePoint(pc, theta), p2.RotatePoint(pc, theta), p3.RotatePoint(pc, theta), pixels)
}

// DrawTriangle3D draws a triangle with z axis
func (b *Board) DrawTriangle3D(p1, p2, p3 *point.Point3D, focalLength int, pixels []byte) {
	pc := point.NewPoint2D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3)
	b.DrawLine3D(p1, p2, pc, focalLength, pixels)
	b.DrawLine3D(p2, p3, pc, focalLength, pixels)
	b.DrawLine3D(p3, p1, pc, focalLength, pixels)
}

// RotateTriangle3DAroundX rotates a triangle around X axis
func (b *Board) RotateTriangle3DAroundX(p1, p2, p3 *point.Point3D, theta float64, focalLength int, pixels []byte) {
	pc := point.NewPoint3D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3, (p1.Z+p2.Z+p3.Z)/3)
	b.DrawTriangle3D(p1.RotatePoint3DAroundX(pc, theta), p2.RotatePoint3DAroundX(pc, theta), p3.RotatePoint3DAroundX(pc, theta), focalLength, pixels)
}

// RotateTriangle3DAroundY rotates a triangle around Y axis
func (b *Board) RotateTriangle3DAroundY(p1, p2, p3 *point.Point3D, theta float64, focalLength int, pixels []byte) {
	pc := point.NewPoint3D(p2.X, p1.Y, p1.Z)
	b.DrawTriangle3D(p1.RotatePoint3DAroundY(pc, theta), p2.RotatePoint3DAroundY(pc, theta), p3.RotatePoint3DAroundY(pc, theta), focalLength, pixels)
}

// RotateTriangle3DAroundZ rotates a trianglearound Z axis
func (b *Board) RotateTriangle3DAroundZ(p1, p2, p3 *point.Point3D, theta float64, focalLength int, pixels []byte) {
	pc := point.NewPoint3D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3, (p1.Z+p2.Z+p3.Z)/3)
	b.DrawTriangle3D(p1.RotatePoint3DAroundZ(pc, theta), p2.RotatePoint3DAroundZ(pc, theta), p3.RotatePoint3DAroundZ(pc, theta), focalLength, pixels)
}

// DrawQuadrangle draws a quadrangle
func (b *Board) DrawQuadrangle(p1, p2, p3, p4 *point.Point2D, pixels []byte) {
	b.DrawLine(p1, p2, pixels)
	b.DrawLine(p2, p3, pixels)
	b.DrawLine(p3, p4, pixels)
	b.DrawLine(p4, p1, pixels)
}

// RotateQuadrangle rotates a quadrangle
func (b *Board) RotateQuadrangle(p1, p2, p3, p4 *point.Point2D, theta float64, pixels []byte) {
	pc := point.NewPoint2D((p1.X+p2.X+p3.X+p4.X)/4, (p1.Y+p2.Y+p3.Y+p4.Y)/4)
	b.DrawQuadrangle(p1.RotatePoint(pc, theta), p2.RotatePoint(pc, theta), p3.RotatePoint(pc, theta), p4.RotatePoint(pc, theta), pixels)
}

// DrawPolygon draws a polygon based on vertices and edges matrices
func (b *Board) DrawPolygon(vertices []*point.Point2D, edges [][2]int, pixels []byte) {
	for _, edge := range edges {
		b.DrawLine(vertices[edge[0]], vertices[edge[1]], pixels)
	}
}

// RotatePolygon rotates a polygon
func (b *Board) RotatePolygon(vertices []*point.Point2D, edges [][2]int, theta float64, pixels []byte) {
	var sumX, sumY int
	lv := len(vertices)
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
	}
	pc := point.NewPoint2D(sumX/lv, sumY/lv)

	rotatedVertices := make([]*point.Point2D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, v.RotatePoint(pc, theta))
	}

	b.DrawPolygon(rotatedVertices, edges, pixels)
}

// DrawPolygon3D draws a 3D polygon based on vertices and edges matrices
func (b *Board) DrawPolygon3D(vertices []*point.Point3D, edges [][2]int, focalLength int, pixels []byte) {
	var sumX, sumY int
	lv := len(vertices)
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
	}

	pc := point.NewPoint2D(sumX/lv, sumY/lv)
	for _, edge := range edges {
		b.DrawLine3D(vertices[edge[0]], vertices[edge[1]], pc, focalLength, pixels)
	}
}

// RotatePolygon3DAroundX rotates a polygon around X axis
func (b *Board) RotatePolygon3DAroundX(vertices []*point.Point3D, edges [][2]int, theta float64, focalLength int, pixels []byte) {
	var sumX, sumY, sumZ int
	lv := len(vertices)
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/lv, sumY/lv, sumZ/lv)

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, v.RotatePoint3DAroundX(pc, theta))
	}

	b.DrawPolygon3D(rotatedVertices, edges, focalLength, pixels)
}

// RotatePolygon3DAroundY rotates a polygon around Y axis
func (b *Board) RotatePolygon3DAroundY(vertices []*point.Point3D, edges [][2]int, theta float64, focalLength int, pixels []byte) {
	var sumX, sumY, sumZ int
	lv := len(vertices)
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/lv, sumY/lv, sumZ/lv)

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, v.RotatePoint3DAroundY(pc, theta))
	}

	b.DrawPolygon3D(rotatedVertices, edges, focalLength, pixels)
}

// RotatePolygon3DAroundZ rotates a polygon around Z axis
func (b *Board) RotatePolygon3DAroundZ(vertices []*point.Point3D, edges [][2]int, theta float64, focalLength int, pixels []byte) {
	var sumX, sumY, sumZ int
	lv := len(vertices)
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/lv, sumY/lv, sumZ/lv)

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, v.RotatePoint3DAroundZ(pc, theta))
	}

	b.DrawPolygon3D(rotatedVertices, edges, focalLength, pixels)
}

// RotatePolygon3DAroundXAndY rotates a polygon around X and Y axis
func (b *Board) RotatePolygon3DAroundXAndY(vertices []*point.Point3D, edges [][2]int, theta float64, focalLength int, pixels []byte) {
	var sumX, sumY, sumZ int
	lv := len(vertices)
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/lv, sumY/lv, sumZ/lv)

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, v.RotatePoint3DAroundXAndY(pc, theta))
	}

	b.DrawPolygon3D(rotatedVertices, edges, focalLength, pixels)
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
