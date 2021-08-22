package util

import (
	"fmt"
	"regexp"
	"strings"
)

func ToLower(s string) string {
	return strings.ToLower(s)
}

// JoinAddr joins host and port into an acceptable address format (socket)
func JoinAddr(host, port string) string {
	addr := fmt.Sprintf("%s:%s", host, port)
	return addr
}

// SplitAddr splits a given host:port
func SplitAddrToHostPort(addr string) (string, string) {
	if addr == "" {
		return "", ""
	}
	s := strings.Split(addr, ":")
	ip := s[0]
	port := s[1]
	return ip, port
}

// SplitAddr splits a given host:port
func SplitAddrToPort(addr string) string {
	_, port := SplitAddrToHostPort(addr)
	return port
}

// Todo:
func CheckHostAliveness(host string) bool {
	ip := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	switch ip.MatchString(host) {
	case true:
		// do something
	case false:
		// do something
	}
	return true
}
