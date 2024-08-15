package main

import (
	"fmt"
	"strconv"
	"strings"
	// "time"
)

var net_host = "192.168.255.255"
var net_mask = 25
var number = 255

func main() {
	fmt.Print("Enter IP address:  ")
	_, _ = fmt.Scanln(&net_host)
	fmt.Print("Enter mask length: ")
	_, _ = fmt.Scanln(&net_mask)

	result := convertIPtoBinary(net_host)

	fmt.Printf("IP Address: 	 %s\n", net_host)
	fmt.Printf("Network Address: %s\n", getNetworkAddress(result, net_mask))
	// fmt.Printf("First IP Address: %s\n", firstIP.String())
	// fmt.Printf("Last IP Address: %s\n", lastIP.String())
	// fmt.Printf("Broadcast Address: %s\n", broadcastIP.String())

}

// In-Progress
// Будет получать IP адрес и маску сети на проверку
func validatorNetwork() string {
	return ""
}

// convertIPtoBinary принимает строку IP-адреса в формате "192.168.20.1"
// и возвращает срез строк, где каждый элемент представляет собой
// бинарное представление соответствующего октета.
// Пример: "192.168.20.1" > [11000000 10101000 00010100 00000001]
func convertIPtoBinary(myNetwork string) []string {
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

// getNetworkAddress принимает срез строк (бинарный IP-адрес) и
// целое число (маску сети). Возвращает строку (адрес сети).
// Например, принимает [11000000 10101000 11111111 11111111] и маску 15.
// Возвращает "192.254.0.0"
func getNetworkAddress(binaryNetwork []string, maskNetwork int) string {
	var data []string

	if maskNetwork >= 24 {
		network_bits := 8 - (32 - maskNetwork)
		octet := binaryNetwork[3]
		octet = octet[:network_bits] + strings.Repeat("0", 8-network_bits)
		binaryNetwork[3] = octet

	} else if maskNetwork >= 16 && maskNetwork < 24 {
		network_bits := 8 - (32 - maskNetwork - 8)
		octet := binaryNetwork[2]
		octet = octet[:network_bits] + strings.Repeat("0", 8-network_bits)
		binaryNetwork[3] = "00000000"
		binaryNetwork[2] = octet

	} else if maskNetwork >= 8 && maskNetwork < 16 {
		network_bits := 8 - (32 - maskNetwork - 8 - 8)
		octet := binaryNetwork[2]
		octet = octet[:network_bits] + strings.Repeat("0", 8-network_bits)
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
	return network
}
