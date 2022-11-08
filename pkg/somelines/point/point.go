package point

import "math"

type Point2D struct {
	X int
	Y int
}

func NewPoint2D(x, y int) Point2D {
	p := Point2D{X: x, Y: y}
	return p
}

func (p *Point2D) Clone() Point2D {
	return *p
}

func (p *Point2D) MoveAlongX(dx int) Point2D {
	return Point2D{p.X + dx, p.Y}
}

func (p *Point2D) MoveAlongY(dy int) Point2D {
	return Point2D{p.X, p.Y + dy}
}

func (p *Point2D) Rotate(pc Point2D, theta float64) Point2D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y - pc.Y

	xr := int(float64(x)*cosTheta - float64(y)*sinTheta)
	yr := int(float64(x)*sinTheta + float64(y)*cosTheta)

	return Point2D{xr + pc.X, yr + pc.Y}
}

func (p *Point2D) To3D() Point3D {
	return Point3D{p.X, p.Y, 0}
}

type Point3D struct {
	X int
	Y int
	Z int
}

func NewPoint3D(x, y, z int) Point3D {
	p := Point3D{X: x, Y: y, Z: z}
	return p
}

func (p *Point3D) Clone() Point3D {
	return *p
}

func (p *Point3D) MoveAlongX(dx int) Point3D {
	return Point3D{p.X + dx, p.Y, p.Z}
}

func (p *Point3D) MoveAlongY(dy int) Point3D {
	return Point3D{p.X, p.Y + dy, p.Z}
}

func (p *Point3D) MoveAlongZ(dz int) Point3D {
	return Point3D{p.X, p.Y, p.Z + dz}
}

func (p *Point3D) MoveAlongXButPointer(dx int) *Point3D {
	return &Point3D{p.X + dx, p.Y, p.Z}
}

func (p *Point3D) MoveAlongYButPointer(dy int) *Point3D {
	return &Point3D{p.X, p.Y + dy, p.Z}
}

func (p *Point3D) MoveAlongZButPointer(dz int) *Point3D {
	return &Point3D{p.X, p.Y, p.Z + dz}
}

func (p *Point3D) RotateAroundX(theta float64) *Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	yr := int(float64(p.Y)*cosTheta - float64(p.Z)*sinTheta)
	zr := int(float64(p.Y)*sinTheta + float64(p.Z)*cosTheta)

	return &Point3D{p.X, yr, zr}
}

func (p *Point3D) RotateAroundXRelativeToPoint(pc Point3D, theta float64) Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X
	y := p.Y - pc.Y
	z := p.Z - pc.Z

	yr := int(float64(y)*cosTheta - float64(z)*sinTheta)
	zr := int(float64(y)*sinTheta + float64(z)*cosTheta)

	return Point3D{x, yr + pc.Y, zr + pc.Z}
}

func (p *Point3D) RotateAroundY(theta float64) *Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	xr := int(float64(p.X)*cosTheta + float64(p.Z)*sinTheta)
	zr := int(-float64(p.X)*sinTheta + float64(p.Z)*cosTheta)

	return &Point3D{xr, p.Y, zr}
}

func (p *Point3D) RotateAroundYRelativeToPoint(pc Point3D, theta float64) Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y
	z := p.Z - pc.Z

	xr := int(float64(x)*cosTheta + float64(z)*sinTheta)
	zr := int(-float64(x)*sinTheta + float64(z)*cosTheta)

	return Point3D{xr + pc.X, y, zr + pc.Z}
}

func (p *Point3D) RotateAroundZ(theta float64) *Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	xr := int(float64(p.X)*cosTheta - float64(p.Y)*sinTheta)
	yr := int(float64(p.X)*sinTheta + float64(p.Y)*cosTheta)

	return &Point3D{xr, yr, p.Z}
}

func (p *Point3D) RotateAroundZRelativeToPoint(pc Point3D, theta float64) Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y - pc.Y
	z := p.Z

	xr := int(float64(x)*cosTheta - float64(y)*sinTheta)
	yr := int(float64(x)*sinTheta + float64(y)*cosTheta)

	return Point3D{xr + pc.X, yr + pc.Y, z}
}

type Point3DNdc struct {
	X float64
	Y float64
	Z int
}

func NewPoint3DNdc(x, y float64, z int) Point3DNdc {
	p := Point3DNdc{X: x, Y: y, Z: z}
	return p
}

func (p *Point3DNdc) Clone() Point3DNdc {
	return *p
}

type ProjectedPoint3D struct {
	X int
	Y int
	D float64
}

func NewProjectedPoint3D(x, y int, d float64) ProjectedPoint3D {
	p := ProjectedPoint3D{X: x, Y: y, D: d}
	return p
}

type Orientation struct {
	X float64
	Y float64
	Z float64
}

func NewOrientation(x, y, z float64) Orientation {
	o := Orientation{X: x, Y: y, Z: z}
	return o
}
