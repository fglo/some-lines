package projector

import (
	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

type Projector interface {
	ProjectPolygon(p shapes.Polygon3D, c *camera.Camera) shapes.ProjectedPolygon3D
}
