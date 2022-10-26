package shapes

import "github.com/fglo/some-lines/pkg/somelines/point"

type Polygon2D struct {
	Vertices []*point.Point2D
	Edges    [][2]int
}

func NewPolygon2D(vertices []*point.Point2D, edges [][2]int) *Polygon2D {
	p := new(Polygon2D)

	p.Vertices = vertices
	p.Edges = edges

	return p
}

// Rotate rotates a polygon
func (p *Polygon2D) Rotate(theta float64) *Polygon2D {
	pc := p.calculateCenterPoint()

	rotatedVertices := make([]*point.Point2D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.Rotate(pc, theta))
	}

	return NewPolygon2D(rotatedVertices, p.Edges)
}

// RotateAroundX rotates a polygon around X axis
func (p *Polygon2D) RotateAroundX(theta float64) *Polygon3D {
	pc := p.calculateCenterPoint()

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.To3D().RotateAroundX(pc.To3D(), theta))
	}

	return NewPolygon3D(rotatedVertices, p.Edges)
}

// RotateAroundY rotates a polygon around Y axis
func (p *Polygon2D) RotateAroundY(theta float64) *Polygon3D {
	pc := p.calculateCenterPoint()

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.To3D().RotateAroundY(pc.To3D(), theta))
	}

	return NewPolygon3D(rotatedVertices, p.Edges)
}

// RotateAroundZ rotates a polygon around Z axis
func (p *Polygon2D) RotateAroundZ(theta float64) *Polygon3D {
	pc := p.calculateCenterPoint()

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.To3D().RotateAroundZ(pc.To3D(), theta))
	}

	return NewPolygon3D(rotatedVertices, p.Edges)
}

type Polygon3D struct {
	Vertices []*point.Point3D
	Edges    [][2]int
}

func NewPolygon3D(vertices []*point.Point3D, edges [][2]int) *Polygon3D {
	p := new(Polygon3D)

	p.Vertices = vertices
	p.Edges = edges

	return p
}

// RotateAroundX rotates a polygon around X axis
func (p *Polygon3D) RotateAroundX(theta float64) *Polygon3D {
	pc := p.calculateCenterPoint()

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundX(pc, theta))
	}

	return NewPolygon3D(rotatedVertices, p.Edges)
}

// RotateAroundY rotates a polygon around Y axis
func (p *Polygon3D) RotateAroundY(theta float64) *Polygon3D {
	pc := p.calculateCenterPoint()

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundY(pc, theta))
	}

	return NewPolygon3D(rotatedVertices, p.Edges)
}

// RotateAroundZ rotates a polygon around Z axis
func (p *Polygon3D) RotateAroundZ(theta float64) *Polygon3D {
	pc := p.calculateCenterPoint()

	rotatedVertices := make([]*point.Point3D, 0)
	for _, v := range p.Vertices {
		rotatedVertices = append(rotatedVertices, v.RotateAroundZ(pc, theta))
	}

	return NewPolygon3D(rotatedVertices, p.Edges)
}

func (p *Polygon2D) calculateCenterPoint() *point.Point2D {
	var sumX, sumY int
	lv := len(p.Vertices)
	for _, v := range p.Vertices {
		sumX += v.X
		sumY += v.Y
	}

	pc := point.NewPoint2D(sumX/lv, sumY/lv)
	return pc
}

func (p *Polygon3D) calculateCenterPoint() *point.Point3D {
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
