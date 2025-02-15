package gterm

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

var saved_term = unix.Termios{}

func isatty(fd uintptr) bool {
	var garbage [128]byte
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd),
		unix.TCGETS,                       // Either TCGETS or TIOCGETA, basically request the termios struct
		uintptr(unsafe.Pointer(&garbage)), // A garbage slice to be filled with termios data
		0, 0, 0)
	return err == 0
}

func Reset() {
	err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, &saved_term)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		panic(err.Error())
	}
}
func Set() {
	if !isatty(uintptr(syscall.Stdin)) {
		panic("You shall not pass! (not a terminal)")
	}
	term_tmp, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		panic("Could not get termios")
	}
	saved_term = *term_tmp
	term_tmp.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	term_tmp.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	term_tmp.Oflag &^= syscall.OPOST
	term_tmp.Cflag &^= syscall.CSIZE | syscall.PARENB
	term_tmp.Cflag |= syscall.CS8
	term_tmp.Cc[syscall.VMIN] = 1
	term_tmp.Cc[syscall.VTIME] = 0

	err2 := unix.IoctlSetTermios(unix.Stdin, unix.TCSETS, term_tmp)
	if err2 != nil {
		panic("Could not set termios")
	}
}

type TerminalContext struct {
}

func Open_terminal() {
	defer func() {
		if r := recover(); r != nil {
			Reset()
			fmt.Printf("\x1b[H\x1b[2J")
		}
	}()
	Set()
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
	line := 0
	for {
		run := <-buff
		if run == 0x3 {
			panic("Exit")
		}
		if run == 0xD {
			line++
			fmt.Printf("\x1b[H\x1b[%dB", line)
		}
		if run < 0xFF {
			fmt.Printf("%c", run)
		}
	}
}
