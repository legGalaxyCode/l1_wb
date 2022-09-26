package main

import (
	"fmt"
	"strconv"
	"strings"
)

// первый способ изменить бит через побитовые операции
func replaceBit(num int64, idx int, val bool) int64 {
	if val {
		num = num | (1 << idx) // маска
		return num
	} else {
		num = num & ^(1 << idx)
		return num
	}
}

// второй способ через преобразование в строку бит и обратно
func replaceBitString(num int64, idx int, val bool) int64 {
	bits := strings.Split(fmt.Sprintf("%064b", num), "")
	if val {
		bits[len(bits)-idx-1] = "1"
	} else {
		bits[len(bits)-idx-1] = "0"
	}
	var binaryNum string
	for _, bit := range bits {
		binaryNum += bit
	}
	res, err := strconv.ParseInt(binaryNum, 2, 64)
	if err != nil {
		fmt.Printf("Error converting string %s\n", binaryNum)
		return 0
	}
	return res
}

// в обоих случаях биты считаются справа налево
func main() {
	test1True := replaceBit(12, 1, true)
	test1False := replaceBit(12, 1, false)
	test2True := replaceBit(8934230, 5, true)
	test2False := replaceBit(8934230, 5, false)
	fmt.Printf("Before:\n12: %064b,  %d\n", 12, 12)
	fmt.Printf("After setting 1 bit in 1:\n12: %064b,  %d\n", test1True, test1True)
	fmt.Printf("After setting 1 bit in 0:\n12: %064b,  %d\n", test1False, test1False)
	fmt.Printf("Before:\n%064b,  %d\n", 8934230, 8934230)
	fmt.Printf("After setting 5 bit in 1:\n%064b,  %d\n", test2True, test2True)
	fmt.Printf("After setting 5 bit in 0:\n%064b,  %d\n", test2False, test2False)
	fmt.Println("-----------------------------------------------------------")
	test1TrueS := replaceBitString(12, 1, true)
	test1FalseS := replaceBitString(12, 1, false)
	test2TrueS := replaceBitString(8934230, 5, true)
	test2FalseS := replaceBitString(8934230, 5, false)
	fmt.Printf("Before:\n12: %064b,  %d\n", 12, 12)
	fmt.Printf("After setting 1 bit in 1:\n12: %064b,  %d\n", test1TrueS, test1TrueS)
	fmt.Printf("After setting 1 bit in 0:\n12: %064b,  %d\n", test1FalseS, test1FalseS)
	fmt.Printf("Before:\n%064b,  %d\n", 8934230, 8934230)
	fmt.Printf("After setting 5 bit in 1:\n%064b,  %d\n", test2TrueS, test2TrueS)
	fmt.Printf("After setting 5 bit in 0:\n%064b,  %d\n", test2FalseS, test2FalseS)
}
