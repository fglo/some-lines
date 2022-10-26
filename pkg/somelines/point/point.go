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

func (p *Point2D) AddToX(dx int) Point2D {
	return Point2D{p.X + dx, p.Y}
}

func (p *Point2D) AddToY(dy int) Point2D {
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

func (p *Point3D) To2D(focalLength int) Point2D {
	xp := int(float64(p.X*focalLength) / float64(p.Z+focalLength))
	yp := int(float64(p.Y*focalLength) / float64(p.Z+focalLength))

	return Point2D{xp, yp}
}

func (p *Point3D) To2DRelativeToPoint(pc Point2D, focalLength int) Point2D {
	x := p.X - pc.X
	y := p.Y - pc.Y

	xp := int(float64(x*focalLength) / float64(p.Z+focalLength))
	yp := int(float64(y*focalLength) / float64(p.Z+focalLength))

	return Point2D{xp + pc.X, yp + pc.Y}
}

func (p *Point3D) AddToX(dx int) Point3D {
	return Point3D{p.X + dx, p.Y, p.Z}
}

func (p *Point3D) AddToY(dy int) Point3D {
	return Point3D{p.X, p.Y + dy, p.Z}
}

func (p *Point3D) AddToZ(dz int) Point3D {
	return Point3D{p.X, p.Y, p.Z + dz}
}

func (p *Point3D) RotateAroundX(pc Point3D, theta float64) Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X
	y := p.Y - pc.Y
	z := p.Z - pc.Z

	yr := int(float64(y)*cosTheta - float64(z)*sinTheta)
	zr := int(float64(y)*sinTheta + float64(z)*cosTheta)

	return Point3D{x, yr + pc.Y, zr + pc.Z}
}

func (p *Point3D) RotateAroundY(pc Point3D, theta float64) Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y
	z := p.Z - pc.Z

	xr := int(float64(x)*cosTheta + float64(z)*sinTheta)
	zr := int(-float64(x)*sinTheta + float64(z)*cosTheta)

	return Point3D{xr + pc.X, y, zr + pc.Z}
}

func (p *Point3D) RotateAroundZ(pc Point3D, theta float64) Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y - pc.Y
	z := p.Z

	xr := int(float64(x)*cosTheta - float64(y)*sinTheta)
	yr := int(float64(x)*sinTheta + float64(y)*cosTheta)

	return Point3D{xr + pc.X, yr + pc.Y, z}
}

type Point2DWithDepth struct {
	X int
	Y int
	D float64
}

func NewPointPoint2DWithDepth(x, y int, d float64) Point2DWithDepth {
	p := Point2DWithDepth{X: x, Y: y, D: d}
	return p
}
