package factory_lib

import "fmt"

// Concrete implementation of the B product that's part of family 2
type B1 struct {
}

func ( b1 B1 ) Scale() int {
	fmt.Println( "B1" )
	return 2
}