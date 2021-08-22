package cli

import (
	"flag"
	"fmt"
	"time"

	"github.com/fahmifj/gscan/scanner"
	"github.com/fahmifj/gscan/util"
)

func parseFlag() (*scanner.Options, error) {
	var err error
	opts := scanner.NewOptions()

	flag.StringVar(&opts.Hostname, "h", "", "IP address or resolvable hostname [e.g: www.google.com]")
	port := flag.String("p", "1-1000", "Port range or single port only [e.g:-p 80,443; -p 1-5000]")
	flag.IntVar(&opts.Threads, "t", 10, "Number of threads [max: 1000]")
	flag.BoolVar(&opts.Verbose, "v", false, "Show closed ports")
	flag.IntVar(&opts.TimeOut, "timeout", 500, "timeout in milliseconds")
	// flag.StringVar(&opts.Mode, "m", "tcp", "Scan mode [tcp|udp]")
	flag.Parse()

	if opts.Hostname == "" {
		return nil, fmt.Errorf("hostname must be specified")
	}

	if opts.Threads < 0 || opts.Threads > 1000 {
		return nil, fmt.Errorf("threads must be a positive number or max value of 1000")
	}

	opts.Ports, err = util.ParsePorts(*port)
	if err != nil {
		return nil, err
	}
	if opts.TimeOut < 0 {
		return nil, fmt.Errorf("timeout must be a positive number")
	}

	return opts, nil
}

func Execute() {
	// not familiar with ctx yet.
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	opts, err := parseFlag()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	scan := scanner.NewScanner(opts)
	fmt.Printf("[*] Scan starting at %s\n", time.Now().Format("2006.01.02 15:04:05"))
	fmt.Printf("[*] Target host: %s\n", opts.Hostname)
	fmt.Printf("[*] Scanning: %d ports\n", len(opts.Ports))
	fmt.Printf("[*] Threads: %d\n\n", opts.Threads)
	scan.Run()

}
