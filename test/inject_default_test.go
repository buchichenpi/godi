package test

import (
	"github.com/buchichenpi/godi/godiInject"
	"testing"
)

type Parent struct {
	Child TestI `inject:""`
}
type Child struct {
}

func (c *Child) get() {
}

func TestDefaultInject(t *testing.T) {
	p := &Parent{}
	godiInject.Provide(p, &Child{})
	godiInject.Populate()
	t.Logf("%v\n", p)
}
