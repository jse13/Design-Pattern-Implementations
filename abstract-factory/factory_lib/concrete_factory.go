package factory_lib

// Concrete implementation of an Abstract_Factory for family 1
type Factory_1 struct {
}

// Returns abstract type of concrete family 1 type
func ( f1 Factory_1 ) Construct_A() abstract_A {
	return A1{}
}

// Returns abstract type of concrete family 1 type
func ( f1 Factory_1 ) Construct_B() abstract_B {
	return B1{}
}

//----------------------------------------------------------------------------------

// // Alternative family of products that implements Abstract_factory
// type factory_2 struct {
// }

// func ( f2 factory_2 ) construct_A() A2 {
// 	return A2
// }

// func ( f2 factory_1 ) construct_B() B2 {
// 	return B2
// }