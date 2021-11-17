package utils

import "net"

func PrivateIPv4() string {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, a := range as {
		ipNet, ok := a.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}

		ip := ipNet.IP.To4()
		if IsPrivateIPv4(ip) {
			return ip.String()
		}
	}
	return ""
}

func IsPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}
