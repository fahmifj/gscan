package util

import (
	"testing"
)

func TestWantErrInvalidPortNumber(t *testing.T) {
	port := "1-65536"
	s := SplitPorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestWantErrInvalidPortRange(t *testing.T) {
	port := "-2-5"
	s := SplitPorts(port)
	if s != nil {
		t.Errorf("nil is expected, got %v", s)
	}
}

func TestWantErrCheckHost(t *testing.T) {
	host := "192.168.123.123"
	ok := CheckHost(host)
	if ok {
		t.Errorf("false is expected, got %v", ok)
	}
}
func TestValidPortNumber(t *testing.T) {
	port := "80,445,"

	s := SplitPorts(port)
	if s == nil {
		t.Errorf("got nil")
	}
}

func TestValidPortRange(t *testing.T) {
	port := "443-445"

	s := SplitPorts(port)
	if s == nil {
		t.Errorf("got nil")
	}
}
