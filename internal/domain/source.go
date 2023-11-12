package domain

type Source interface {
	LoadSystem(sys System)
	Listen()
}
