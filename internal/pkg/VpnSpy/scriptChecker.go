package VpnSpy

import (
	"net"
	"strconv"
	"strings"
)

func checkBlock(ip string, port int) (bool, error) {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil && !strings.Contains(err.Error(), "connection refused") && !strings.Contains(err.Error(), "timeout") {
		return false, err
	} else if err != nil {
		return false, nil
	}
	defer conn.Close()
	return true, nil
}
