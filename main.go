package main

import (
	"fmt"

	"github.com/m-nt/gim/src/rope"
)

func main() {
	fmt.Println("hey this is not a commit")
	str := "hellow world"
	the_rope := rope.Rope{}
	the_rope.From_str(&str)
	str2 := " my life "
	the_rope.Print()
	str3 := "another one"
	// time.Sleep(time.Second)
	the_rope.Insert(10, &str3)
	// res := the_rope.Report()
	// fmt.Printf("%s\n", *res)
	the_rope.Insert(15, &str2)
	// fmt.Printf("result: %+v\n", the_rope)
	// time.Sleep(time.Second)
	// depth := the_rope.Depth()
	// fmt.Printf("%d\n", depth)
	// res2 := the_rope.Report()
	// fmt.Printf("%s\n", *res2)
	the_rope.Print()
}
