package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	cubeRoot := int(math.Cbrt(float64(n)))
	for i := cubeRoot; i > 0; i-- {
		cube := i * i * i
		cubeS := strconv.Itoa(cube)
		cycle := true
		for j := 0; j < len(cubeS)/2; j++ {
			if cubeS[j] != cubeS[len(cubeS)-j-1] {
				cycle = false
				break
			}
		}
		if cycle {
			fmt.Println(cube)
			break
		}
	}
}
