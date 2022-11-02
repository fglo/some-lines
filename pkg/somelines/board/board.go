package board

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/projector"
	"github.com/fglo/some-lines/pkg/somelines/renderer"
	"github.com/fglo/some-lines/pkg/somelines/scene"
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

	b.shaded = true

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
// func (b *Board) Draw(pixels []byte, counter, focalLength int) {
// 	theta := float64((int(1.5*float64(counter)))%360) * math.Pi / 180.0

// 	triangle := shapes.NewTriangle(point.NewPoint2D(30, 30), point.NewPoint2D(60, 25), point.NewPoint2D(70, 50))
// 	b.DrawPolygon(triangle.Rotate(theta), pixels)

// 	quadrangle := shapes.NewQuadrangle(point.NewPoint2D(30, 70), point.NewPoint2D(60, 65), point.NewPoint2D(70, 90), point.NewPoint2D(40, 90))
// 	b.DrawPolygon(quadrangle.Rotate(-theta), pixels)

// 	line := shapes.NewLine(point.NewPoint2D(30, 120), point.NewPoint2D(70, 130))
// 	b.DrawLine(line, pixels)
// 	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 140, 0), point.NewPoint3D(70, 150, 20), point.NewPoint2D(30, 140), focalLength, pixels)
// 	b.DrawLine3DRelativeToPoint(point.NewPoint3D(30, 160, 10), point.NewPoint3D(70, 175, 20), point.NewPoint2D(55, 160), focalLength, pixels)

// 	dy := 40
// 	p1 := point.NewPoint2D(80, 20)
// 	p2 := point.NewPoint2D(100, 40)
// 	p3 := point.NewPoint2D(120, 20)

// 	triangle = shapes.NewTriangle(p1, p2, p3)
// 	b.DrawPolygon(&triangle, pixels)

// 	triangle = triangle.MoveAlongY(dy)
// 	b.DrawPolygon3D(triangle.RotateAroundX(theta), focalLength, pixels)

// 	triangle = triangle.MoveAlongY(dy)
// 	b.DrawPolygon3D(triangle.RotateAroundY(theta), focalLength, pixels)

// 	triangle = triangle.MoveAlongY(dy)
// 	b.DrawPolygon3D(triangle.RotateAroundZ(theta), focalLength, pixels)

// 	cube := shapes.NewCube(point.NewPoint3D(140, 20, 0), point.NewPoint3D(170, 50, 30))
// 	b.DrawPolygon3D(&cube, focalLength, pixels)

// 	cube = *cube.MoveAlongYButPointer(dy)
// 	b.DrawPolygon3D(cube.RotateAroundX(-theta), focalLength, pixels)

// 	cube = cube.MoveAlongY(dy)
// 	b.DrawPolygon3D(cube.RotateAroundY(-theta), focalLength, pixels)

// 	cube = cube.MoveAlongY(dy)
// 	b.DrawPolygon3D(cube.RotateAroundZ(-theta), focalLength, pixels)

// 	vertices3d := []point.Point3D{
// 		point.NewPoint3D(50, 40, 20),
// 		point.NewPoint3D(130, 40, 20),
// 		point.NewPoint3D(60, 100, 20),
// 		point.NewPoint3D(120, 100, 20),
// 		point.NewPoint3D(70, 40, 80),
// 		point.NewPoint3D(150, 40, 80),
// 		point.NewPoint3D(60, 100, 80),
// 		point.NewPoint3D(120, 100, 80),
// 	}
// 	edges3d := [][2]int{
// 		{0, 1}, {0, 2}, {0, 4},
// 		{3, 1}, {3, 2}, {3, 7},
// 		{5, 1}, {5, 4}, {5, 7},
// 		{6, 2}, {6, 4}, {6, 7},
// 	}

// 	polygon3d := shapes.NewPolygon3D(vertices3d, edges3d)

// 	d := 30
// 	dt := d + counter%b.width
// 	ds := int(float64(d) * math.Sin(theta))
// 	b.DrawPolygon3D(polygon3d.MoveAlongXButPointer(dt).MoveAlongZButPointer(ds).RotateAroundX(-theta).RotateAroundY(-theta), focalLength, pixels)
// }

