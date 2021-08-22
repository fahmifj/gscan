package scanner

type Options struct {
	Hostname string
	Ports    []string
	Threads  int
	TimeOut  int
	Verbose  bool
}

func NewOptions() *Options {
	return &Options{}
}
