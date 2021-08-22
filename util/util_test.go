package util

import (
	"testing"
)

func TestWantErrWeirdArguments(t *testing.T) {
	port := "-50-asdasd,50,"
	_, err := ParsePorts(port)
	if err == nil {
		t.Errorf("error is expected, got %v", err)
	}
}

func TestWantErrSinglePortButString(t *testing.T) {
	port := "iamf"
	_, err := ParsePorts(port)
	if err == nil {
		t.Errorf("error is expected, got %v", err)
	}
}

func TestWantErrMultiPortButString(t *testing.T) {
	port := "50-iamf"
	_, err := ParsePorts(port)
	if err == nil {
		t.Errorf("error is expected, got %v", err)
	}
	port = "iamf-50"
	_, err = ParsePorts(port)
	if err == nil {
		t.Errorf("error is expected, got %v", err)
	}
}

func TestWantErrLowerPortHigherPort(t *testing.T) {
	port := "0-65535"
	_, err := ParsePorts(port)
	if err == nil {
		t.Errorf("error is expected, got %v", err)
	}
}

func TestSwapHigherPortToLowerPort(t *testing.T) {
	port := "1024-1"
	s, _ := ParsePorts(port)
	if s == nil {
		t.Errorf("Swap is expected, got %v", s)
	}
}

func TestWantErrInvalidPortNumber(t *testing.T) {
	port := "65536"
	s, err := ParsePorts(port)
	if err == nil {
		t.Errorf("error is expected, got %v", s)
	}
}

func TestWantErrInvalidPortRange(t *testing.T) {
	port := "-2-5"
	_, err := ParsePorts(port)
	if err == nil {
		t.Errorf("error is expected, got %v", err)
	}
}

func TestWantErrCheckHost(t *testing.T) {
	t.Skip("Not implemented")
	// host := "192.168.123.123"
	// ok := CheckHost(host)
	// if ok {
	// 	t.Errorf("false is expected, got %v", ok)
	// }
}
func TestValidPortNumber(t *testing.T) {
	port := "80,445,"

	_, err := ParsePorts(port)
	if err != nil {
		t.Errorf("got error %v", err)
	}
}

func TestValidPortRange(t *testing.T) {
	port := "443-445"
	_, err := ParsePorts(port)
	if err != nil {
		t.Errorf("got error %v", err)
	}
}
