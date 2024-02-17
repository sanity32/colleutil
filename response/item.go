package response

import (
	"fmt"
)

type Item[T any] struct {
	DocumentID string `json:"documentId"`
	FrameId    int    `json:"frameId"`
	Result     *T     `json:"result"`
}

func (i Item[T]) reportZeroInt() {
	fmt.Println("DEBUG WARNING \n\n\n Result is zero integer.")
	fmt.Println("Further check is required. \n\nIt was not OK answer previously.")
}

func (i Item[T]) Ok() bool {
	return i.Result != nil
}

func (i Item[T]) zeroInt(v any) bool {
	if v == nil {
		return false
	}
	t := fmt.Sprintf("%T", v)
	isInt := t == "int" || t == "*int"
	if !isInt {
		return false
	}
	r, _ := v.(int)
	return isInt && r == 0
}

func (i Item[T]) Get() (v T, ok bool) {
	if r := i.Result; r != nil {
		if i.zeroInt(*i.Result) {
			i.reportZeroInt()
		}
		return *r, true
	}
	return
}

// returns zero value if no result
func (i Item[T]) GetResult() (r T) {
	if i.Result != nil {
		r = *i.Result
	}
	return r
}
