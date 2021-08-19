package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func JoinAddr(host, port string) string {
	addr := fmt.Sprintf("%s:%s", host, port)
	return addr
}

func SplitAddr(addr string) (string, string) {
	if addr == "" {
		return "", ""
	}
	s := strings.Split(addr, ":")
	ip := s[0]
	port := s[1]
	return ip, port
}

// SplitPorts
func SplitPorts(port string) []string {
	var cleanPorts []string
	if !strings.Contains(port, ",") && !strings.Contains(port, "-") {
		cleanPorts = append(cleanPorts, port)

	} else if strings.Contains(port, ",") {
		dirtyPorts := strings.Split(port, ",")
		for _, val := range dirtyPorts {
			if val != "" {
				cleanPorts = append(cleanPorts, val)
			}
		}

	} else if strings.Contains(port, "-") {
		dirtyPorts := strings.Split(port, "-")
		if len(dirtyPorts) != 2 {
			fmt.Println("[!] Invalid port range")
			return nil
		}
		loPort, _ := strconv.Atoi(dirtyPorts[0])
		hiPort, _ := strconv.Atoi(dirtyPorts[1])
		if hiPort < 0 || loPort < 0 || hiPort > 65535 || loPort > 65535 {
			fmt.Println("[!] Invalid port number")
			return nil
		}
		if hiPort < loPort {
			loPort, hiPort = hiPort, loPort
		}
		// TODO :
		// Add port based on well-known list!
		for i := loPort; i <= hiPort; i++ {
			_, ok := KnownPorts[strconv.Itoa(i)]
			if ok {
				cleanPorts = append(cleanPorts, strconv.Itoa(i))
			}
		}
	}
	return cleanPorts
}

func CheckHost(host string) bool {
	ip := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	switch ip.MatchString(host) {
	case true:
		// do something
	case false:
		// do something
	}
	return true
}

// p := '1,2,3' =>  [ '1','2','3'] => len(p)
// what if p,1,2,3 ?
// if p[element] != ""
//	}

// p:= '1-3' =>  [ '1','2','3'] => min=1 max=3
// if p-1-2?
// if len(p) != 2
//
//separator := []string{"-", ","}

//return
