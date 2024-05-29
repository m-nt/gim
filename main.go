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
	the_rope.Print()
	// time.Sleep(time.Second)
	the_rope.Insert(10, "another one")
	fmt.Printf("%d\n", the_rope.Depth())
	the_rope.Print()
	// res := the_rope.Report()
	// fmt.Printf("%s\n", *res)
	the_rope.Insert(15, "ok")
	fmt.Printf("%d\n", the_rope.Depth())
	the_rope.Print()
	// fmt.Printf("str: %s\n", *res2)
	the_rope.Insert(13, "THIS IS WILD")
	fmt.Printf("%d\n", the_rope.Depth())
	the_rope.Print()
	the_rope.Append("THIS IS APPENDED")
	the_rope.Prepend("THIS IS PREPENDED")
	// fmt.Printf("result: %+v\n", the_rope)
	// time.Sleep(time.Second)
	fmt.Printf("%d\n", the_rope.Depth())
	// res2 := the_rope.Report()
	// fmt.Printf("%s\n", *res2)
	the_rope.Print()
}
