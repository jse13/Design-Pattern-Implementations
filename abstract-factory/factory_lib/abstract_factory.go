package factory_lib

// Interface to a concrete factory type.
// Exposes a factory method for each family member, but the
// implementation is up to the concrete family factory.
type Abstract_Factory interface {
	Construct_A() abstract_A
	Construct_B() abstract_B
}