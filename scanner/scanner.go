package scanner

import (
	"fmt"
	"net"

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
	return &Scanner{
		Hostname: h,
		Ports:    util.SplitPorts(p),
	}
}

func TCPScan(host, port string) *Result {
	addr := util.JoinAddr(host, port)
	svc := util.KnownPorts[port]
	conn, err := net.Dial("tcp", addr)
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
