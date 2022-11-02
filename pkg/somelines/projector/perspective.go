package projector

import (
	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

type PerspectiveProjector struct {
}

func NewPerspectiveProjector() Projector {
	pp := PerspectiveProjector{}
	return &pp
}

func (pp *PerspectiveProjector) ProjectPolygon(polygon shapes.Polygon3D, c *camera.Camera) shapes.ProjectedPolygon3D {
	vs := make([]point.ProjectedPoint3D, 0)
	for _, v := range polygon.Vertices {
		vs = append(vs, pp.projectPoint(v, c))
	}
	return shapes.NewProjectedPolygon3D(vs, polygon.Edges)
}

func (pp *PerspectiveProjector) projectPoint(point3d point.Point3D, c *camera.Camera) point.ProjectedPoint3D {
	ac := point.NewPoint3D(point3d.X-c.Position.X, point3d.Y-c.Position.Y, point3d.Z-c.Position.Z)
	d := ac.RotateAroundX(c.Orientation.X).RotateAroundY(c.Orientation.Y).RotateAroundZ(c.Orientation.Z)

	x := d.X * c.FocalLength
	y := d.Y * c.FocalLength
	if d.Z != 0 {
		x /= d.Z
		y /= d.Z
	}

	x = x * 300 / c.Vw
	y = y * 300 / c.Vh

	return point.NewProjectedPoint3D(x+c.Position.X, y+c.Position.Y, float64(d.Z))
}
