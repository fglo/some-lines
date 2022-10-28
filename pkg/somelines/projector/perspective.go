package projector

import (
	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

type PerspectiveProjector struct {
	Camera camera.Camera
}

func NewPerspectiveProjector() PerspectiveProjector {
	pp := PerspectiveProjector{}
	return pp
}

func (pp *PerspectiveProjector) SetCamera(c camera.Camera) {
	pp.Camera = c
}

func (pp *PerspectiveProjector) ProjectPolygon(p shapes.Polygon3D) shapes.Polygon2D {
	vs2d := make([]point.Point2D, 0)
	for _, v := range p.Vertices {
		vs2d = append(vs2d, pp.ProjectPoint(v))
	}
	return shapes.NewPolygon2D(vs2d, p.Edges)
}

// func (pp *PerspectiveProjector) ProjectLine(l shapes.Line3D) shapes.Line {

// }

func (pp *PerspectiveProjector) ProjectPoint(p point.Point3D) point.Point2D {
	ac := point.NewPoint3D(p.X-pp.Camera.Position.X, p.Y-pp.Camera.Position.Y, p.Z-pp.Camera.Position.Z)
	d := ac.RotateAroundX(pp.Camera.Orientation.X).RotateAroundY(pp.Camera.Orientation.Y).RotateAroundZ(pp.Camera.Orientation.Z)

	x := d.X * pp.Camera.FocalLength
	y := d.Y * pp.Camera.FocalLength
	if d.Z != 0 {
		x /= d.Z
		y /= d.Z
	}

	x = x * 300 / pp.Camera.Vw
	y = y * 300 / pp.Camera.Vh

	return point.NewPoint2D(x+pp.Camera.Position.X, y+pp.Camera.Position.Y)
	// return point.NewPoint2D(x+pp.Camera.Position.X, y+pp.Camera.Position.Y)
}
