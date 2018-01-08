package codefights

import "net"

func isIPv4Address(inputString string) bool {
	ip := net.ParseIP(inputString)
	return ip != nil && ip.To4() != nil
}
