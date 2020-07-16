// +build linux

package chap1

import (
	"fmt"
	"github.com/joomcode/errorx"
	"golang.org/x/sys/unix"
	"net"
)

func serverDemo() error {
	serverSock, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		return errorx.Decorate(err, "fail to create socket")
	}
	ip := net.ParseIP("192.168.211.129")
	var lsa = &unix.SockaddrInet4{
		Port: 8234,
	}
	copy(lsa.Addr[:], ip.To4())
	if err = unix.Bind(serverSock, lsa); err != nil {
		return errorx.Decorate(err, "fail to bind socket")
	}
	if err = unix.Listen(serverSock, 5); err != nil {
		return errorx.Decorate(err, "fail to listen socket")
	}
	nfd, _, err := unix.Accept(serverSock)
	if err != nil {
		return errorx.Decorate(err, "fail to accept socket")
	}
	res, err := unix.Write(nfd, []byte("hello"))
	if err != nil {
		return errorx.Decorate(err, "fail to write socket")
	}
	fmt.Println(res)
	fmt.Println(serverSock)
	return nil
}

func clientDemo() error {
	sock, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		return errorx.Decorate(err, "client fail to create socket")
	}
	ip := net.ParseIP("192.168.211.129")
	var lsa = &unix.SockaddrInet4{
		Port: 8234,
	}
	copy(lsa.Addr[:], ip.To4())
	if err = unix.Connect(sock, lsa); err != nil {
		return errorx.Decorate(err, "client fail to connect remote socket")
	}
	var res = new([100]byte)
	num, err := unix.Read(sock, res[:])
	if err != nil {
		return errorx.Decorate(err, "client fail to read remote socket")
	}
	fmt.Println(num)
	fmt.Printf("response is %s\n", string(res[:]))
	return nil
}
