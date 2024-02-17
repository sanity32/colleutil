package selector

import (
	"testing"
)

func TestPrase(t *testing.T) {
	s := "foo:contains(привет) <-body:visible"
	if r := ParseDeep(s); r != nil {
		t.Log(r.JSON())
	}

}
