package factory_lib

import "fmt"

type B1 struct {
}

func ( b1 B1 ) Scale() int {
	fmt.Println( "B1" )
	return 2
}