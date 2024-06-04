package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/m-nt/gim/src/gterm"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			gterm.Reset()
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	gterm.Set()
	fmt.Printf("\x1b[H\x1b[2J")
	reader := bufio.NewReader(os.Stdin)
	buff := make(chan byte)
	go func() {
		for {
			char, _ := reader.ReadByte()
			buff <- char
		}
	}()
	for {
		select {
		case run := <-buff:
			_ = run
			if run == 0xA {
				fmt.Printf("\x1b[H\x1b[1B")
			}
			if run == 0x63 {
				fmt.Printf("%c", run)
			}
			if run == 0x61 {
				fmt.Printf("%c", run)
			}
			// fmt.Printf("%b", run)
		default:
			continue
		}
	}
}
