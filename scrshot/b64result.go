package scrshot

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

const (
	PREFIX_B64_PNG = "data:image/png;base64,"
	PREFIX_B64_JPG = "data:image/jpeg;base64,"
)

type B64Result string

func NewB64ResultFromBytes(input []byte) B64Result {
	str := base64.StdEncoding.EncodeToString(input)
	return NewB64Result(str)
}

func NewB64Result(input string) B64Result {
	return B64Result(input)
}

func (b B64Result) String() string {
	return string(b)
}

func (b B64Result) Bytes() []byte {
	return []byte(b)
}

func (b B64Result) HasPrefix(p string) bool {
	return strings.HasPrefix(b.String(), p)
}

func (b B64Result) AddPrefix(p string) B64Result {
	if b.HasPrefix(p) {
		return b
	}
	return B64Result(p + b.String())
}

func (b B64Result) AddPngPrefix() B64Result {
	return b.AddPrefix(PREFIX_B64_PNG)
}

func (b B64Result) AddJpgPrefix() B64Result {
	return b.AddPrefix(PREFIX_B64_JPG)
}

func (b B64Result) RemovePrefix(p string) B64Result {
	if !b.HasPrefix(p) {
		return b
	}
	return B64Result(strings.TrimPrefix(b.String(), p))
}

func (b B64Result) RemovePngPrefix() B64Result {
	return b.RemovePrefix(PREFIX_B64_PNG)
}

func (b B64Result) RemoveJpgPrefix() B64Result {
	return b.RemovePrefix(PREFIX_B64_JPG)
}

func (b B64Result) RemovePrefixAll() B64Result {
	needle := ";base64,"
	if !strings.Contains(b.String(), needle) {
		return b
	}
	s := strings.Split(b.String(), needle)[1]
	return NewB64Result(s)
}

func (b B64Result) SaveB64Png(filename string) error {
	return os.WriteFile(filename, b.Bytes(), 0644)
}

func (b B64Result) Md5Hash() string {
	clean := b.RemovePrefixAll()
	return fmt.Sprintf("%x", md5.Sum(clean.Bytes()))
}
