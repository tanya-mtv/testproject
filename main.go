package main

import (
	"fmt"

	"github.com/tanya-mtv/testproject/testmod"
	"github.com/tanya-mtv/testproject/testmod1"
	"github.com/tanya-mtv/testproject/testmod2"
)

func main() {
	fmt.Println("My test")
	fmt.Println(testmod.Hi("Ваня"))

	fmt.Println(testmod1.GodBy("Ваня"))

	testmod2.Dosomething()

}
