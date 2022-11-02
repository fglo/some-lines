package renderer

import (
	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/point"
	"github.com/fglo/some-lines/pkg/somelines/projector"
	"github.com/fglo/some-lines/pkg/somelines/scene"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
	"github.com/fglo/some-lines/pkg/somelines/utils"
)

type Renderer struct {
	Projector projector.Projector
}

func New(projector projector.Projector) Renderer {
	r := Renderer{
		Projector: projector,
	}

	return r
}

func (r *Renderer) RenderScene(s scene.Scene3D, screenWidth, screenHeight int, pixels []byte) {
	for _, p := range s.Polygons3D {
		projected := r.Projector.ProjectPolygon(p, s.ActiveCamera)
		r.RenderPolygon(&projected, *s.ActiveCamera, screenWidth, screenHeight, pixels)
	}
}

func (r *Renderer) RenderPolygon(polygon *shapes.ProjectedPolygon3D, c camera.Camera, screenWidth, screenHeight int, pixels []byte) {
	for _, edge := range polygon.Edges {
		r.RenderLine(polygon.Vertices[edge[0]], polygon.Vertices[edge[1]], c, screenWidth, screenHeight, pixels)
	}
}

func (r *Renderer) RenderLine(startPoint, endPoint point.ProjectedPoint3D, c camera.Camera, screenWidth, screenHeight int, pixels []byte) {
	for _, p := range shapes.PlotLine(startPoint, endPoint, c) {
		utils.Color3DPixel(p.X, p.Y, p.D, screenWidth, screenHeight, pixels)
	}
}
