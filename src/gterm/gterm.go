package gterm

import (
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
	err := unix.IoctlSetTermios(unix.Stdin, unix.TCSETA, &saved_term)
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
	// term_tmp.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	// term_tmp.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	// term_tmp.Oflag &^= syscall.OPOST
	// term_tmp.Cflag &^= syscall.CSIZE | syscall.PARENB
	// term_tmp.Cflag |= syscall.CS8
	// term_tmp.Cc[syscall.VMIN] = 1
	// term_tmp.Cc[syscall.VTIME] = 0

	err2 := unix.IoctlSetTermios(unix.Stdin, unix.TCSETA, term_tmp)
	if err2 != nil {
		panic("Could not set termios")
	}
}
