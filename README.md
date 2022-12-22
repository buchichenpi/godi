This is a wrapper of facebookarchive/inject which is a tool of di used in Go.

这是一个封装了facebookarchive/inject的依赖注入项目

注意事项：
    依照facebookarchive/inject的要求，所有需要注入的字段必须对其他包可见

使用方式：
step1:
```go
func init() {
    godiInject.Provide(objs...)
}
```
如果需要对obj进行初始化，例如进行构造或者做一些其他动作，可以继承Constructor接口：
```go
func (s *SayHello) Init() (obj interface{}, name string) {
	handler := ginwrapper.NewHandler(
		ginwrapper.WithCodec(&codec.CodeImpl{}),
		ginwrapper.WithProtoName("friendship.proto"),
	)
	handler.RegisterService(s)
	server.Handle("/test/", handler)
	return s, ""
}
```
如果需要使用名字注入，则可以在Provide的时候使用ObjectWithName
```go
godiInject.Provide(p, &godiInject.ObjectWithName{
		Object: &Child2{1},
		Bean:   "b",
	}, &godiInject.ObjectWithName{
		Object: &Child2{2},
		Bean:   "c",
	})
```
step2:
```go
func main() {
    godi.Populate()
}
```

具体的使用案例可以查看test目录。