func (b *Board) DrawScene(pixels []byte, counter, focalLength int) {
	theta := float64(counter%360) * math.Pi / 180.0
	d := 50
	dt := d + counter%b.width
	_ = dt
	ds := int(float64(d) * math.Sin(theta))
	dc := int(float64(d) * math.Cos(theta))
	_ = ds
	_ = dc

	dthetax := (15 * (math.Sin(theta) + 1)) * math.Pi / 180.0
	_ = dthetax
	dy := int(-80 * (math.Sin(theta) + 1))
	_ = dy

	cameraPosition := point.NewPoint3D(b.width/2, b.height/2, 0)
	cameraOrientation := point.NewOrientation(1*math.Pi/180.0, 0, 0)

	camera := camera.New(cameraPosition, cameraOrientation)
	// camera.SetFoV(46)
	// camera.SetFoV(84)

	scene := scene.New()
	scene.AddCamera("main", &camera)

	cube := shapes.NewCube(point.NewPoint3D(500, 400, 200), point.NewPoint3D(700, 600, 400))
	scene.AddPolygon3D(cube)

	cube = cube.MoveAlongY(250)
	// scene.AddPolygon3D(cube)
	scene.AddPolygon3D(*cube.RotateAroundX(-theta))

	cube = cube.MoveAlongY(250)
	// scene.AddPolygon3D(cube)
	scene.AddPolygon3D(*cube.RotateAroundY(-theta))

	cube = cube.MoveAlongY(250)
	// scene.AddPolygon3D(cube)
	scene.AddPolygon3D(*cube.RotateAroundZ(-theta))

	vertices3d := []point.Point3D{
		point.NewPoint3D(50, 640, 200),
		point.NewPoint3D(130, 640, 200),
		point.NewPoint3D(60, 700, 200),
		point.NewPoint3D(120, 700, 200),
		point.NewPoint3D(70, 640, 260),
		point.NewPoint3D(150, 640, 260),
		point.NewPoint3D(60, 700, 260),
		point.NewPoint3D(120, 700, 260),
	}
	edges3d := [][2]int{
		{0, 1}, {0, 2}, {0, 4},
		{3, 1}, {3, 2}, {3, 7},
		{5, 1}, {5, 4}, {5, 7},
		{6, 2}, {6, 4}, {6, 7},
	}

	polygon3d := shapes.NewPolygon3D(vertices3d, edges3d)
	_ = polygon3d
	scene.AddPolygon3D(*polygon3d.MoveAlongXButPointer(dt).MoveAlongZButPointer(ds).RotateAroundX(-theta).RotateAroundY(-theta))

	projector := projector.NewPerspectiveProjector()
	// projector := projector.NewOrthogonalProjector()
	renderer := renderer.New(projector)
	renderer.RenderScene(scene, b.width, b.height, pixels)

	b.drawCrosshair(cameraPosition, pixels)
}

func (b *Board) DrawScene2(pixels []byte, counter, focalLength int) {
	theta := float64(counter%360) * math.Pi / 180.0
	d := 50
	dt := d + counter%b.width
	_ = dt
	ds := int(float64(d) * math.Sin(theta))
	dc := int(float64(d) * math.Cos(theta))
	_ = ds
	_ = dc

	dthetax := 15 * (math.Sin(theta) + 1) * math.Pi / 180.0
	_ = dthetax
	dtheta := 15 * math.Sin(theta) * math.Pi / 180.0
	_ = dtheta
	dy := int(-80 * (math.Sin(theta) + 1))
	_ = dy

	cameraPosition := point.NewPoint3D(b.width/2, b.height/2, 0)
	// cameraPosition = cameraPosition.MoveAlongZ(ds)
	// cameraPosition = cameraPosition.MoveAlongYButPointer(dy).MoveAlongZ(ds)
	// cameraPosition = cameraPosition.MoveAlongXButPointer(ds).MoveAlongY(dc)
	// cameraPosition = cameraPosition.MoveAlongZButPointer(ds).MoveAlongXButPointer(ds).MoveAlongY(dc)
	// cameraOrientation := point.NewOrientation(dthetax, 0, 0)
	cameraOrientation := point.NewOrientation(0, dtheta, 0)

	camera := camera.New(cameraPosition, cameraOrientation)
	// camera.SetFoV(46)
	// camera.SetFoV(53)

	scene := scene.New()
	scene.AddCamera("main", &camera)

	plane := shapes.NewPlane(point.NewPoint3D(b.width/2, 2*b.height/3, 300))
	scene.AddPolygon3D(plane)

	cube := shapes.NewCube(point.NewPoint3D(b.width/2-100, 2*b.height/3-200, 200), point.NewPoint3D(b.width/2+100, 2*b.height/3, 400))
	cube = *cube.RotateAroundY(theta)
	scene.AddPolygon3D(cube)

	projector := projector.NewPerspectiveProjector()
	// projector := projector.NewOrthogonalProjector()
	renderer := renderer.New(projector)
	renderer.RenderScene(scene, b.width, b.height, pixels)

	b.drawCrosshair(cameraPosition, pixels)
}

func (b *Board) drawCrosshair(cameraPosition point.Point3D, pixels []byte) {
	b.colorPixel(cameraPosition.X, cameraPosition.Y, pixels)
	b.colorPixel(cameraPosition.X+1, cameraPosition.Y, pixels)
	b.colorPixel(cameraPosition.X-1, cameraPosition.Y, pixels)
	b.colorPixel(cameraPosition.X, cameraPosition.Y+1, pixels)
	b.colorPixel(cameraPosition.X, cameraPosition.Y-1, pixels)
}

func (b *Board) colorPixel(x, y int, pixels []byte) {
	if x >= 0 && x < b.width && y > 0 && y < b.height {
		i := b.getPixelsIndex(x, y)
		pixels[4*i] = 0xf0
		pixels[4*i+1] = 0xf0
		pixels[4*i+2] = 0xf0
		pixels[4*i+3] = 0xff
	}
}

func (b *Board) color3DPixel(x, y int, z float64, pixels []byte) {
	if x >= 0 && x < b.width && y > 0 && y < b.height {
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
