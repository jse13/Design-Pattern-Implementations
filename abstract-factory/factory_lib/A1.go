package factory_lib

import "fmt"

// Concrete implementation of the A product that's part of family 1
type A1 struct {
}

func ( a1 A1 ) Print() {
	fmt.Println( "A1" )
}