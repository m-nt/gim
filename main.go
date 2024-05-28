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
	str2 := " my life "
	time.Sleep(time.Second)
	the_rope.Insert(4, &str2)
	// fmt.Printf("result: %+v\n", the_rope)
	time.Sleep(time.Second)
	the_rope.Print()

}
