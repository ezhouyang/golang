// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import syscall "syscall"
import os "os"

// FDs are wrappers for file descriptors
export type FD struct {
	fd int64
}

export func NewFD(fd int64) *FD {
	if fd < 0 {
		return nil
	}
	n := new(FD);
	n.fd = fd;
	return n;
}

export var (
	Stdin = NewFD(0);
	Stdout = NewFD(1);
	Stderr = NewFD(2);
)

export func Open(name string, mode int64, flags int64) (fd *FD, err *Error) {
	r, e := syscall.open(name, mode, flags);
	return NewFD(r), ErrnoToError(e)
}

func (fd *FD) Close() *Error {
	if fd == nil {
		return EINVAL
	}
	r, e := syscall.close(fd.fd);
	fd.fd = -1;  // so it can't be closed again
	return ErrnoToError(e)
}

func (fd *FD) Read(b *[]byte) (ret int64, err *Error) {
	if fd == nil {
		return -1, EINVAL
	}
	r, e := syscall.read(fd.fd, &b[0], int64(len(b)));
	return r, ErrnoToError(e)
}

func (fd *FD) Write(b *[]byte) (ret int64, err *Error) {
	if fd == nil {
		return -1, EINVAL
	}
	r, e := syscall.write(fd.fd, &b[0], int64(len(b)));
	return r, ErrnoToError(e)
}

func (fd *FD) WriteString(s string) (ret int64, err *Error) {
	if fd == nil {
		return -1, EINVAL
	}
	b := new([]byte, len(s)+1);
	if !syscall.StringToBytes(b, s) {
		return -1, EINVAL
	}
	r, e := syscall.write(fd.fd, &b[0], int64(len(s)));
	return r, ErrnoToError(e)
}
