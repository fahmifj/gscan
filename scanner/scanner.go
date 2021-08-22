package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/fahmifj/gscan/util"
)

type Scanner struct {
	*Options
}

func NewScanner(opts *Options) *Scanner {
	return &Scanner{opts}
}

func (s *Scanner) Run() {
	// portChannel is used for transfering a single port
	portChan := make(chan string)
	resultChan := make(chan Result)

	// Create goroutines according to the thread value for scanning
	// and wg is needed to control the goroutines
	var scanWg sync.WaitGroup
	scanWg.Add(s.Threads)
	for i := 0; i < s.Threads; i++ {
		// Do scan
		go s.TCPScan(portChan, &scanWg, resultChan)
	}

	var resWg sync.WaitGroup
	resWg.Add(1)

	// NOTE: should be separated for clearance
	// Print the results in another goroutine
	go func() {
		start := time.Now()
		// This is going to be ugly if the port that has 3 digits number
		fmt.Printf("PORT	STATE	SERVICE\n")
		for {
			res, ok := <-resultChan
			if !ok {
				break
			}
			fmt.Printf("%v/tcp	%v	%v\n", res.Port, res.State, res.Service)
		}
		duration := time.Since(start)
		fmt.Printf("\n")
		fmt.Printf("[+] Completed in %.3fs\n", duration.Seconds())
		resWg.Done()
	}()

	// Iterate over ports and send each port
	// to portChan which is then consumed by TCPScan()
	for _, port := range s.Ports {
		portChan <- port
	}
	// Close channels because all ports has been transferred
	// but wait until all goroutines done.
	close(portChan)
	scanWg.Wait()

	// Not sure about this,
	// but if I put this in the upper, it's going to block the iteration over ports
	close(resultChan)
	resWg.Wait()

}

func (s *Scanner) TCPScan(portChan chan string, wg *sync.WaitGroup, result chan<- Result) {
	defer wg.Done()
	for {
		port, ok := <-portChan
		// if no more data, just exit the loop
		if !ok {
			break
		}

		// Map the known ports to list the assigned service.
		// It's not dynamic, so it can be wrong, (e.g SSH port assigned to 2222 instead of 22).
		svc := util.KnownPorts[port]

		// Construct address that consists of host:port
		addr := util.JoinAddr(util.ToLower(s.Hostname), port)

		// Get connection.
		conn, err := net.DialTimeout("tcp", addr, time.Millisecond*time.Duration(s.TimeOut))
		if conn != nil && err == nil {
			// If it got the connection and no error assume the port open.
			result <- Result{Port: port, State: "open", Service: svc}
			conn.Close()
		} else if s.Verbose {
			// Ignore the error assume that the port is closed
			result <- Result{Port: port, State: "closed", Service: svc}
		}
		// Note: This going to be error because the conn is nil
		//
		// defer conn.Close()
	}

}
