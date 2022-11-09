package camera

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/point"
)

type Camera struct {
	Position    point.Point3Df
	Orientation point.Orientation

	Vw          float64
	Vh          float64
	FocalLength float64
	Aspect      float64

	NearClippingPlane float64
	FarClippingPlane  float64

	Screen Screen

	CameraToWorld [4][4]float64
}

type Screen struct {
	TopLeft     point.Point2D
	TopRight    point.Point2D
	BottomLeft  point.Point2D
	BottomRight point.Point2D
}

func New(position point.Point3Df, orientation point.Orientation) Camera {
	aspect := 16.0 / 9.0
	// aspect := 0.980 / 0.735
	vh := 0.9
	vw := vh * aspect

	sinX, cosX := math.Sincos(orientation.X)
	sinY, cosY := math.Sincos(orientation.Y)
	sinZ, cosZ := math.Sincos(orientation.Z)

	cameraToWorld := [4][4]float64{
		{cosY * cosZ, -cosY * sinZ, sinY, position.X},
		{cosX*sinZ + sinX*sinY*cosZ, cosX*cosZ - sinX*sinY*sinZ, -sinX * cosY, position.Y},
		{sinX*sinZ - cosX*sinY*cosZ, sinX*cosZ - cosX*sinY*sinZ, cosX * cosY, position.Z},
		{0, 0, 0, 1},
	}

	c := Camera{
		Position:          position,
		Orientation:       orientation,
		Vw:                vw,
		Vh:                vh,
		Aspect:            aspect,
		NearClippingPlane: 0.1,
		FarClippingPlane:  1000,
		// FocalLength:       35,
		CameraToWorld: cameraToWorld,
	}
	c.SetFoV(46)
	// c.SetFoV(53)

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
