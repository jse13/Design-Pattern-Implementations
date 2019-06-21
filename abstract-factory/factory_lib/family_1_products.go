package factory_lib

import "fmt"

// Concrete implementation of the A product that's part of family 1
type A1 struct {
}

func ( a1 A1 ) Print() {
	fmt.Println( "A1" )
}

// Concrete implementation of the B product that's part of family 1
type B1 struct {
}

func ( b1 B1 ) Scale() int {
	fmt.Println( "B1" )
	return 2
}