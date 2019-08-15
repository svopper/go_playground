package main

import "fmt"

func main() {
	arr := []byte{0x00, 0x80}
	test1 := (arr[0] << 8) | arr[1]
	test2 := (0x00 << 8) | 0x80

	fmt.Printf("%x\n", test1)
	fmt.Printf("%x\n", test2)
}
