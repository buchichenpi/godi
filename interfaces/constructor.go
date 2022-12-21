package interfaces

type Constructor interface {
	Construct() (obj interface{}, name string) // name不能是private和inline
}
