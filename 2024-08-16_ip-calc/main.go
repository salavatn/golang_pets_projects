package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "time"
)

var net_host = "192.168.100.155"
var net_mask = 21
var myBinaryIP []string

func main() {
	myBinaryIP := getBinaryIP(net_host) // [11111111 11111111 11111111 11111010]
	validatorNetwork(myBinaryIP, net_mask)

	fmt.Printf("Address: 	%s\n", net_host)
	fmt.Printf("Network: 	%s\n", getNetwork(myBinaryIP, net_mask))
	fmt.Printf("HostMin:	%s\n", getHostMin(myBinaryIP, net_mask))
	fmt.Printf("HostMax:	%s\n", getHostMax(myBinaryIP, net_mask))
	fmt.Printf("Broadcast:	%s\n", getBroadcast(myBinaryIP, net_mask))
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
	var networkAddress []string

	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)

	if netmask < 8 {
		octet := binaryNetwork[1]
		octet = octet[:netmask] + strings.Repeat("0", 8-netmask)
		copy(binaryNetwork[0:], []string{octet, "00000000", "00000000", "00000000"})
	} else if netmask >= 8 && netmask < 16 {
		networkOctet := -8 + netmask
		octet := binaryNetwork[1]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		copy(binaryNetwork[1:], []string{octet, "00000000", "00000000"})
	} else if netmask >= 16 && netmask < 24 {
		networkOctet := -16 + netmask
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		copy(binaryNetwork[2:], []string{octet, "00000000"})
	} else if netmask >= 24 {
		networkOctet := -24 + netmask
		octet := binaryNetwork[3]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		copy(binaryNetwork[3:], []string{octet})
	} // binaryNetwork = [00001010 11111111 11100000 00000000]

	for _, octet := range binaryNetwork {
		octetIP, _ := strconv.ParseInt(octet, 2, 64)
		networkAddress = append(networkAddress, fmt.Sprintf("%d", octetIP))
	} // networkAddress = [10 255 224 0]

	network := strings.Join(networkAddress, ".")
	mask := strconv.Itoa(netmask)

	return network + "/" + mask // 10.255.224.0/19
}

func getHostMin(binaryNet []string, netmask int) string {
	var networkAddress []string

	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)

	if netmask < 8 {
		octet := binaryNetwork[1]
		octet = octet[:netmask] + strings.Repeat("0", 8-netmask)
		copy(binaryNetwork[0:], []string{octet, "00000000", "00000000", "00000001"})
	} else if netmask >= 8 && netmask < 16 {
		networkOctet := -8 + netmask
		octet := binaryNetwork[1]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		copy(binaryNetwork[1:], []string{octet, "00000000", "00000001"})
	} else if netmask >= 16 && netmask < 24 {
		networkOctet := -16 + netmask
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet)
		copy(binaryNetwork[2:], []string{octet, "00000001"})
	} else if netmask >= 24 {
		networkOctet := -24 + netmask
		octet := binaryNetwork[3]
		octet = octet[:networkOctet] + strings.Repeat("0", 8-networkOctet-1) + "1"
		copy(binaryNetwork[3:], []string{octet})
	} // binaryNetwork = [00001010 11111111 11100000 00000000]

	for _, octet := range binaryNetwork {
		octetIP, _ := strconv.ParseInt(octet, 2, 64)
		networkAddress = append(networkAddress, fmt.Sprintf("%d", octetIP))
	} // networkAddress = [10 255 224 0]

	minHost := strings.Join(networkAddress, ".")

	return minHost // 10.255.255.241
}

func getHostMax(binaryNet []string, netmask int) string {
	var networkAddress []string

	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)

	if netmask < 8 {
		octet := binaryNetwork[1]
		octet = octet[:netmask] + strings.Repeat("1", 8-netmask)
		copy(binaryNetwork[0:], []string{octet, "11111111", "11111111", "11111110"})
	} else if netmask >= 8 && netmask < 16 {
		networkOctet := -8 + netmask
		octet := binaryNetwork[1]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		copy(binaryNetwork[1:], []string{octet, "11111111", "11111110"})
	} else if netmask >= 16 && netmask < 24 {
		networkOctet := -16 + netmask
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		copy(binaryNetwork[2:], []string{octet, "11111110"})
	} else if netmask >= 24 {
		networkOctet := -24 + netmask
		octet := binaryNetwork[3]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet-1) + "0"
		copy(binaryNetwork[3:], []string{octet})
	} // binaryNetwork = [00001010 11111111 11100000 00000000]

	for _, octet := range binaryNetwork {
		octetIP, _ := strconv.ParseInt(octet, 2, 64)
		networkAddress = append(networkAddress, fmt.Sprintf("%d", octetIP))
	} // networkAddress = [10 255 224 0]

	maxHost := strings.Join(networkAddress, ".")

	return maxHost // 10.255.255.241
}

func getBroadcast(binaryNet []string, netmask int) string {
	var networkAddress []string

	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)

	if netmask < 8 {
		octet := binaryNetwork[1]
		octet = octet[:netmask] + strings.Repeat("1", 8-netmask)
		copy(binaryNetwork[0:], []string{octet, "11111111", "11111111", "11111111"})
	} else if netmask >= 8 && netmask < 16 {
		networkOctet := -8 + netmask
		octet := binaryNetwork[1]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		copy(binaryNetwork[1:], []string{octet, "11111111", "11111111"})
	} else if netmask >= 16 && netmask < 24 {
		networkOctet := -16 + netmask
		octet := binaryNetwork[2]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		copy(binaryNetwork[2:], []string{octet, "11111111"})
	} else if netmask >= 24 {
		networkOctet := -24 + netmask
		octet := binaryNetwork[3]
		octet = octet[:networkOctet] + strings.Repeat("1", 8-networkOctet)
		copy(binaryNetwork[3:], []string{octet})
	}

	for _, octet := range binaryNetwork {
		octetIP, _ := strconv.ParseInt(octet, 2, 64)
		networkAddress = append(networkAddress, fmt.Sprintf("%d", octetIP))
	} // networkAddress = [10 255 224 0]

	maxHost := strings.Join(networkAddress, ".")

	return maxHost // 10.255.255.255
}

// In-Progress
// Будет получать IP адрес и маску сети на проверку
func validatorNetwork(binaryNet []string, netmask int) string {
	if netmask < 1 || netmask > 32 {
		fmt.Printf("Error: Вы ввели %d! Установить маску сети [1-32].\n", netmask)
		os.Exit(0)
	}

	binaryNetwork := make([]string, 4)
	copy(binaryNetwork, binaryNet)

	counter := 1

	for _, octet := range binaryNetwork {
		octetIP, _ := strconv.ParseInt(octet, 2, 64)
		if counter == 1 && octetIP == 0 {
			fmt.Printf("Error: Не верный адрес сети/хоста.\n")
			os.Exit(0)
		} else if octetIP < 0 || octetIP > 255 {
			fmt.Printf("Error: Указали не верный адрес сети/хоста.\n")
			os.Exit(0)
		}
		counter++

	}

	return ""
}

// For green GitHub Timeline! :D
