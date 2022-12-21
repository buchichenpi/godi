package test

import (
	"github.com/buchichenpi/godi/godiInject"
	"testing"
	"time"
)

// 测试带有初始化方法

type Parent3 struct {
	Count int
	B     TestI `inject:""`
}

func (p *Parent3) Init() (interface{}, string) {
	p.Count = 11
	return p, ""
}

type Child3 struct {
	A int
}

func (c *Child3) get() {
}

func (c *Child3) Init() (interface{}, string) {
	c.A = time.Now().Second()
	return c, ""
}

func TestInitInject(t *testing.T) {
	p := &Parent3{}
	godiInject.Provide(p, &Child3{})
	godiInject.Populate()
	t.Logf("%v\n", p)
}
