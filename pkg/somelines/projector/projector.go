package projector

import (
	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

type Projector interface {
	ProjectPolygon(polygon shapes.Polygon3D)
	// ProjectLine(line shapes.Line3D)
	ProjectPoint(point point.Point3D)
}
