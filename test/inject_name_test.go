package test

import (
	"github.com/buchichenpi/godi/godiInject"
	"testing"
)

type TestI interface {
	get()
}

type Parent2 struct {
	C TestI `inject:"c"`
	B TestI `inject:"b"`
}
type Child2 struct {
	A int
}

func (c *Child2) get() {
}

func TestNameInject(t *testing.T) {
	p := &Parent2{}
	godiInject.Provide(p, &godiInject.ObjectWithName{
		Object: &Child2{1},
		Bean:   "b",
	}, &godiInject.ObjectWithName{
		Object: &Child2{2},
		Bean:   "c",
	})
	godiInject.Populate()
	t.Logf("%v\n", p)
}
