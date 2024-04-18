package scrshot

import (
	"image"

	"github.com/sanity32/lr"
)

type Screenshot struct {
	localresident *lr.Mgr
	Source        *image.RGBA
	LastErr       error
}

func NewScreenshotAll(l *lr.Mgr) (r Screenshot, err error) {
	r.localresident = l
	r.TakeAll()
	return r, r.LastErr
}

func NewScreenshotXYWH(l *lr.Mgr, x, y, w, h int) (r Screenshot, err error) {
	r.localresident = l
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
	// img, err := screenshot.CaptureScreen()
	img, err := s.localresident.Client().Screenshot()
	s.Save(img, err)
}

func (s *Screenshot) TakeXYWH(x, y, w, h int) error {
	rect := image.Rect(x, y, x+w, y+h)
	// img, err := screenshot.CaptureRect(rect)
	img, err := s.localresident.Client().ScreenshotRect(rect)
	if err != nil {
		return err
	}
	s.Save(img, err)
	return nil
}

func (s Screenshot) CropToXYWH(x, y, w, h int) Image {
	min := image.Pt(x, y)
	max := image.Pt(x+w, y+h)
	return NewImage(s.Source.SubImage(image.Rectangle{
		Min: min,
		Max: max,
	}))
}

func (s Screenshot) Out() Image {
	return NewImage(s.Source)
}
