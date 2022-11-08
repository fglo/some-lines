package projector

import (
	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

// TODO: real orthographic projection
// Possibility to do isometric projection

type OrthogonalProjector struct {
	Cw int
	Ch int
}

func NewOrthogonalProjector(cw, ch int) Projector {
	op := OrthogonalProjector{
		Cw: cw,
		Ch: ch,
	}
	return &op
}

func (op *OrthogonalProjector) ProjectPolygon(polygon shapes.Polygon3D, c *camera.Camera) shapes.ProjectedPolygon3D {
	pc := polygon.CalculateFlatCenterPoint()
	vs := make([]point.ProjectedPoint3D, 0)
	for _, v := range polygon.Vertices {
		vs = append(vs, op.projectPoint(v, pc, c))
	}
	return shapes.NewProjectedPolygon3D(vs, polygon.Edges)
}

func (op *OrthogonalProjector) projectPoint(point3d point.Point3D, pc point.Point2D, c *camera.Camera) point.ProjectedPoint3D {
	x := point3d.X - pc.X
	y := point3d.Y - pc.Y
	z := point3d.Z - c.Position.Z
	fl := c.FocalLength

	xp := int(float64(x)*fl/float64(z) + fl)
	yp := int(float64(y)*fl/float64(z) + fl)

	// xp = xp * op.Cw / c.Vw
	// yp = yp * op.Ch / c.Vh

	return point.NewProjectedPoint3D(xp+pc.X, yp+pc.Y, float64(point3d.Z))
}
