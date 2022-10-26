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
		rotatedVertices = append(rotatedVertices, v3d.RotateAroundX(pc.To3D(), theta))
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
		rotatedVertices = append(rotatedVertices, v3d.RotateAroundY(pc.To3D(), theta))
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
		rotatedVertices = append(rotatedVertices, v3d.RotateAroundZ(pc.To3D(), theta))
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

// RotateAroundX rotates a polygon around X axis
func (p *Polygon3D) RotateAroundX(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundX(pc, theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

// RotateAroundY rotates a polygon around Y axis
func (p *Polygon3D) RotateAroundY(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundY(pc, theta))
	}

	np := NewPolygon3D(rotatedVertices, p.Edges)
	return &np
}

// RotateAroundZ rotates a polygon around Z axis
func (p *Polygon3D) RotateAroundZ(theta float64) *Polygon3D {
	pc := p.CalculateCenterPoint()

	rotatedVertices := make([]point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundZ(pc, theta))
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
