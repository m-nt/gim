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
			fmt.Printf("\x1b[H\x1b[2J")
		}
	}()
	gterm.Set()
	fmt.Printf("\x1b[H\x1b[2J")
	reader := bufio.NewReader(os.Stdin)
	buff := make(chan byte)
	go func() {
		for {
			char, err := reader.ReadByte()
			if err == nil {
				buff <- char
			}
		}
	}()
	for {
		run := <-buff
		if run == 0x3 {
			panic("Exit")
		}
		if run == 0xD {
			fmt.Printf("\x1b[H\x1b[1B")
		}
		if run == 0x63 {
			fmt.Printf("%c", run)
		}
		if run == 0x61 {
			fmt.Printf("%c", run)
		}
		fmt.Printf("%x\n", run)
	}
}
