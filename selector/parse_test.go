package selector

import (
	"testing"
)

func TestPrase(t *testing.T) {
	s := "foo:contains(привет) <-body:visible"
	if r := Deep(s); r != nil {
		t.Log(r.JSON())
	}

}
