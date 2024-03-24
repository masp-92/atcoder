package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := range a {
		tmp := make([]int, n)
		for j := range tmp {
			fmt.Scan(&tmp[j])
		}
		a[i] = tmp
	}

	for i := range a {
		for j := range a[i] {
			if a[i][j] == 1 {
				fmt.Print(j+1, " ")
			}
		}
		fmt.Println()
	}

}
