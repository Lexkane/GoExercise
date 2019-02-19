package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printer(n int32) {
	output := 1
	for z := int32(1); z < n; z++ {
		for y := int32(1); y <= z; y++ {
			for x := int32(1); x <= y; x++ {
				if (x*x+y*y == z*z) {
					fmt.Println(output, z, x, y)
					output++
					if output > int(n) {
						goto Exit
					}
				}
			}
		}
	}
Exit:
}

//print all square ints
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter positive integer number till which you wish to print Pifagorian numbers")
	text, _ := reader.ReadString(' ')
	n,_ := strconv.Atoi(text)
	//fmt.Println(n)
	printer(int32(n))
}
