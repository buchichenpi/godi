package godiInject

import (
	"github.com/buchichenpi/godi/interfaces"
	"github.com/facebookarchive/inject"
)

var beans []*inject.Object // 存放初始化后的实例

func Provide(objects ...interface{}) {

	for _, obj := range objects {
		if obj == nil {
			continue
		}

		if constructor, ok := obj.(interfaces.Constructor); ok {
			bean, name := constructor.Init()
			if bean != nil {
				beans = append(beans, &inject.Object{
					Value: bean,
					Name:  name,
				})
			}
		} else if bean, ok := obj.(*inject.Object); ok {
			// 支持业务代码中使用原生inject.Object
			beans = append(beans, bean)
		} else if v, ok := obj.(ObjectWithNameI); ok {
			beans = append(beans, &inject.Object{
				Value: v.GetBean(),
				Name:  v.GetName(),
			})
		} else { // 普通的实例
			beans = append(beans, &inject.Object{
				Value: obj,
			})
		}
	}
}

func Populate() {
	graph := inject.Graph{}
	graph.Provide(beans...)
	err := graph.Populate()
	if err != nil {
		panic(err)
	}
}
