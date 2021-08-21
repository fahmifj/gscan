package scanner

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/fahmifj/gscan/util"
)

type Scanner struct {
	*CLIOptions
}

func NewScanner(opts *CLIOptions) *Scanner {
	return &Scanner{opts}
}

func (s *Scanner) Run(ctx context.Context) {

	// portChannel is used for transfering a single port
	portChan := make(chan string)

	// if ports is less than the specified threads, make it 1 thread for 1 port.
	if len(s.Ports) < s.Threads {
		s.Threads = len(s.Ports)
	}

	// Create goroutines according to the thread value for scanning
	// and wg is needed to control the goroutines
	var scanWg sync.WaitGroup
	scanWg.Add(s.Threads)
	for i := 0; i < s.Threads; i++ {
		go func() {
			// Inform that the scan has been done upon exiting this goroutine.
			defer scanWg.Done()
			for {
				// receive single port from portChannel and do scan
				port, ok := <-portChan

				// if no more data, just exit the loop
				if !ok {
					break
				}

				// Do scan
				res := s.TCPScan(port)
				// Print the results, looks dumb, but still works
				if res.State == "open" {
					fmt.Printf("%v	%v	%v\n", res.Port, res.State, res.Service)
				}
			}

		}()
	}
	// Iterate over port send each port to portChannel
	for _, port := range s.Ports {
		portChan <- port
	}
	// close channels because all ports has been transferred
	// but wait until all goroutines scanner done.
	close(portChan)
	scanWg.Wait()
	ctx.Done()

}

func (s *Scanner) TCPScan(port string) Result {
	svc := util.KnownPorts[port]
	// construct address
	addr := util.JoinAddr(util.ToLower(s.Hostname), port)
	// idk why this returns nil
	conn, err := net.DialTimeout(util.ToLower(s.Mode), addr, 1*time.Second)
	// conn, err := net.Dial("tcp", addr)
	if err != nil {
		// if error, assume that the port is closed
		return Result{
			Port:    port,
			State:   "closed",
			Service: svc,
		}
	}
	defer conn.Close()
	return Result{
		Port:    port,
		State:   "open",
		Service: svc,
	}
}
