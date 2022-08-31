package main

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"syscall"
)

func goID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func reuseportGoroutine() {
	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gID := goID()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Client [%s] Received msg from Server gID: [%d] \n", r.RemoteAddr, gID)
	})
	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			var lc = net.ListenConfig{
				Control: func(network, address string, c syscall.RawConn) error {
					var opErr error
					if err := c.Control(func(fd uintptr) {
						opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
					}); err != nil {
						return err
					}
					return opErr
				},
			}
			l, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:8000")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Server with PID: [%d] is running \n", goID())
			_ = server.Serve(l)
		}()
	}
	wg.Wait()
}

var lc = net.ListenConfig{
	Control: func(network, address string, c syscall.RawConn) error {
		var opErr error
		if err := c.Control(func(fd uintptr) {
			opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		}); err != nil {
			return err
		}
		return opErr
	},
}

func main() {
	pid := os.Getpid()
	l, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Client [%s] Received msg from Server PID: [%d] \n", r.RemoteAddr, pid)
	})
	fmt.Printf("Server with PID: [%d] is running \n", pid)
	_ = server.Serve(l)
}