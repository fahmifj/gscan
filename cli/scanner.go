package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/fahmifj/gscan/scanner"
	"github.com/fahmifj/gscan/util"
)

func parseFlag() (*scanner.CLIOptions, error) {
	var err error
	opts := scanner.NewCliOpt()

	flag.StringVar(&opts.Hostname, "h", "", "IP address or resolvable hostname [e.g: www.google.com]")
	port := flag.String("p", "1-1024", "Port range or single port only [e.g:-p 80,443; -p 1-5000]")
	flag.IntVar(&opts.Threads, "t", 10, "Number of threads [max: 50]")
	flag.StringVar(&opts.Mode, "m", "tcp", "Scan mode [tcp|udp]")
	flag.Parse()

	if opts.Hostname == "" {
		return nil, fmt.Errorf("hostname must be specified")
	}

	if opts.Threads < 0 || opts.Threads > 50 {
		return nil, fmt.Errorf("threads must be a positive number and max value of 50")
	}

	opts.Ports, err = util.ParsePorts(*port)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func Execute() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opts, err := parseFlag()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	scan := scanner.NewScanner(opts)
	start := time.Now()
	fmt.Printf("[*] Scan starting at %s\n", start.Format("2006.01.02 15:04:05"))
	fmt.Printf("[*] Target host: %s\n", opts.Hostname)
	fmt.Printf("[*] Threads: %d\n\n", opts.Threads)
	fmt.Printf("PORT	STATE	SERVICE\n")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		select {
		case <-signalChan:
			fmt.Println("[!] Keyboard interrupt detected, exiting...")
			cancel()
		case <-ctx.Done():

		}
	}()
	scan.Run(ctx)
	duration := time.Since(start)
	fmt.Printf("[+] Completed in %.3fs\n", duration.Seconds())
}
