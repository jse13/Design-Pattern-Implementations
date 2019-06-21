package main
import "./factory_lib"

func run( factory factory_lib.Abstract_Factory ) {
	a := factory.Construct_A()
	a.Print()
	b := factory.Construct_B()
	b.Scale()
}

func main() {
	factory := factory_lib.Factory_1{}
	run( factory )
}