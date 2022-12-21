package interfaces

type Constructor interface {
	Init() (obj interface{}, name string) // name不能是private和inline
}
