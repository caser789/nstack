package main

import (
	"fmt"
	"github.com/caser789/nstack/tcpip/link/tun"
)

func main() {
	name := "tunabc"
	fd, err := tun.Open(name)

	fmt.Println(fd)
	fmt.Println(err)
}
