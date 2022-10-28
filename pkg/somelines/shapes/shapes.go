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
	dx := downRightCorner.X - leftUpCorner.X
	dy := downRightCorner.Y - leftUpCorner.Y
	dz := downRightCorner.Z - leftUpCorner.Z

	p1 := leftUpCorner
	p2 := leftUpCorner.MoveAlongX(dx)
	p3 := leftUpCorner.MoveAlongY(dy)
	p4 := leftUpCorner.MoveAlongXButPointer(dx).MoveAlongY(dy)
	p5 := leftUpCorner.MoveAlongZ(dz)
	p6 := leftUpCorner.MoveAlongXButPointer(dx).MoveAlongZ(dz)
	p7 := leftUpCorner.MoveAlongYButPointer(dy).MoveAlongZ(dz)
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

func NewPlane(middlePoint point.Point3D) Polygon3D {
	vertices := []point.Point3D{}
	edges := [][2]int{}

	vertices = append(vertices, middlePoint)
	vertex := point.NewPoint3D(middlePoint.X+200, middlePoint.Y, middlePoint.Z+200)
	vertices = append(vertices, vertex)
	vertex = point.NewPoint3D(middlePoint.X+200, middlePoint.Y, middlePoint.Z-200)
	vertices = append(vertices, vertex)
	vertex = point.NewPoint3D(middlePoint.X-200, middlePoint.Y, middlePoint.Z+200)
	vertices = append(vertices, vertex)
	vertex = point.NewPoint3D(middlePoint.X-200, middlePoint.Y, middlePoint.Z-200)
	vertices = append(vertices, vertex)
	for i := -9; i < 10; i++ {
		vertex := point.NewPoint3D(middlePoint.X+i*20, middlePoint.Y, middlePoint.Z+200)
		vertices = append(vertices, vertex)
		vertex = point.NewPoint3D(middlePoint.X+i*20, middlePoint.Y, middlePoint.Z-200)
		vertices = append(vertices, vertex)
		vertex = point.NewPoint3D(middlePoint.X+200, middlePoint.Y, middlePoint.Z+i*20)
		vertices = append(vertices, vertex)
		vertex = point.NewPoint3D(middlePoint.X-200, middlePoint.Y, middlePoint.Z+i*20)
		vertices = append(vertices, vertex)
	}

	for i := 0; i < len(vertices); i++ {
		v1 := vertices[i]
		for j := i + 1; j < len(vertices); j++ {
			v2 := vertices[j]
			if (v1.X == v2.X && v1.Z-v2.Z == 400) || (v1.Z == v2.Z && v1.X-v2.X == 400) {
				edges = append(edges, [2]int{i, j})
			}
		}
	}

	return NewPolygon3D(vertices, edges)
}

