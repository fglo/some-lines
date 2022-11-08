package camera

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/point"
)

type Camera struct {
	Position    point.Point3D
	Orientation point.Orientation

	Vw          float64
	Vh          float64
	FocalLength float64
	Aspect      float64

	NearClippingPlane float64
	FarClippingPlane  float64

	Screen Screen
}

type Screen struct {
	TopLeft     point.Point2D
	TopRight    point.Point2D
	BottomLeft  point.Point2D
	BottomRight point.Point2D
}

func New(position point.Point3D, orientation point.Orientation) Camera {
	aspect := 16.0 / 9.0
	vh := 3.0
	vw := vh * aspect
	c := Camera{
		Position:          position,
		Orientation:       orientation,
		Vw:                vw,
		Vh:                vh,
		Aspect:            aspect,
		NearClippingPlane: 1,
		FarClippingPlane:  1000,
	}
	c.SetFoV(46)

	right := (c.Vw * 0.5 / c.FocalLength) * c.NearClippingPlane
	left := -right

	top := right * c.Aspect
	bottom := -top

	c.Screen = Screen{
		TopLeft:     point.NewPoint2D(int(left), int(top)),
		TopRight:    point.NewPoint2D(int(right), int(top)),
		BottomLeft:  point.NewPoint2D(int(left), int(bottom)),
		BottomRight: point.NewPoint2D(int(right), int(top)),
	}

	return c
}

func (c *Camera) SetFoV(fov float64) {
	fov = math.Mod(fov, 360)
	c.FocalLength = float64(c.Vw) * 0.5 / math.Tan(fov*0.5*math.Pi/180.0)
}

func (c *Camera) ProjectPoint(point3d point.Point3D) point.Point3DNdc {
	ac := point.NewPoint3D(point3d.X-c.Position.X, point3d.Y-c.Position.Y, point3d.Z-c.Position.Z)
	// ac := point3d
	d := ac.RotateAroundX(c.Orientation.X).RotateAroundY(c.Orientation.Y).RotateAroundZ(c.Orientation.Z)

	x := float64(d.X) * c.NearClippingPlane
	y := float64(d.Y) * c.NearClippingPlane
	if d.Z != 0 {
		x /= float64(-d.Z)
		y /= float64(-d.Z)
	}

	xNormalized := (float64(x) + float64(c.Vw)/2.0) / float64(c.Vw)
	yNormalized := (float64(y) + float64(c.Vh)/2.0) / float64(c.Vh)

	return point.NewPoint3DNdc(xNormalized, yNormalized, -d.Z)
}
