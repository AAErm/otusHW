package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var timeout = flag.Duration("timeout", 10*time.Second, "timeout")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("must be 2 args")
		return
	}

	addr := net.JoinHostPort(args[0], args[1])
	client := NewTelnetClient(addr, *timeout, os.Stdin, os.Stdout)
	defer client.Close()

	err := client.Connect()
	if err != nil {
		fmt.Printf("failed to client connect %v", err)

		return
	}

	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGHUP, syscall.SIGINT)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
			cancel()
		}()

		if err := client.Receive(); err != nil {
			os.Stderr.Write([]byte(err.Error() + "\n"))

			return
		}
	}()

	go func() {
		defer func() {
			wg.Done()
			cancel()
		}()

		if err := client.Send(); err != nil {
			os.Stderr.Write([]byte(err.Error() + "\n"))
			return
		}
	}()

	<-ctx.Done()

	wg.Wait()
}
