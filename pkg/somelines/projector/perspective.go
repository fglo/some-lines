package projector

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

type PerspectiveProjector struct {
	Cw int
	Ch int
}

func NewPerspectiveProjector(cw, ch int) Projector {
	pp := PerspectiveProjector{
		Cw: cw,
		Ch: ch,
	}
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
	pointNDC := c.ProjectPoint(point3d)

	x := int(math.Floor(pointNDC.X * float64(pp.Cw)))
	y := int(math.Floor((1 - pointNDC.Y) * float64(pp.Ch)))

	return point.NewProjectedPoint3D(x, y, float64(pointNDC.Z))
}

// func (pp *PerspectiveProjector) projectPoint(point3d point.Point3D, c *camera.Camera) point.ProjectedPoint3D {
// 	ac := point.NewPoint3D(point3d.X-c.Position.X, point3d.Y-c.Position.Y, point3d.Z-c.Position.Z)
// 	// ac := point3d
// 	d := ac.RotateAroundX(c.Orientation.X).RotateAroundY(c.Orientation.Y).RotateAroundZ(c.Orientation.Z)

// 	x := d.X * c.NearClippingPlane
// 	y := d.Y * c.NearClippingPlane
// 	if d.Z != 0 {
// 		x /= -d.Z
// 		y /= -d.Z
// 	}

// 	xNormalized := (float64(x) + float64(c.Vw)/2.0) / float64(c.Vw)
// 	yNormalized := (float64(y) + float64(c.Vh)/2.0) / float64(c.Vh)

// 	x = int(math.Floor(xNormalized * float64(pp.Cw)))
// 	y = int(math.Floor((1 - yNormalized) * float64(pp.Ch)))

// 	// x = x * pp.Cw / c.Vw
// 	// y = y * pp.Ch / c.Vh

// 	// return point.NewProjectedPoint3D(x+c.Position.X, y+c.Position.Y, float64(d.Z))
// 	return point.NewProjectedPoint3D(x, y, float64(d.Z))
// }
