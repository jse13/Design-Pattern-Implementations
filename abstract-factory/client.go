package main
import af "./factory_lib"

// The "main application loop" that has a factory injected into it.
// The function doesn't care what family A and B belong to so long as they
// implement the same interface.
func run( factory af.Abstract_Factory ) {
	a := factory.Construct_A()
	a.Print()
	b := factory.Construct_B()
	b.Scale()
}

func main() {
	// Create the correct concrete factory for injection into the rest of the application.
	factory := af.Factory_1{}

	run( factory )
}