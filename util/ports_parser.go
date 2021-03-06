package util

import (
	"fmt"
	"strconv"
	"strings"
)

// ParsePorts split a single string that contains ports into a slice of ports
func ParsePorts(port string) ([]string, error) {
	var cleanPorts []string

	if !strings.Contains(port, ",") && !strings.Contains(port, "-") {

		// if it doesn't contain any delimiter, assume it as single port
		n, err := strconv.Atoi(port)
		if err != nil {

			return nil, fmt.Errorf("specified port must be a number")
		}

		if n < 0 || n > 65535 {
			return nil, fmt.Errorf("invalid port range/number")
		}
		cleanPorts = append(cleanPorts, port)

	} else if strings.Contains(port, ",") {

		dirtyPorts := strings.Split(port, ",")
		for _, val := range dirtyPorts {
			// This cleanup something like -p,123, 231, => [ '', '123', ' 123', '']
			if val != "" {
				_, err := strconv.Atoi(strings.TrimSpace(val))
				if err != nil {
					return nil, fmt.Errorf("specified port must be a number")
				}
				cleanPorts = append(cleanPorts, val)
			}
		}

	} else if strings.Contains(port, "-") {

		dirtyPorts := strings.Split(port, "-")

		// A strip (-) indicates port range, so the length must be equal to 2, lower port and higher port.
		if len(dirtyPorts) != 2 {

			return nil, fmt.Errorf("invalid port range/number")
		}

		// Before generating port numbers
		// lower port and higher port need to be converted to integer.
		loPort, err := strconv.Atoi(strings.TrimSpace(dirtyPorts[0]))
		if err != nil {
			return nil, fmt.Errorf("specified port must be a number")
		}

		hiPort, err := strconv.Atoi(strings.TrimSpace(dirtyPorts[1]))
		if err != nil {
			return nil, fmt.Errorf("specified port must be a number")
		}

		// Ports specified must be in valid range.
		if hiPort <= 0 || loPort <= 0 || hiPort > 65535 || loPort > 65535 {
			return nil, fmt.Errorf("invalid port range/number")
		}

		// Just swap if the range in reverse order: higher port, lowerport.
		// e.g: -p 65535,1
		if hiPort < loPort {
			loPort, hiPort = hiPort, loPort
		}

		for i := loPort; i <= hiPort; i++ {
			// Map the generated port numbers with known ports
			// so the scanner doesn't have to deal with unassigned ports.
			n := strconv.Itoa(i)
			_, ok := KnownPorts[n]
			if ok {
				cleanPorts = append(cleanPorts, n)
			}
		}
	}
	return cleanPorts, nil
}
