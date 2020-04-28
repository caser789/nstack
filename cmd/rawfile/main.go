package main

import (
	"github.com/caser789/nstack/tcpip/link/rawfile"
	"github.com/caser789/nstack/tcpip/link/tun"

	"fmt"
	"sync"
)

func main() {
	name := "tunabc"
	fd, err := tun.Open(name)

	fmt.Println("tun")
	fmt.Println(fd)
	fmt.Println(err)

	mtu, err := rawfile.GetMTU(name)
	fmt.Println("mtu")
	fmt.Println(mtu)
	fmt.Println(err)

	var wg sync.WaitGroup

	// doesn't work. Invalid parameter
	wg.Add(1)
	go func(fd int, wg *sync.WaitGroup) {
		fmt.Println("NonBlocking Write")

		defer wg.Done()

		buf := []byte("test nonblocking write")
		err := rawfile.NonBlockingWrite(fd, buf)
		fmt.Println(err)

	}(fd, &wg)

	wg.Add(1)
	go func(fd int, wg *sync.WaitGroup) {
		fmt.Println("Blocking read")

		defer wg.Done()

		b := make([]byte, 100)
		n, err := rawfile.BlockingRead(fd, b)

		fmt.Println(b)
		fmt.Println(n)
		fmt.Println(err)

	}(fd, &wg)

	wg.Wait()
}
