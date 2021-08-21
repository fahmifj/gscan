package scanner

type CLIOptions struct {
	Hostname string
	Ports    []string
	Threads  int
	Mode     string
}

func NewCliOpt() *CLIOptions {
	return &CLIOptions{}
}
