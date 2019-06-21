package factory_lib

type Factory_1 struct {
}

func ( f1 Factory_1 ) Construct_A() abstract_A {
	return A1{}
}

func ( f1 Factory_1 ) Construct_B() abstract_B {
	return B1{}
}

//----------------------------------------------------------------------------------

// type factory_2 struct {
// }

// func ( f2 factory_2 ) construct_A() A2 {
// 	return A2
// }

// func ( f2 factory_1 ) construct_B() B2 {
// 	return B2
// }