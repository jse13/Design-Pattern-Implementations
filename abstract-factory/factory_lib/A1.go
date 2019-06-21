package factory_lib

import "fmt"

type A1 struct {
}

func ( a1 A1 ) Print() {
	fmt.Println( "A1" )
}