package main

import (
	"fmt"
	"os"
	"primitive-webapp/primitive"
)

func main() {
	strstd, err := primitive.Primitive("golang.jpeg", "out.jpeg", 10, primitive.WithMode(primitive.Circle))
	if err != nil {
		panic(err) 
	}
	fmt.Fprint(os.Stdout, strstd)

}
