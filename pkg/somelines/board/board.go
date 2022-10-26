package board

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/point"
)

// Board encapsulates simulation logic
type Board struct {
	width  int
	height int

	paused    bool
	forwarded bool
	// reversed  bool
}

// New is a Board constructor
func New(w, h int) *Board {
	b := new(Board)

	b.width = w
	b.height = h

	return b
}

// Setup prepares board
func (b *Board) Setup(numberOfCells int) {
	b.paused = false
}

// TogglePause toggles board pause
func (b *Board) TogglePause() {
	b.paused = !b.paused
}

// Forward sets forward
func (b *Board) Forward(forward bool) {
	b.forwarded = forward
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
func (b *Board) Draw(pix []byte, counter, focalLength int) {
	theta := float64((2*counter)%360) * math.Pi / 180.0
	// b.RotateTriangle(point.NewPoint2D(30, 30), point.NewPoint2D(60, 25), point.NewPoint2D(40, 50), theta, pix)
	b.RotateQuadrangle(point.NewPoint2D(30, 30), point.NewPoint2D(60, 25), point.NewPoint2D(70, 50), point.NewPoint2D(40, 50), theta, pix)

	vertices := []*point.Point2D{point.NewPoint2D(30, 70), point.NewPoint2D(60, 65), point.NewPoint2D(70, 90), point.NewPoint2D(40, 90)}
	edges := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}
	b.RotatePolygon(vertices, edges, -theta, pix)

	b.DrawLine(point.NewPoint2D(30, 120), point.NewPoint2D(70, 130), pix)
	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 140, 0), point.NewPoint3D(70, 150, 20), point.NewPoint2D(30, 140), focalLength, pix)

	p1 := point.NewPoint3D(80, 20, 0)
	p2 := point.NewPoint3D(100, 40, 0)
	p3 := point.NewPoint3D(120, 20, 0)
	b.DrawTriangle(point.NewPoint2D(80, 20), point.NewPoint2D(100, 40), point.NewPoint2D(120, 20), pix)
	// b.DrawTriangle3D(p1, p2, p3, focalLength, pix)
	dy := 40
	b.RotateTriangle3DAroundX(p1.AddToY(dy), p2.AddToY(dy), p3.AddToY(dy), theta, focalLength, pix)
	b.RotateTriangle3DAroundY(p1.AddToY(dy*2), p2.AddToY(dy*2), p3.AddToY(dy*2), theta, focalLength, pix)
	b.RotateTriangle3DAroundZ(p1.AddToY(dy*3), p2.AddToY(dy*3), p3.AddToY(dy*3), theta, focalLength, pix)

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

	b.DrawPolygon3D(vertices3d, edges3d, focalLength, pix)
	b.RotatePolygon3DAroundX(vertices3d2, edges3d, -theta, focalLength, pix)
	b.RotatePolygon3DAroundY(vertices3d3, edges3d, -theta, focalLength, pix)
	b.RotatePolygon3DAroundZ(vertices3d4, edges3d, -theta, focalLength, pix)
}

// DrawTriangle draws a triangle
func (b *Board) DrawTriangle(p1, p2, p3 *point.Point2D, pix []byte) {
	b.DrawLine(p1, p2, pix)
	b.DrawLine(p2, p3, pix)
	b.DrawLine(p3, p1, pix)
}

// RotateTriangle rotates a triangle
func (b *Board) RotateTriangle(p1, p2, p3 *point.Point2D, theta float64, pix []byte) {
	pc := point.NewPoint2D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3)
	b.DrawTriangle(rotatePoint(p1, pc, theta), rotatePoint(p2, pc, theta), rotatePoint(p3, pc, theta), pix)
}

// DrawTriangle3D draws a triangle with z axis
func (b *Board) DrawTriangle3D(p1, p2, p3 *point.Point3D, focalLength int, pix []byte) {
	pc := point.NewPoint2D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3)
	// _ = pc
	// b.DrawTriangle(p1.To2D(focalLength), p2.To2D(focalLength), p3.To2D(focalLength), pix)
	b.DrawTriangle(p1.To2DRelativeToPoint(pc, focalLength), p2.To2DRelativeToPoint(pc, focalLength), p3.To2DRelativeToPoint(pc, focalLength), pix)
}

// RotateTriangle3DAroundX rotates a triangle around X axis
func (b *Board) RotateTriangle3DAroundX(p1, p2, p3 *point.Point3D, theta float64, focalLength int, pix []byte) {
	pc := point.NewPoint3D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3, (p1.Z+p2.Z+p3.Z)/3)
	// pc := point.NewPoint3D(p2.X, p1.Y, p1.Z)
	b.DrawTriangle3D(rotatePoint3DAroundX(p1, pc, theta), rotatePoint3DAroundX(p2, pc, theta), rotatePoint3DAroundX(p3, pc, theta), focalLength, pix)
}

