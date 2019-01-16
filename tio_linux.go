// +build linux

package main

import (
	"golang.org/x/sys/unix"
)

const ioctlReadTermios = unix.TCGETA
const ioctlWriteTermios = unix.TCSETA

