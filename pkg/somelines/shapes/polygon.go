package shapes

import "github.com/fglo/some-lines/pkg/somelines/point"

type Polygon2D struct {
	Vertices []point.Point2D
	Edges    [][2]int
}

func NewPolygon2D(vertices []point.Point2D, edges [][2]int) Polygon2D {
	p := Polygon2D{Vertices: vertices, Edges: edges}
	return p
}

func (p *Polygon2D) MoveAlongX(dx int) Polygon2D {
	p2 := *p
	vertices := make([]point.Point2D, 0)
	for _, v := range p2.Vertices {
		v.X += dx
		vertices = append(vertices, v)
	}
	p2.Vertices = vertices
	return p2
}

func (p *Polygon2D) MoveAlongY(dy int) Polygon2D {
	p2 := *p
	vertices := make([]point.Point2D, 0)
	for _, v := range p2.Vertices {
		v.Y += dy
		vertices = append(vertices, v)
	}
	p2.Vertices = vertices
	return p2
}

func (p *Polygon2D) MoveAlongXButPointer(dx int) *Polygon2D {
	p2 := p.MoveAlongX(dx)
	return &p2
}

func (p *Polygon2D) MoveAlongYButPointer(dy int) *Polygon2D {
	p2 := p.MoveAlongY(dy)
	return &p2
}

// Rotate rotates a polygon
func (p *Polygon2D) Rotate(theta float64) *Polygon2D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point2D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.Rotate(pc, theta))
	}

	np := NewPolygon2D(rotatedVertices, p.Edges)
	return &np
}

// RotateAroundX rotates a polygon around X axis
func (p *Polygon2D) RotateAroundX(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		v3d := v.To3D()
		rotatedVertices = append(rotatedVertices, v3d.RotateAroundXRelativeToPoint(pc.To3D(), theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

// RotateAroundY rotates a polygon around Y axis
func (p *Polygon2D) RotateAroundY(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		v3d := v.To3D()
		rotatedVertices = append(rotatedVertices, v3d.RotateAroundYRelativeToPoint(pc.To3D(), theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

// RotateAroundZ rotates a polygon around Z axis
func (p *Polygon2D) RotateAroundZ(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		v3d := v.To3D()
		rotatedVertices = append(rotatedVertices, v3d.RotateAroundZRelativeToPoint(pc.To3D(), theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

func (p *Polygon2D) CalculateCenterPoint() point.Point2D {
	var sumX, sumY int
	lv := len(p.Vertices)
	for _, v := range p.Vertices {
		sumX += v.X
		sumY += v.Y
	}

	pc := point.NewPoint2D(sumX/lv, sumY/lv)
	return pc
}

type Polygon3D struct {
	Vertices []point.Point3D
	Edges    [][2]int
}

func NewPolygon3D(vertices []point.Point3D, edges [][2]int) Polygon3D {
	p := Polygon3D{Vertices: vertices, Edges: edges}
	return p
}

func (p *Polygon3D) Project(cameraPosition point.Point3D, cameraOrientation point.Orientation) []Line {
	lines := make([]Line, 0)
	for _, edge := range p.Edges {
		v1 := p.Vertices[edge[0]]
		v1p := v1.Project(cameraPosition, cameraOrientation)
		// v1.X = v1p.X
		// v1.Y = v1p.Y

		v2 := p.Vertices[edge[1]]
		v2p := v2.Project(cameraPosition, cameraOrientation)
		// v2.X = v2p.X
		// v2.Y = v2p.Y

		line := NewLine(v1p, v2p)
		lines = append(lines, line)
	}
	return lines
}

func (p *Polygon3D) MoveAlongX(dx int) Polygon3D {
	p2 := *p
	vertices := make([]point.Point3D, 0)
	for _, v := range p2.Vertices {
		v.X += dx
		vertices = append(vertices, v)
	}
	p2.Vertices = vertices
	return p2
}

func (p *Polygon3D) MoveAlongY(dy int) Polygon3D {
	p2 := *p
	vertices := make([]point.Point3D, 0)
	for _, v := range p2.Vertices {
		v.Y += dy
		vertices = append(vertices, v)
	}
	p2.Vertices = vertices
	return p2
}

func (p *Polygon3D) MoveAlongZ(dz int) Polygon3D {
	p2 := *p
	vertices := make([]point.Point3D, 0)
	for _, v := range p2.Vertices {
		v.Z += dz
		vertices = append(vertices, v)
	}
	p2.Vertices = vertices
	return p2
}

func (p *Polygon3D) MoveAlongXButPointer(dx int) *Polygon3D {
	p2 := p.MoveAlongX(dx)
	return &p2
}

func (p *Polygon3D) MoveAlongYButPointer(dy int) *Polygon3D {
	p2 := p.MoveAlongY(dy)
	return &p2
}

func (p *Polygon3D) MoveAlongZButPointer(dz int) *Polygon3D {
	p2 := p.MoveAlongZ(dz)
	return &p2
}

// RotateAroundX rotates a polygon around X axis
func (p *Polygon3D) RotateAroundX(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundXRelativeToPoint(pc, theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

// RotateAroundY rotates a polygon around Y axis
func (p *Polygon3D) RotateAroundY(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundYRelativeToPoint(pc, theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

// RotateAroundZ rotates a polygon around Z axis
func (p *Polygon3D) RotateAroundZ(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundZRelativeToPoint(pc, theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

func (p *Polygon3D) CalculateCenterPoint() point.Point3D {
	var sumX, sumY, sumZ int
	lv := len(p.Vertices)
	for _, v := range p.Vertices {
		sumX += v.X
		sumY += v.Y
		sumZ += v.Z
	}

	pc := point.NewPoint3D(sumX/lv, sumY/lv, sumZ/lv)
	return pc
}

func (p *Polygon3D) CalculateFlatCenterPoint() point.Point2D {
	var sumX, sumY int
	lv := len(p.Vertices)
	for _, v := range p.Vertices {
		sumX += v.X
		sumY += v.Y
	}

	pc := point.NewPoint2D(sumX/lv, sumY/lv)
	return pc
}