// RotateTriangle3DAroundY rotates a triangle around Y axis
func (b *Board) RotateTriangle3DAroundY(p1, p2, p3 *point.Point3D, theta float64, focalLength int, pix []byte) {
	// pc := point.NewPoint3D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3, (p1.Z+p2.Z+p3.Z)/3)
	pc := point.NewPoint3D(p2.X, p1.Y, p1.Z)
	b.DrawTriangle3D(rotatePoint3DAroundY(p1, pc, theta), rotatePoint3DAroundY(p2, pc, theta), rotatePoint3DAroundY(p3, pc, theta), focalLength, pix)
}

// RotateTriangle3DAroundZ rotates a trianglearound Z axis
func (b *Board) RotateTriangle3DAroundZ(p1, p2, p3 *point.Point3D, theta float64, focalLength int, pix []byte) {
	pc := point.NewPoint3D((p1.X+p2.X+p3.X)/3, (p1.Y+p2.Y+p3.Y)/3, (p1.Z+p2.Z+p3.Z)/3)
	// pc := point.NewPoint3D(p2.X, p1.Y, p1.Z)
	b.DrawTriangle3D(rotatePoint3DAroundZ(p1, pc, theta), rotatePoint3DAroundZ(p2, pc, theta), rotatePoint3DAroundZ(p3, pc, theta), focalLength, pix)
}

// DrawQuadrangle draws a quadrangle
func (b *Board) DrawQuadrangle(p1, p2, p3, p4 *point.Point2D, pix []byte) {
	b.DrawLine(p1, p2, pix)
	b.DrawLine(p2, p3, pix)
	b.DrawLine(p3, p4, pix)
	b.DrawLine(p4, p1, pix)
}

// RotateQuadrangle rotates a quadrangle
func (b *Board) RotateQuadrangle(p1, p2, p3, p4 *point.Point2D, theta float64, pix []byte) {
	pc := point.NewPoint2D((p1.X+p2.X+p3.X+p4.X)/4, (p1.Y+p2.Y+p3.Y+p4.Y)/4)
	b.DrawQuadrangle(rotatePoint(p1, pc, theta), rotatePoint(p2, pc, theta), rotatePoint(p3, pc, theta), rotatePoint(p4, pc, theta), pix)
}

// DrawPolygon draws a polygon based on vertices and edges matrices
func (b *Board) DrawPolygon(vertices []*point.Point2D, edges [][2]int, pix []byte) {
	for _, edge := range edges {
		b.DrawLine(vertices[edge[0]], vertices[edge[1]], pix)
	}
}

// RotatePolygon rotates a polygon
func (b *Board) RotatePolygon(vertices []*point.Point2D, edges [][2]int, theta float64, pix []byte) {
	sumX := 0
	sumY := 0
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
	}
	pc := point.NewPoint2D(sumX/len(vertices), sumY/len(vertices))

	rotatedVertices := make([]*point.Point2D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, rotatePoint(v, pc, theta))
	}

	b.DrawPolygon(rotatedVertices, edges, pix)
}

// DrawPolygon3D draws a 3D polygon based on vertices and edges matrices
func (b *Board) DrawPolygon3D(vertices []*point.Point3D, edges [][2]int, focalLength int, pix []byte) {
	var sumX, sumY int
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
	}

	pc := point.NewPoint2D(sumX/len(vertices), sumY/len(vertices))
	for _, edge := range edges {
		b.DrawLine3DRelativeToPoint(vertices[edge[0]], vertices[edge[1]], pc, focalLength, pix)
	}
}

// RotatePolygon3DAroundX rotates a polygon around X axis
func (b *Board) RotatePolygon3DAroundX(vertices []*point.Point3D, edges [][2]int, theta float64, focalLength int, pix []byte) {
	var sumX, sumY, sumZ int
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/len(vertices), sumY/len(vertices), sumZ/len(vertices))
	pc2d := point.NewPoint2D(sumX/len(vertices), sumY/len(vertices))

	rotatedVertices := make([]*point.Point2D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, rotatePoint3DAroundX(v, pc, theta).To2DRelativeToPoint(pc2d, focalLength))
	}

	b.DrawPolygon(rotatedVertices, edges, pix)
}

// RotatePolygon3DAroundY rotates a polygon around Y axis
func (b *Board) RotatePolygon3DAroundY(vertices []*point.Point3D, edges [][2]int, theta float64, focalLength int, pix []byte) {
	var sumX, sumY, sumZ int
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/len(vertices), sumY/len(vertices), sumZ/len(vertices))
	pc2d := point.NewPoint2D(sumX/len(vertices), sumY/len(vertices))

	rotatedVertices := make([]*point.Point2D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, rotatePoint3DAroundY(v, pc, theta).To2DRelativeToPoint(pc2d, focalLength))
	}

	b.DrawPolygon(rotatedVertices, edges, pix)
}

