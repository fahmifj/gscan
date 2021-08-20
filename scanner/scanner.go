package scanner

import (
	"fmt"
	"net"
	"time"

	"github.com/fahmifj/gscan/util"
)

type Scanner struct {
	Hostname string
	Ports    []string
}

type Result struct {
	Port    string
	State   string
	Service string
}

func NewGoScan(h, p string) *Scanner {
	port := util.ParsePorts(p)
	if port == nil {
		return nil
	}
	return &Scanner{
		Hostname: h,
		Ports:    port,
	}
}

func TCPScan(host, port string) *Result {
	addr := util.JoinAddr(host, port)
	svc := util.KnownPorts[port]
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return &Result{

			Port:    port,
			State:   "close",
			Service: svc,
		}
	}
	defer conn.Close()

	return &Result{
		Port:    port,
		State:   "open",
		Service: svc,
	}
}

func Printer(res *Result) {

	// fmt.Printf("%#v\n", res)
	fmt.Printf("tcp/%s	%s	%s\n", res.Port, res.State, res.Service)

}
