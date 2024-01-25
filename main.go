package main

import (
	"fmt"
	"math"
	"net"
)

const (
	ip     = "192.19.193.209"
	subnet = "29"
)

func main() {
	ip, network, err := net.ParseCIDR(fmt.Sprintf("%s/%s", ip, subnet))
	if err != nil {
		fmt.Println(err)
	}

	size, bits := network.Mask.Size()
	usableAddresses := math.Pow(2, float64(bits)-float64(size))
	availableAddresses := usableAddresses - 2
	subnets := 256 / usableAddresses

	fmt.Println("Usable addresses:", usableAddresses)
	fmt.Println("Available addresses:", availableAddresses)
	fmt.Println("Total Subnets:", subnets)

	fmt.Println()

	fmt.Println("IP:", ip)
	fmt.Printf("Range: %s - %s\n", network.IP, getLast(network))

	fmt.Println()
	fmt.Println("Subnet:", network)
}

func getLast(network *net.IPNet) net.IP {
	last := make(net.IP, len(network.IP))
	copy(last, network.IP)
	for i := range last {
		last[i] |= ^network.Mask[i]
	}
	return last
}
