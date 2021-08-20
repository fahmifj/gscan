package cli

import (
	"flag"
	"fmt"
	"time"

	"github.com/fahmifj/gscan/scanner"
)

type CLIOptions struct {
	Hostname string
	Ports    string
	// Threads  int
}

func NewCliOpt() *CLIOptions {
	return &CLIOptions{}
}

func parseFlag(opt *CLIOptions) *CLIOptions {
	flag.StringVar(&opt.Hostname, "h", "", "IP address or resolvable hostname [e.g: www.google.com]")
	flag.StringVar(&opt.Ports, "p", "1-1024", "Port range or single port only [e.g:-p 80,443; -p 1-5000]")
	// flag.IntVar(&opt.Threads, "t", 10, "Number of threads [max: 100]")
	flag.Parse()

	// if opt.Threads < 0 {
	// 	fmt.Println("[Error] Threads must be a positive number.")
	// 	return nil
	// }

	return opt
}

func Run() {
	opt := parseFlag(NewCliOpt())
	if opt == nil {
		return
	}
	goscan := scanner.NewGoScan(opt.Hostname, opt.Ports)
	if goscan == nil {
		return
	}
	start := time.Now()
	fmt.Printf("[*] Scan starting at %s\n", start.Format("2006.01.02 15:04:05"))
	fmt.Printf("[*] Target host: %s\n", goscan.Hostname)
	fmt.Printf("PORT	STATE	SERVICE\n")
	for _, port := range goscan.Ports {
		res := scanner.TCPScan(goscan.Hostname, port)
		scanner.Printer(res)

	}
	duration := time.Since(start)
	fmt.Printf("[+] Completed in %.3fs\n", duration.Seconds())

}
