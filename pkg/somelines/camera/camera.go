package camera

import (
	"math"

	"github.com/fglo/some-lines/pkg/somelines/point"
)

type Camera struct {
	Position    point.Point3D
	Orientation point.Orientation

	Vw          int
	Vh          int
	FocalLength int
}

func New(position point.Point3D, orientation point.Orientation) Camera {
	c := Camera{
		Position:    position,
		Orientation: orientation,
		Vw:          100,
		Vh:          100,
		FocalLength: 100,
	}
	return c
}

func (c *Camera) SetFoV(fov float64) {
	fov = math.Mod(fov, 360)
	c.Vw = int(2.0 * float64(c.FocalLength) * math.Tan(fov*math.Pi/180.0) * 0.5)
}
