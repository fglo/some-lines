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
