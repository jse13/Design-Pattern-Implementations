package factory_lib

// Interface for A products; each family should have it's own concrete implementation of this
type abstract_A interface {
	Print()
}

// Interface for B products; each family should have it's own concrete implementation of this
type abstract_B interface {
	Scale() int
}