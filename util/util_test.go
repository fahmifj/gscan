package util

import (
	"testing"
)

func TestWantErrWeirdArguments(t *testing.T) {
	port := "-50-asdasd,50,"
	s := ParsePorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestWantErrSinglePortButString(t *testing.T) {
	port := "iamf"
	s := ParsePorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestWantErrMultiPortButString(t *testing.T) {
	port := "50-iamf"
	s := ParsePorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
	port = "iamf-50"
	s = ParsePorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestWantErrLowerPortHigherPort(t *testing.T) {
	port := "0-65535"
	s := ParsePorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestSwapHigherPortToLowerPort(t *testing.T) {
	port := "1024-1"
	s := ParsePorts(port)
	if s == nil {
		t.Errorf("Swap is expected, got %v", s)
	}
}

func TestWantErrInvalidPortNumber(t *testing.T) {
	port := "65536"
	s := ParsePorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestWantErrInvalidPortRange(t *testing.T) {
	port := "-2-5"
	s := ParsePorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestWantErrCheckHost(t *testing.T) {
	// host := "192.168.123.123"
	// ok := CheckHost(host)
	// if ok {
	// 	t.Errorf("false is expected, got %v", ok)
	// }
}
func TestValidPortNumber(t *testing.T) {
	port := "80,445,"

	s := ParsePorts(port)
	if s == nil {
		t.Errorf("got nil")
	}
}

func TestValidPortRange(t *testing.T) {
	port := "443-445"

	s := ParsePorts(port)
	if s == nil {
		t.Errorf("got nil")
	}
}
