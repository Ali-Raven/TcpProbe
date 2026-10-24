package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TcpProb/internal/probe"
)

func main() {
	host := flag.String("h", "", "Use this like tcprobe -h <HOST>")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(ch)

	portF := flag.String("p", "", "Use this like tcprobe -p <HOST>")
	flag.Parse()

	if len(os.Args[1:]) == 0 {
		fmt.Println("Usage of ./bin/tcprobe\n -h <HOST> -p <PORT>")
		panic("Error : Argument Needed for process.")
	}

	// Default PORT
	port := "443"

	if *portF != "" {
		port = *portF
	}

	var counter probe.Counter
	timeout := 2 * time.Second
	fullAddr := net.JoinHostPort(*host, port)

	fmt.Printf("Probing %s ...\n", *host)
	for {

		_, _ = probe.Probe(timeout, fullAddr, &counter)

		select {
		case <-ch:
			fmt.Println("\nInterrupt received ....")
			fmt.Printf("Sent : %d\n", counter.Sent)
			fmt.Printf("Success : %d\n", counter.Success)
			fmt.Printf("Failed: %d\n", counter.Failed)
			return
		default:
		}
	}
}
