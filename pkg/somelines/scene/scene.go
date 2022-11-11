package scene

import (
	"fmt"

	"github.com/fglo/some-lines/pkg/somelines/camera"
	"github.com/fglo/some-lines/pkg/somelines/shapes"
)

type Scene3D struct {
	Polygons3D  []shapes.Polygon3D
	Polygons3Df []shapes.Polygon3Df

	Cameras      map[string]*camera.Camera
	ActiveCamera *camera.Camera
}

func New() Scene3D {
	return Scene3D{
		Polygons3D: make([]shapes.Polygon3D, 0),
		Cameras:    make(map[string]*camera.Camera),
	}
}

func (s *Scene3D) AddPolygon3D(p shapes.Polygon3D) int {
	s.Polygons3D = append(s.Polygons3D, p)
	return len(s.Polygons3D) - 1
}

func (s *Scene3D) GetPolygon3D(id int) shapes.Polygon3D {
	if id > 0 && id < len(s.Polygons3D) {
		return s.Polygons3D[id]
	}
	return s.Polygons3D[0]
}

func (s *Scene3D) AddPolygon3Df(p shapes.Polygon3Df) int {
	s.Polygons3Df = append(s.Polygons3Df, p)
	return len(s.Polygons3Df) - 1
}

func (s *Scene3D) AddCamera(lbl string, c *camera.Camera) {
	s.Cameras[lbl] = c
	if len(s.Cameras) == 1 {
		s.ActiveCamera = c
	}
}

func (s *Scene3D) SetActiveCamera(lbl string) error {
	if camera, found := s.Cameras[lbl]; found {
		s.ActiveCamera = camera
		return nil
	}
	return fmt.Errorf("Camera not found.")
}
