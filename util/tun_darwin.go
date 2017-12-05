package util

import (
	"fmt"
	"log"
	"net"
)

func Ifconfig(tunName, network string, mtu uint32) {
	var ip, ipv4Net, _ = net.ParseCIDR(network)
	ipStr := ip.To4().String()
	sargs := fmt.Sprintf("%s %s %s mtu %d netmask %s up", tunName, ipStr, ipStr, mtu, Ipv4MaskString(ipv4Net.Mask))
	if err := ExecCommand("/sbin/ifconfig", sargs); err != nil {
		log.Fatal("execCommand failed", err)
	}
}