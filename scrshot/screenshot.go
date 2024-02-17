package scrshot

import (
	"fmt"
	"image"

	"github.com/vova616/screenshot"
)

type Screenshot struct {
	Source  *image.RGBA
	LastErr error
}

func NewScreenshotAll() (r Screenshot, err error) {
	r.TakeAll()
	return r, r.LastErr
}

func NewScreenshotXYWH(x, y, w, h int) (r Screenshot, err error) {
	err = r.TakeXYWH(x, y, w, h)
	return
}

func (s Screenshot) IsEmpty() bool {
	return s.Source == nil || s.LastErr != nil
}

func (s *Screenshot) Save(img *image.RGBA, err error) {
	if err != nil {
		s.LastErr = err
		return
	}
	s.Source = img
}

func (s *Screenshot) TakeAll() {
	img, err := screenshot.CaptureScreen()
	s.Save(img, err)
}

func (s *Screenshot) TakeXYWH(x, y, w, h int) error {
	rect := image.Rect(x, y, x+w, y+h)
	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		return err
	}
	s.Save(img, err)
	return nil
}

func (s Screenshot) CropToXYWH(x, y, w, h int) Image {
	min := image.Pt(x, y)
	max := image.Pt(x+w, y+h)
	fmt.Println("min", min, "max", max)
	return NewImage(s.Source.SubImage(image.Rectangle{
		Min: min,
		Max: max,
	}))
}

func (s Screenshot) Out() Image {
	return NewImage(s.Source)
}
