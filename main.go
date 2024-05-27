package main

import (
	"fmt"
	"time"

	"github.com/m-nt/gim/src/rope"
)

func main() {
	fmt.Println("hey this is not a commit")
	str := "hellow world"
	the_rope := rope.Rope{}
	the_rope.From_str(&str)
	time.Sleep(time.Second)
	// fmt.Printf("result: %+v\n", the_rope)
	the_rope.Print()

}
