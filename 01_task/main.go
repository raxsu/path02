package main

import (
	"fmt"
	"strconv"
)

func main() {
	str, nubmer := "104", 35
	numberA := strconv.Itoa(nubmer)

	if str1, err := strconv.Atoi(str); err == nil {
		fmt.Printf("str - %T, %v", str1, str1)
	}

	fmt.Printf(" number - %T, %v\n", numberA, numberA)

}
