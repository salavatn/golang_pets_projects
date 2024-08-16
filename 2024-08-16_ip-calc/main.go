package main

import (
	"fmt"
	"strconv"
	"strings"
	// "time"
)

var net_host = "188.234.148.235"
var net_mask = 17
var number = 255

func main() {
	// fmt.Print("Enter IP address:  ")
	// _, _ = fmt.Scanln(&net_host)
	// fmt.Print("Enter mask length: ")
	// _, _ = fmt.Scanln(&net_mask)

	result := getBinaryIP(net_host)
	// fmt.Printf("result1:	%s\n", result)
	fmt.Printf("Address:	%s\n", net_host)
	// fmt.Printf("result2:	%s\n", result)
	fmt.Printf("Network: 	%s\n", getNetwork(result, net_mask))
	// fmt.Printf("result3:	%s\n", result)
	fmt.Printf("HostMin:	%s\n", getHostMin(result, net_mask))
	// fmt.Printf("result3:	%s\n", result)
	// fmt.Printf("Last IP Address: %s\n", lastIP.String())
	fmt.Printf("Broadcast:	%s\n", getBroadcast(result, net_mask))

}

// In-Progress
// Будет получать IP адрес и маску сети на проверку
func validatorNetwork() string {
	return ""
}

// getBinaryIP принимает строку IP-адреса в формате "192.168.20.1"
// и возвращает срез строк, где каждый элемент представляет собой
// бинарное представление соответствующего октета.
// Пример: "192.168.20.1" > [11000000 10101000 00010100 00000001]
func getBinaryIP(myNetwork string) []string {
	var binaryIP = make([]string, 0, 4)

	ipParts := strings.Split(myNetwork, ".")

	for _, part := range ipParts {
		ipOctet, _ := strconv.Atoi(part)
		binaryOctet := strconv.FormatInt(int64(ipOctet), 2)

		if len(binaryOctet) < 8 {
			prefix := strings.Repeat("0", 8-len(binaryOctet))
			binaryOctet = prefix + binaryOctet
		}

		binaryIP = append(binaryIP, binaryOctet)
	}

	return binaryIP
}

// getNetwork принимает срез строк (бинарный IP-адрес) и
// целое число (маску сети). Возвращает строку (адрес сети).
// Например, принимает [11000000 10101000 11111111 11111111] и маску 15.
// Возвращает "192.254.0.0"
func getNetwork(binaryNet []string, netmask int) string {
	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)

	// binaryNetwork := binaryNet
	var data []string

	if netmask >= 24 {
		networkOctet := 8 - (32 - netmask)
		octet := binaryNetwork[3]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		binaryNetwork[3] = octet

	} else if netmask >= 16 && netmask < 24 {
		networkOctet := 8 - (32 - netmask - 8)
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		binaryNetwork[3] = "00000000"
		binaryNetwork[2] = octet

	} else if netmask >= 8 && netmask < 16 {
		networkOctet := 8 - (32 - netmask - 8 - 8)
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		binaryNetwork[3] = "00000000"
		binaryNetwork[2] = "00000000"
		binaryNetwork[1] = octet
	}

	for _, octet := range binaryNetwork {
		ip_octet, _ := strconv.ParseInt(octet, 2, 64)
		str := fmt.Sprintf("%d", ip_octet)
		data = append(data, str)
	}

	network := strings.Join(data, ".")
	mask := strconv.Itoa(netmask)
	return network + "/" + mask
}

func getBroadcast(binaryNet []string, netmask int) string {
	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)
	var data []string

	if netmask >= 24 {
		networkOctet := 8 - (32 - netmask)
		octet := binaryNetwork[3]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		binaryNetwork[3] = octet

	} else if netmask >= 16 && netmask < 24 {
		networkOctet := 8 - (32 - netmask - 8)
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		binaryNetwork[3] = "11111111"
		binaryNetwork[2] = octet

	} else if netmask >= 8 && netmask < 16 {
		networkOctet := 8 - (32 - netmask - 8 - 8)
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		binaryNetwork[3] = "11111111"
		binaryNetwork[2] = "11111111"
		binaryNetwork[1] = octet
	}

	for _, octet := range binaryNetwork {
		ip_octet, _ := strconv.ParseInt(octet, 2, 64)
		str := fmt.Sprintf("%d", ip_octet)
		data = append(data, str)
	}
	// fmt.Println("getBroadcast:  ", binaryNetwork)
	network := strings.Join(data, ".")
	return network
}

func getHostMin(binaryNet []string, netmask int) string {
	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)
	var data []string

	if netmask >= 24 {
		networkOctet := 8 - (32 - netmask)
		octet := binaryNetwork[3]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		// fmt.Println("getHostMin, octet1: ", octet)
		octet = octet[:7] + "1"
		// fmt.Println("getHostMin, octet2: ", octet)

		binaryNetwork[3] = octet

	} else if netmask >= 16 && netmask < 24 {
		networkOctet := 8 - (32 - netmask - 8)
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		binaryNetwork[3] = "00000001"
		binaryNetwork[2] = octet

	} else if netmask >= 8 && netmask < 16 {
		networkOctet := 8 - (32 - netmask - 8 - 8)
		octet := binaryNetwork[2]
		fmt.Println("Octet1: ", octet)
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		fmt.Println("Octet2: ", octet)
		binaryNetwork[3] = "00000001"
		binaryNetwork[2] = "00000000"
		binaryNetwork[1] = octet
	}

	for _, octet := range binaryNetwork {
		ip_octet, _ := strconv.ParseInt(octet, 2, 64)
		str := fmt.Sprintf("%d", ip_octet)
		data = append(data, str)
	}
	// fmt.Println("getHostMin:    ", binaryNetwork)
	network := strings.Join(data, ".")
	return network
}
