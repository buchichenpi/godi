package godiInject

type ObjectWithNameI interface {
	GetName() string
	GetBean() interface{}
}
type ObjectWithName struct {
	Object interface{}
	Bean   string
}

func (o *ObjectWithName) GetName() string {
	return o.Bean
}
func (o *ObjectWithName) GetBean() interface{} {
	return o.Object
}
