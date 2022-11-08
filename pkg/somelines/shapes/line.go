package shapes

import (
	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/point"
)

type Line struct {
	P1 point.Point2D
	P2 point.Point2D
}

func NewLine(p1, p2 point.Point2D) Line {
	return Line{P1: p1, P2: p2}
}

func (l *Line) PlotLine() []point.Point2D {
	var (
		dx, sx      int
		dy, sy      int
		coordinates []point.Point2D
	)

	if l.P1.X < l.P2.X {
		sx = 1
		dx = l.P2.X - l.P1.X
	} else {
		sx = -1
		dx = l.P1.X - l.P2.X
	}

	if l.P1.Y < l.P2.Y {
		sy = 1
		dy = l.P2.Y - l.P1.Y
	} else {
		sy = -1
		dy = l.P1.Y - l.P2.Y
	}

	p0 := l.P1.Clone()

	switch {
	case dx > dy:
		ai := (dy - dx) * 2
		bi := dy * 2
		d := bi - dx
		coordinates = append(coordinates, p0)
		for p0.X != l.P2.X {
			if d >= 0 {
				p0.Y += sy
				d += ai
			} else {
				d += bi
			}
			p0.X += sx
			coordinates = append(coordinates, p0)
		}
	case dx <= dy:
		ai := (dx - dy) * 2
		bi := dx * 2
		d := bi - dy
		coordinates = append(coordinates, p0)
		for p0.Y != l.P2.Y {
			if d >= 0 {
				p0.X += sx
				d += ai
			} else {
				d += bi
			}
			p0.Y += sy
			coordinates = append(coordinates, p0)
		}
	}

	return coordinates
}

type Line3D struct {
	P1 point.Point3D
	P2 point.Point3D
}

func NewLine3D(p1, p2 point.Point3D) Line3D {
	return Line3D{P1: p1, P2: p2}
}

func PlotLine(startPoint, endPoint point.ProjectedPoint3D, c camera.Camera) []point.ProjectedPoint3D {
	var (
		dx, sx        int
		dy, sy        int
		dz            int
		depth, sdepth float64
		ps, pe        point.Point2D
		coordinates   []point.ProjectedPoint3D
	)

	if startPoint.D < endPoint.D {
		ps = point.NewPoint2D(startPoint.X, startPoint.Y)
		pe = point.NewPoint2D(endPoint.X, endPoint.Y)
		dz = int(endPoint.D - startPoint.D)
	} else {
		ps = point.NewPoint2D(endPoint.X, endPoint.Y)
		pe = point.NewPoint2D(startPoint.X, startPoint.Y)
		dz = int(startPoint.D - endPoint.D)
	}

	if ps.X < pe.X {
		sx = 1
		dx = pe.X - ps.X
	} else {
		sx = -1
		dx = ps.X - pe.X
	}

	if ps.Y < pe.Y {
		sy = 1
		dy = pe.Y - ps.Y
	} else {
		sy = -1
		dy = ps.Y - pe.Y
	}

	switch {
	case dx > dy:
		ai := (dy - dx) * 2
		bi := dy * 2
		d := bi - dx
		depth = startPoint.D
		if startPoint.D == endPoint.D {
			sdepth = 0
		} else if dy == 0 {
			sdepth = float64(dz)
		} else {
			sdepth = float64(dz) / float64(dy)
		}

		coordinates = append(coordinates, point.NewProjectedPoint3D(ps.X, ps.Y, depth))
		for ps.X != pe.X {
			if d >= 0 {
				ps.Y += sy
				d += ai
			} else {
				d += bi
			}
			ps.X += sx
			if (sdepth > 0 && depth <= endPoint.D) || (sdepth < 0 && depth >= endPoint.D) {
				depth += sdepth
			}
			coordinates = append(coordinates, point.NewProjectedPoint3D(ps.X, ps.Y, depth))
		}
	case dx <= dy:
		ai := (dx - dy) * 2
		bi := dx * 2
		d := bi - dy
		depth = startPoint.D
		if startPoint.D == endPoint.D {
			sdepth = 0
		} else if dx == 0 {
			sdepth = float64(dz)
		} else {
			sdepth = float64(dz) / float64(dx)
		}
		coordinates = append(coordinates, point.NewProjectedPoint3D(ps.X, ps.Y, depth))
		for ps.Y != pe.Y {
			if d >= 0 {
				ps.X += sx
				d += ai
			} else {
				d += bi
			}
			ps.Y += sy
			if (sdepth > 0 && depth <= endPoint.D) || (sdepth < 0 && depth >= endPoint.D) {
				depth += sdepth
			}
			coordinates = append(coordinates, point.NewProjectedPoint3D(ps.X, ps.Y, depth))
		}
	}

	return coordinates
}
