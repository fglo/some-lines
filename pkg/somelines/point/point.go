package point

import "math"

type Point2D struct {
	X int
	Y int
}

func NewPoint2D(x, y int) *Point2D {
	p := new(Point2D)
	p.X = x
	p.Y = y

	return p
}

func (p *Point2D) RotatePoint(pc *Point2D, theta float64) *Point2D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y - pc.Y

	xr := int(float64(x)*cosTheta - float64(y)*sinTheta)
	yr := int(float64(x)*sinTheta + float64(y)*cosTheta)

	pr := NewPoint2D(xr+pc.X, yr+pc.Y)
	return pr
}

type Point3D struct {
	X int
	Y int
	Z int
}

func NewPoint3D(x, y, z int) *Point3D {
	p := new(Point3D)
	p.X = x
	p.Y = y
	p.Z = z

	return p
}

func (p *Point3D) To2D(focalLength int) *Point2D {
	xp := int(float64(p.X*focalLength) / float64(p.Z+focalLength))
	yp := int(float64(p.Y*focalLength) / float64(p.Z+focalLength))

	return NewPoint2D(xp, yp)
}

func (p *Point3D) To2DRelativeToPoint(pc *Point2D, focalLength int) *Point2D {
	x := p.X - pc.X
	y := p.Y - pc.Y

	xp := int(float64(x*focalLength) / float64(p.Z+focalLength))
	yp := int(float64(y*focalLength) / float64(p.Z+focalLength))

	return NewPoint2D(xp+pc.X, yp+pc.Y)
}

func (p *Point3D) AddToX(dx int) *Point3D {
	return NewPoint3D(p.X+dx, p.Y, p.Z)
}

func (p *Point3D) AddToY(dy int) *Point3D {
	return NewPoint3D(p.X, p.Y+dy, p.Z)
}

func (p *Point3D) AddToZ(dz int) *Point3D {
	return NewPoint3D(p.X, p.Y, p.Z+dz)
}

func (p *Point3D) RotatePoint3DAroundX(pc *Point3D, theta float64) *Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X
	y := p.Y - pc.Y
	z := p.Z - pc.Z

	yr := int(float64(y)*cosTheta - float64(z)*sinTheta)
	zr := int(float64(y)*sinTheta + float64(z)*cosTheta)

	pr := NewPoint3D(x, yr+pc.Y, zr+pc.Z)
	return pr
}

func (p *Point3D) RotatePoint3DAroundY(pc *Point3D, theta float64) *Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y
	z := p.Z - pc.Z

	xr := int(float64(x)*cosTheta + float64(z)*sinTheta)
	zr := int(-float64(x)*sinTheta + float64(z)*cosTheta)

	pr := NewPoint3D(xr+pc.X, y, zr+pc.Z)
	return pr
}

func (p *Point3D) RotatePoint3DAroundZ(pc *Point3D, theta float64) *Point3D {
	sinTheta, cosTheta := math.Sincos(theta)

	x := p.X - pc.X
	y := p.Y - pc.Y
	z := p.Z

	xr := int(float64(x)*cosTheta - float64(y)*sinTheta)
	yr := int(float64(x)*sinTheta + float64(y)*cosTheta)

	pr := NewPoint3D(xr+pc.X, yr+pc.Y, z)
	return pr
}

func (p *Point3D) RotatePoint3DAroundXAndY(pc *Point3D, theta float64) *Point3D {
	return p.RotatePoint3DAroundX(pc, theta).RotatePoint3DAroundY(pc, theta)
}
