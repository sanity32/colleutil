package scrshot

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

type Image struct {
	Content image.Image
}

func NewImage(input image.Image) Image {
	return Image{Content: input}
}

func (i Image) ToPng() []byte {
	var b bytes.Buffer
	png.Encode(&b, i.Content)
	return b.Bytes()
}

func (i Image) ToJpeg() []byte {
	var b bytes.Buffer
	jpeg.Encode(&b, i.Content, nil)
	return b.Bytes()
}

func (i Image) ToBase64Png() B64Result {
	return NewB64ResultFromBytes(i.ToPng())
}

func (i Image) ToBase64Jpg() B64Result {
	return NewB64ResultFromBytes(i.ToJpeg())
}

func (i Image) SaveJpeg(filename string, q ...int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var opts *jpeg.Options = nil
	if len(q) > 0 {
		opts = &jpeg.Options{Quality: q[0]}
	}
	return jpeg.Encode(f, i.Content, opts)
}
