package point

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
