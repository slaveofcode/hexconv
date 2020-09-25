package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverseStr(str string) string {
	length := len(str)
	var reversedStr string
	for charIdx := 0; charIdx < length; charIdx++ {
		reversedStr = string(str[charIdx]) + reversedStr
	}
	return reversedStr
}

func binaryToDecimal(bits string) float64 {
	var base float64 = 2
	var decimal float64
	// 00000101 ~> 10100000 for easy loop
	reversedBits := reverseStr(bits)
	for x := 0; x < len(reversedBits); x++ {
		// only count binary true
		if reversedBits[x] == '1' {
			decimal += math.Pow(base, float64(x))
		}
	}
	return decimal
}

func decimalToHex(decimal float64) string {
	src := decimal
	var originSrc float64
	var res string

	hex := map[int]byte{}
	hex[10] = 'A'
	hex[11] = 'B'
	hex[12] = 'C'
	hex[13] = 'D'
	hex[14] = 'E'
	hex[15] = 'F'

	for src > 0 {
		originSrc = src
		src = math.Floor(src / 16)
		if src > 0 {
			remainder := int(originSrc) % 16

			got := false
			for idx, item := range hex {
				if idx == remainder {
					got = true
					res += string(item)
				}
			}

			if !got {
				res += strconv.Itoa(remainder)
			}
		} else {
			got := false
			for idx, item := range hex {
				if idx == int(originSrc) {
					got = true
					res += string(item)
				}
			}

			if !got {
				res += strconv.Itoa(int(originSrc))
			}
		}

	}

	return reverseStr(res)
}

func hexToDecimal(hex string) float64 {
	hexMap := map[byte]int{}
	hexMap['A'] = 10
	hexMap['B'] = 11
	hexMap['C'] = 12
	hexMap['D'] = 13
	hexMap['E'] = 14
	hexMap['F'] = 15

	totalDigits := len(hex)
	revStr := reverseStr(hex)
	var decimal float64
	for x := 0; x < totalDigits; x++ {
		found := false
		var replaceMentValue int
		for idx, item := range hexMap {
			if idx == revStr[x] {
				found = true
				replaceMentValue = item
			}
		}

		if !found {
			res, _ := strconv.Atoi(string(revStr[x]))
			decimal += float64(res) * math.Pow(16, float64(x))
		} else {
			decimal += float64(replaceMentValue) * math.Pow(16, float64(x))
		}
	}

	return decimal
}

func decimalToBinary(decimal float64) string {
	src := decimal
	var originSrc float64
	var res string

	for src > 0 {
		originSrc = src
		src = math.Floor(src / 2)
		if src > 0 {
			remainder := int(originSrc) % 2
			res += strconv.Itoa(remainder)
		} else {
			res += strconv.Itoa(int(originSrc))
		}
	}

	// auto add padding after, for 1 byte = 8 bits
	for x := len(res); x < 8; x++ {
		res += "0"
	}

	return reverseStr(res)
}

func main() {
	// bin := 0B00000101
	// oneByte := []byte("F")
	// fmt.Println(bin)
	// fmt.Printf("%08b %2x\n", oneByte, oneByte)
	// fmt.Println(binaryToDecimal())
	fmt.Println("111111111100110011000000 to Decimal:", binaryToDecimal("111111111100110011000000"))
	fmt.Println("2345 to Hexadecimal:", decimalToHex(2345))
	fmt.Println("278 to binary:", decimalToBinary(278))
	fmt.Printf("ASCII: %08b \n", []byte("FFCCC0"))
	fmt.Println("FFCCC0 to binary: ", decimalToBinary(hexToDecimal("FFCCC0")))

	fmt.Print(hexToDecimal("FF"), hexToDecimal("F0"), hexToDecimal("0F"), hexToDecimal("11"))
	fmt.Print(" = ")
	fmt.Print(decimalToBinary(hexToDecimal("FF")), " ", decimalToBinary(hexToDecimal("F0")), " ", decimalToBinary(hexToDecimal("0F")), " ", decimalToBinary(hexToDecimal("11")), "\n")

	fmt.Printf("Bitwise AND %08b \n", 0b10000001&0b10001001)
	fmt.Printf("Bitwise OR %08b \n", 0b10000001|0b10001001)
	fmt.Printf("Bitwise XOR %08b \n", 0b10000001^0b10001001)
	fmt.Printf("Left shift %08b \n", 0b10000001>>2)
	fmt.Printf("RIght shift %08b \n", 0b10000001<<2)

	// see https://golang.org/pkg/fmt/ for more formatting output
	fmt.Printf("source %08b %d \n", 0b10000001, 0b10000001)
	fmt.Printf("shift left by 2 %10b %d \n", 0b10000001<<2, 0b10000001<<2)
}
