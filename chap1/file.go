package chap1

import (
	"fmt"
	"github.com/joomcode/errorx"
	"golang.org/x/sys/unix"
)

func fileOps() error {
	fd, err := unix.Open("data.txt", unix.O_CREAT|unix.O_RDWR|unix.O_APPEND, 0)
	if err != nil {
		return errorx.Decorate(err, "fail to open file")
	}
	if _, err = unix.Write(fd, []byte("world\n")); err != nil {
		return errorx.Decorate(err, "fail to write file")
	}
	_ = unix.Close(fd)

	fd, err = unix.Open("data.txt", unix.O_RDONLY, 0)
	if err != nil {
		return errorx.Decorate(err, "fail to open file")
	}
	var res = new([100]byte)
	if _, err = unix.Read(fd, res[:]); err != nil {
		return errorx.Decorate(err, "fail to write file")
	}
	_ = unix.Close(fd)
	fmt.Printf("%s", string(res[:]))

	fmt.Println(fd)
	return nil
}
