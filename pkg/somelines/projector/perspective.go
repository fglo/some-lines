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

func (pp *PerspectiveProjector) ProjectPolygon3Df(polygon shapes.Polygon3Df, c *camera.Camera) shapes.ProjectedPolygon3D {
	vs := make([]point.ProjectedPoint3D, 0)
	for _, v := range polygon.Vertices {
		vs = append(vs, pp.projectPoint3Df(v, c))
	}
	return shapes.NewProjectedPolygon3D(vs, polygon.Edges)
}

func (pp *PerspectiveProjector) projectPoint(point3d point.Point3D, c *camera.Camera) point.ProjectedPoint3D {
	// TODO: implement world_to_camera transformation matrix
	ac := point.NewPoint3D(point3d.X-int(c.Position.X), point3d.Y-int(c.Position.Y), point3d.Z-int(c.Position.Z))
	d := ac.RotateAroundX(c.Orientation.X).RotateAroundY(c.Orientation.Y).RotateAroundZ(c.Orientation.Z)

	x := float64(d.X) * c.NearClippingPlane
	y := float64(d.Y) * c.NearClippingPlane
	if d.Z != 0 {
		x /= float64(-d.Z)
		y /= float64(-d.Z)
	}

	xNdc := (float64(x) + float64(c.Vw)/2.0) / float64(c.Vw)
	yNdc := (float64(y) + float64(c.Vh)/2.0) / float64(c.Vh)

	xRastered := int(math.Floor(xNdc * float64(pp.Cw)))
	yRastered := int(math.Floor((1 - yNdc) * float64(pp.Ch)))

	return point.NewProjectedPoint3D(xRastered, yRastered, float64(d.Z))
}

func (pp *PerspectiveProjector) projectPoint3Df(point3d point.Point3Df, c *camera.Camera) point.ProjectedPoint3D {
	// TODO: implement world_to_camera transformation matrix
	// ac := point.NewPoint3Df(point3d.X-c.Position.X, point3d.Y-c.Position.Y, point3d.Z-c.Position.Z)
	// d := ac.RotateAroundX(c.Orientation.X).RotateAroundY(c.Orientation.Y).RotateAroundZ(c.Orientation.Z)

	d := point.NewPoint3Df(
		point3d.X*c.CameraToWorld[0][0]+point3d.Y*c.CameraToWorld[0][1]+point3d.Z*c.CameraToWorld[0][2]+c.CameraToWorld[0][3],
		point3d.X*c.CameraToWorld[1][0]+point3d.Y*c.CameraToWorld[1][1]+point3d.Z*c.CameraToWorld[1][2]+c.CameraToWorld[1][3],
		point3d.X*c.CameraToWorld[2][0]+point3d.Y*c.CameraToWorld[2][1]+point3d.Z*c.CameraToWorld[2][2]+c.CameraToWorld[2][3],
	)

	x := float64(d.X) * c.NearClippingPlane
	y := float64(d.Y) * c.NearClippingPlane
	if d.Z != 0 {
		x /= float64(-d.Z)
		y /= float64(-d.Z)
	}

	xNdc := (x + c.Vw*0.5) / c.Vw
	yNdc := (y + c.Vh*0.5) / c.Vh

	xRastered := int(math.Floor(xNdc * float64(pp.Cw)))
	yRastered := int(math.Floor((1 - yNdc) * float64(pp.Ch)))

	return point.NewProjectedPoint3D(xRastered, yRastered, float64(-d.Z))
}
