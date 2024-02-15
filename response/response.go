package colleresponse

import (
	"encoding/json"
	"errors"
)

var ErrParseFailure = errors.New("failed to parse")

func Parse[T any](input any) (r Response[T], err error) {
	return r, r.Populate(input)
}

type Response[T any] []Item[T]

func (aa Response[T]) Ok() bool {
	return len(aa) != 0
}

func (aa *Response[T]) Populate(input any) error {
	if j, err := json.Marshal(input); err != nil {
		return err
	} else {
		return json.Unmarshal(j, &aa)
	}
}

func (aa *Response[T]) FilterOnlyOk() (r Response[T]) {
	for _, item := range *aa {
		if item.Ok() {
			r = append(r, item)
		}
	}
	return
}

// If no answers, returns empty answer, i.e. Ok() == false
func (aa *Response[T]) First() (r Item[T]) {
	if a := aa.FilterOnlyOk(); a.Ok() {
		return a[0]
	}
	return
}

func (aa *Response[T]) Find(frameID int) (r Item[T]) {
	for _, item := range *aa {
		if item.FrameId == frameID {
			return item
		}
	}
	return
}
