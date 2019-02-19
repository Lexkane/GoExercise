package main

import (
	"fmt"
	"math"
)

/*func ipow(base int,  exp int) int64 {
	result := 1;
	for {
	if (exp & 1)
	result *= base;
	exp = exp>>1;
	if (!exp)
	break;
	base *= base;
	}

	return result;
}
*/
func main() {
	 s := float64(10000)
	 z:=float64(1)
	 for (math.Pow(2, z) <=s) {
			z++
		}
	 fmt.Println(z+1)

}