// RotatePolygon3DAroundZ rotates a polygon around Z axis
func (b *Board) RotatePolygon3DAroundZ(vertices []*point.Point3D, edges [][2]int, theta float64, focalLength int, pix []byte) {
	var sumX, sumY, sumZ int
	for _, v := range vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/len(vertices), sumY/len(vertices), sumZ/len(vertices))
	pc2d := point.NewPoint2D(sumX/len(vertices), sumY/len(vertices))

	rotatedVertices := make([]*point.Point2D, 0)
	for _, v := range vertices {
		rotatedVertices = append(rotatedVertices, rotatePoint3DAroundZ(v, pc, theta).To2DRelativeToPoint(pc2d, focalLength))
	}

	b.DrawPolygon(rotatedVertices, edges, pix)
}

// DrawLine3DRelativeToPoint draws a line
func (b *Board) DrawLine3DRelativeToPoint(p1, p2 *point.Point3D, pc *point.Point2D, focalLength int, pix []byte) {
	b.DrawLine(p1.To2DRelativeToPoint(pc, focalLength), p2.To2DRelativeToPoint(pc, focalLength), pix)
}

// DrawLine3D draws a line
func (b *Board) DrawLine3D(p1, p2 *point.Point3D, focalLength int, pix []byte) {
	pc := point.NewPoint2D(b.width/2, b.height/2)
	b.DrawLine(p1.To2DRelativeToPoint(pc, focalLength), p2.To2DRelativeToPoint(pc, focalLength), pix)
}

// DrawLine draws a line
func (b *Board) DrawLine(p1, p2 *point.Point2D, pix []byte) {
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
		for p0.X != p2.X {
			b.colorPixel(p0.X, p0.Y, pix)
			if d >= 0 {
				p0.Y += sy
				d += ai
			} else {
				d += bi
			}
			p0.X += sx
		}
	case dx <= dy:
		ai := (dx - dy) * 2
		bi := dx * 2
		d := bi - dy
		for p0.Y != p2.Y {
			b.colorPixel(p0.X, p0.Y, pix)
			if d >= 0 {
				p0.X += sx
				d += ai
			} else {
				d += bi
			}
			p0.Y += sy
		}
	}
}

func (b *Board) getPixelsIndex(x, y int) int {
	if x < 0 {
		x = b.width + x
	}
	if x >= b.width {
		x = x - b.width
	}

	if y < 0 {
		y = b.height + y
	}
	if y >= b.height {
		y = y - b.height
	}

	max := b.width * b.height
	i := y*b.width + x
	if i < 0 {
		i = max + i
	}
	if i >= max {
		i = i - max
	}

	return i
}

func (b *Board) colorPixel(x, y int, pix []byte) {
	i := b.getPixelsIndex(x, y)
	pix[4*i] = 0xf0
	pix[4*i+1] = 0xf0
	pix[4*i+2] = 0xf0
	pix[4*i+3] = 0xff
}

func rotatePoint(p, pc *point.Point2D, theta float64) *point.Point2D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y - pc.Y

	xr := int(float64(x)*cosTheta - float64(y)*sinTheta)
	yr := int(float64(x)*sinTheta + float64(y)*cosTheta)

	pr := point.NewPoint2D(xr+pc.X, yr+pc.Y)
	return pr
}

func rotatePoint3DAroundX(p, pc *point.Point3D, theta float64) *point.Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X
	y := p.Y - pc.Y
	z := p.Z - pc.Z

	yr := int(float64(y)*cosTheta - float64(z)*sinTheta)
	zr := int(float64(y)*sinTheta + float64(z)*cosTheta)

	pr := point.NewPoint3D(x, yr+pc.Y, zr+pc.Z)
	return pr
}

func rotatePoint3DAroundY(p, pc *point.Point3D, theta float64) *point.Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y
	z := p.Z - pc.Z

	xr := int(float64(x)*cosTheta + float64(z)*sinTheta)
	zr := int(-float64(x)*sinTheta + float64(z)*cosTheta)

	pr := point.NewPoint3D(xr+pc.X, y, zr+pc.Z)
	return pr
}

func rotatePoint3DAroundZ(p, pc *point.Point3D, theta float64) *point.Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y - pc.Y
	z := p.Z

	xr := int(float64(x)*cosTheta - float64(y)*sinTheta)
	yr := int(float64(x)*sinTheta + float64(y)*cosTheta)

	pr := point.NewPoint3D(xr+pc.X, yr+pc.Y, z)
	return pr
}
