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

	triangle := shapes.NewTriangle(point.NewPoint2D(30, 30), point.NewPoint2D(60, 25), point.NewPoint2D(70, 50))
	b.DrawPolygon(triangle.Rotate(theta), pixels)

	quadrangle := shapes.NewQuadrangle(point.NewPoint2D(30, 70), point.NewPoint2D(60, 65), point.NewPoint2D(70, 90), point.NewPoint2D(40, 90))
	b.DrawPolygon(quadrangle.Rotate(-theta), pixels)

	line := shapes.NewLine(point.NewPoint2D(30, 120), point.NewPoint2D(70, 130))
	b.DrawLine(line, pixels)
	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 140, 0), point.NewPoint3D(70, 150, 20), point.NewPoint2D(30, 140), focalLength, pixels)
	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 160, 10), point.NewPoint3D(70, 175, 20), point.NewPoint2D(55, 160), focalLength, pixels)

	dy := 40
	p1 := point.NewPoint2D(80, 20)
	p2 := point.NewPoint2D(100, 40)
	p3 := point.NewPoint2D(120, 20)

	triangle = shapes.NewTriangle(p1, p2, p3)
	b.DrawPolygon(&triangle, pixels)

	triangle = shapes.NewTriangle(p1.AddToY(dy), p2.AddToY(dy), p3.AddToY(dy))
	b.DrawPolygon3D(triangle.RotateAroundX(theta), focalLength, pixels)

	triangle = shapes.NewTriangle(p1.AddToY(dy*2), p2.AddToY(dy*2), p3.AddToY(dy*2))
	b.DrawPolygon3D(triangle.RotateAroundY(theta), focalLength, pixels)

	triangle = shapes.NewTriangle(p1.AddToY(dy*3), p2.AddToY(dy*3), p3.AddToY(dy*3))
	b.DrawPolygon3D(triangle.RotateAroundZ(theta), focalLength, pixels)

	vertices3d := []point.Point3D{
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
	vertices3d2 := []point.Point3D{
		vertices3d[0].AddToY(dy),
		vertices3d[1].AddToY(dy),
		vertices3d[2].AddToY(dy),
		vertices3d[3].AddToY(dy),
		vertices3d[4].AddToY(dy),
		vertices3d[5].AddToY(dy),
		vertices3d[6].AddToY(dy),
		vertices3d[7].AddToY(dy),
	}
	vertices3d3 := []point.Point3D{
		vertices3d2[0].AddToY(dy),
		vertices3d2[1].AddToY(dy),
		vertices3d2[2].AddToY(dy),
		vertices3d2[3].AddToY(dy),
		vertices3d2[4].AddToY(dy),
		vertices3d2[5].AddToY(dy),
		vertices3d2[6].AddToY(dy),
		vertices3d2[7].AddToY(dy),
	}
	vertices3d4 := []point.Point3D{
		vertices3d3[0].AddToY(dy),
		vertices3d3[1].AddToY(dy),
		vertices3d3[2].AddToY(dy),
		vertices3d3[3].AddToY(dy),
		vertices3d3[4].AddToY(dy),
		vertices3d3[5].AddToY(dy),
		vertices3d3[6].AddToY(dy),
		vertices3d3[7].AddToY(dy),
	}

	polygon3d := shapes.NewPolygon3D(vertices3d, edges3d)
	b.DrawPolygon3D(&polygon3d, focalLength, pixels)

	polygon3d = shapes.NewPolygon3D(vertices3d2, edges3d)
	b.DrawPolygon3D(polygon3d.RotateAroundX(-theta), focalLength, pixels)

	polygon3d = shapes.NewPolygon3D(vertices3d3, edges3d)
	b.DrawPolygon3D(polygon3d.RotateAroundY(-theta), focalLength, pixels)

	polygon3d = shapes.NewPolygon3D(vertices3d4, edges3d)
	b.DrawPolygon3D(polygon3d.RotateAroundZ(-theta), focalLength, pixels)
}

// Draw2 draws board
func (b *Board) Draw2(pixels []byte, counter, focalLength int) {
	theta := float64((int(1.5*float64(counter)))%360) * math.Pi / 180.0
	_ = theta

	vertices3d := []point.Point3D{
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

	polygon3d := shapes.NewPolygon3D(vertices3d, edges3d)
	b.DrawPolygon3D(polygon3d.RotateAroundX(-theta).RotateAroundY(-theta), focalLength, pixels)
}

// DrawPolygon draws a polygon based on vertices and edges matrices
func (b *Board) DrawPolygon(polygon *shapes.Polygon2D, pixels []byte) {
	for _, edge := range polygon.Edges {
		line := shapes.NewLine(polygon.Vertices[edge[0]], polygon.Vertices[edge[1]])
		b.DrawLine(line, pixels)
	}
}

// DrawPolygon3D draws a 3D polygon based on vertices and edges matrices
func (b *Board) DrawPolygon3D(polygon *shapes.Polygon3D, focalLength int, pixels []byte) {
	pc := polygon.CalculateFlatCenterPoint()
	for _, edge := range polygon.Edges {
		line := shapes.NewLine3D(polygon.Vertices[edge[0]], polygon.Vertices[edge[1]])
		b.DrawLine3D(line, pc, focalLength, pixels)
	}
}

// DrawLine3DRelativeToPoint draws a line
func (b *Board) DrawLine3DRelativeToPoint(p1, p2 point.Point3D, pc point.Point2D, focalLength int, pixels []byte) {
	line := shapes.NewLine3D(p1, p2)
	b.DrawLine3D(line, pc, focalLength, pixels)
}

// DrawLine draws a line
func (b *Board) DrawLine(l shapes.Line, pixels []byte) {
	for _, p := range l.PlotLine() {
		b.colorPixel(p.X, p.Y, pixels)
	}
}

// DrawLine3D draws a 3D line
func (b *Board) DrawLine3D(l shapes.Line3D, pc point.Point2D, focalLength int, pixels []byte) {
	for _, p := range l.PlotLine(pc, focalLength) {
		b.color3DPixel(p.X, p.Y, p.D, pixels)
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
	x = keepCoordBetweenMinAndMax(x, 0, b.width)
	y = keepCoordBetweenMinAndMax(y, 0, b.height)

	max := b.width * b.height
	i := keepCoordBetweenMinAndMax(y*b.width+x, 0, max)
	return i
}

func keepCoordBetweenMinAndMax(coord, min, max int) int {
	for coord < min || coord >= max {
		switch {
		case coord < min:
			coord += max
		case coord >= max:
			coord -= max
		}
	}
	return coord
}
