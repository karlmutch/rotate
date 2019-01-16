// +build !linux

package main

import (
	"syscall"
)

const ioctlReadTermios = syscall.TIOCGETA
const ioctlWriteTermios = syscall.TIOCSETA

