package scrshot

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
)

type Image struct {
	Content image.Image
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

func NewImage(input image.Image) Image {
	return Image{
		Content: input,
	}
}
