package main

import (
	"fmt"

	"github.com/tanya-mtv/testproject/testmod"
	"github.com/tanya-mtv/testproject/testmod1"
)

func main() {
	fmt.Println("My test")
	fmt.Println(testmod.Hi("Ваня"))

	fmt.Println(testmod1.GodBy("Ваня"))

}
