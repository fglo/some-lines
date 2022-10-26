package shapes

import "github.com/fglo/some-lines/pkg/somelines/point"

func NewTriangle(p1, p2, p3 point.Point2D) Polygon2D {
	vertices := []point.Point2D{p1, p2, p3}
	edges := [][2]int{{0, 1}, {1, 2}, {2, 0}}
	return NewPolygon2D(vertices, edges)
}

func NewQuadrangle(p1, p2, p3, p4 point.Point2D) Polygon2D {
	vertices := []point.Point2D{p1, p2, p3, p4}
	edges := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}
	return NewPolygon2D(vertices, edges)
}

func NewCube(leftUpCorner, downRightCorner point.Point3D) Polygon3D {
	d := downRightCorner.X - leftUpCorner.X
	p1 := leftUpCorner
	p2 := leftUpCorner.MoveAlongX(d)
	p3 := leftUpCorner.MoveAlongY(d)
	p4 := leftUpCorner.MoveAlongXButPointer(d).MoveAlongY(d)
	p5 := leftUpCorner.MoveAlongZ(d)
	p6 := leftUpCorner.MoveAlongXButPointer(d).MoveAlongZ(d)
	p7 := leftUpCorner.MoveAlongYButPointer(d).MoveAlongZ(d)
	p8 := downRightCorner
	vertices := []point.Point3D{p1, p2, p3, p4, p5, p6, p7, p8}
	edges := [][2]int{
		{0, 1}, {0, 2}, {0, 4},
		{3, 1}, {3, 2}, {3, 7},
		{5, 1}, {5, 4}, {5, 7},
		{6, 2}, {6, 4}, {6, 7},
	}
	return NewPolygon3D(vertices, edges)
}