func NewTeapot() Polygon3D {
	fvertices := [][3]float64{
		{0.2000, 0.0000, 2.70000}, {0.2000, -0.1120, 2.70000},
		{0.1120, -0.2000, 2.70000}, {0.0000, -0.2000, 2.70000},
		{1.3375, 0.0000, 2.53125}, {1.3375, -0.7490, 2.53125},
		{0.7490, -1.3375, 2.53125}, {0.0000, -1.3375, 2.53125},
		{1.4375, 0.0000, 2.53125}, {1.4375, -0.8050, 2.53125},
		{0.8050, -1.4375, 2.53125}, {0.0000, -1.4375, 2.53125},
		{1.5000, 0.0000, 2.40000}, {1.5000, -0.8400, 2.40000},
		{0.8400, -1.5000, 2.40000}, {0.0000, -1.5000, 2.40000},
		{1.7500, 0.0000, 1.87500}, {1.7500, -0.9800, 1.87500},
		{0.9800, -1.7500, 1.87500}, {0.0000, -1.7500, 1.87500},
		{2.0000, 0.0000, 1.35000}, {2.0000, -1.1200, 1.35000},
		{1.1200, -2.0000, 1.35000}, {0.0000, -2.0000, 1.35000},
		{2.0000, 0.0000, 0.90000}, {2.0000, -1.1200, 0.90000},
		{1.1200, -2.0000, 0.90000}, {0.0000, -2.0000, 0.90000},
		{-2.0000, 0.0000, 0.90000}, {2.0000, 0.0000, 0.45000},
		{2.0000, -1.1200, 0.45000}, {1.1200, -2.0000, 0.45000},
		{0.0000, -2.0000, 0.45000}, {1.5000, 0.0000, 0.22500},
		{1.5000, -0.8400, 0.22500}, {0.8400, -1.5000, 0.22500},
		{0.0000, -1.5000, 0.22500}, {1.5000, 0.0000, 0.15000},
		{1.5000, -0.8400, 0.15000}, {0.8400, -1.5000, 0.15000},
		{0.0000, -1.5000, 0.15000}, {-1.6000, 0.0000, 2.02500},
		{-1.6000, -0.3000, 2.02500}, {-1.5000, -0.3000, 2.25000},
		{-1.5000, 0.0000, 2.25000}, {-2.3000, 0.0000, 2.02500},
		{-2.3000, -0.3000, 2.02500}, {-2.5000, -0.3000, 2.25000},
		{-2.5000, 0.0000, 2.25000}, {-2.7000, 0.0000, 2.02500},
		{-2.7000, -0.3000, 2.02500}, {-3.0000, -0.3000, 2.25000},
		{-3.0000, 0.0000, 2.25000}, {-2.7000, 0.0000, 1.80000},
		{-2.7000, -0.3000, 1.80000}, {-3.0000, -0.3000, 1.80000},
		{-3.0000, 0.0000, 1.80000}, {-2.7000, 0.0000, 1.57500},
		{-2.7000, -0.3000, 1.57500}, {-3.0000, -0.3000, 1.35000},
		{-3.0000, 0.0000, 1.35000}, {-2.5000, 0.0000, 1.12500},
		{-2.5000, -0.3000, 1.12500}, {-2.6500, -0.3000, 0.93750},
		{-2.6500, 0.0000, 0.93750}, {-2.0000, -0.3000, 0.90000},
		{-1.9000, -0.3000, 0.60000}, {-1.9000, 0.0000, 0.60000},
		{1.7000, 0.0000, 1.42500}, {1.7000, -0.6600, 1.42500},
		{1.7000, -0.6600, 0.60000}, {1.7000, 0.0000, 0.60000},
		{2.6000, 0.0000, 1.42500}, {2.6000, -0.6600, 1.42500},
		{3.1000, -0.6600, 0.82500}, {3.1000, 0.0000, 0.82500},
		{2.3000, 0.0000, 2.10000}, {2.3000, -0.2500, 2.10000},
		{2.4000, -0.2500, 2.02500}, {2.4000, 0.0000, 2.02500},
		{2.7000, 0.0000, 2.40000}, {2.7000, -0.2500, 2.40000},
		{3.3000, -0.2500, 2.40000}, {3.3000, 0.0000, 2.40000},
		{2.8000, 0.0000, 2.47500}, {2.8000, -0.2500, 2.47500},
		{3.5250, -0.2500, 2.49375}, {3.5250, 0.0000, 2.49375},
		{2.9000, 0.0000, 2.47500}, {2.9000, -0.1500, 2.47500},
		{3.4500, -0.1500, 2.51250}, {3.4500, 0.0000, 2.51250},
		{2.8000, 0.0000, 2.40000}, {2.8000, -0.1500, 2.40000},
		{3.2000, -0.1500, 2.40000}, {3.2000, 0.0000, 2.40000},
		{0.0000, 0.0000, 3.15000}, {0.8000, 0.0000, 3.15000},
		{0.8000, -0.4500, 3.15000}, {0.4500, -0.8000, 3.15000},
		{0.0000, -0.8000, 3.15000}, {0.0000, 0.0000, 2.85000},
		{1.4000, 0.0000, 2.40000}, {1.4000, -0.7840, 2.40000},
		{0.7840, -1.4000, 2.40000}, {0.0000, -1.4000, 2.40000},
		{0.4000, 0.0000, 2.55000}, {0.4000, -0.2240, 2.55000},
		{0.2240, -0.4000, 2.55000}, {0.0000, -0.4000, 2.55000},
		{1.3000, 0.0000, 2.55000}, {1.3000, -0.7280, 2.55000},
		{0.7280, -1.3000, 2.55000}, {0.0000, -1.3000, 2.55000},
		{1.3000, 0.0000, 2.40000}, {1.3000, -0.7280, 2.40000},
		{0.7280, -1.3000, 2.40000}, {0.0000, -1.3000, 2.40000}}
	vertices := []point.Point3D{}
	multiplier := 20.0
	for _, v := range fvertices {
		vertices = append(vertices, point.Point3D{int(v[0]*multiplier) + 75, -int(v[1]*multiplier) + 50, int(v[2] * multiplier)})
	}

	rim := []int{102, 103, 104, 105, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	body1 := []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	body2 := []int{24, 25, 26, 27, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	lid1 := []int{96, 96, 96, 96, 97, 98, 99, 100, 101, 101, 101, 101, 0, 1, 2, 3}
	lid2 := []int{0, 1, 2, 3, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117}
	handle1 := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56}
	handle2 := []int{53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 28, 65, 66, 67}
	spout1 := []int{68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83}
	spout2 := []int{80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95}

	edges := [][2]int{}
	edges = appendToEdges(edges, rim)
	edges = appendToEdges(edges, body1)
	edges = appendToEdges(edges, body2)
	edges = appendToEdges(edges, lid1)
	edges = appendToEdges(edges, lid2)
	edges = appendToEdges(edges, handle1)
	edges = appendToEdges(edges, handle2)
	edges = appendToEdges(edges, spout1)
	edges = appendToEdges(edges, spout2)

	return NewPolygon3D(vertices, edges)
}

func appendToEdges(edges [][2]int, patch []int) [][2]int {
	var vp int = -1
	for _, v := range patch {
		if vp > -1 {
			edges = append(edges, [2]int{vp, v})
		}
		vp = v
	}
	return edges
}
