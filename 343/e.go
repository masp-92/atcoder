package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

// readInt は次の整数を読み込むヘルパー関数です。
func readInt() int {
	sc.Scan()
	val, _ := strconv.Atoi(sc.Text())
	return val
}

func max(v ...int) int {
	max := v[0]
	for _, vi := range v {
		if vi > max {
			max = vi
		}
	}
	return max
}

func min(v ...int) int {
	min := v[0]
	for _, vi := range v {
		if vi < min {
			min = vi
		}
	}
	return min
}

func calcArea(c1 []int, c2 []int, c3 []int) (int, int, int) {
	x1, y1, z1 := c1[0], c1[1], c1[2]
	x2, y2, z2 := c2[0], c2[1], c2[2]
	x3, y3, z3 := c3[0], c3[1], c3[2]

	// v3を求める
	xw := min(x1+7, x2+7, x3+7) - max(x1, x2, x3)
	yw := min(y1+7, y2+7, y3+7) - max(y1, y2, y3)
	zw := min(z1+7, z2+7, z3+7) - max(z1, z2, z3)
	v3 := xw * yw * zw

	// v2を求める
	// c1,c2
	xw12 := min(x1+7, x2+7) - max(x1, x2)
	yw12 := min(y1+7, y2+7) - max(y1, y2)
	zw12 := min(z1+7, z2+7) - max(z1, z2)
	//c2,c3
	xw23 := min(x2+7, x3+7) - max(x2, x3)
	yw23 := min(y2+7, y3+7) - max(y2, y3)
	zw23 := min(z2+7, z3+7) - max(z2, z3)
	//c3,c1
	xw31 := min(x3+7, x1+7) - max(x3, x1)
	yw31 := min(y3+7, y1+7) - max(y3, y1)
	zw31 := min(z3+7, z1+7) - max(z3, z1)

	v2 := (xw12 * yw12 * zw12) + (xw23 * yw23 * zw23) + (xw31 * yw31 * zw31) - (3 * v3)

	v1 := 3*7*7*7 - 2*v2 - 3*v3
	return v1, v2, v3
}

func main() {
	defer wt.Flush()          // プログラムの最後でバッファの内容をフラッシュする
	sc.Split(bufio.ScanWords) // 単語単位で読み込むように設定

	V1, V2, V3 := readInt(), readInt(), readInt()

	c1 := []int{0, 0, 0}
	c2 := []int{0, 0, 0}
	c3 := []int{0, 0, 0}

	for c2i := 0; c2i <= 7; c2i++ {
		for c2j := 0; c2j <= 7; c2j++ {
			for c2k := -7; c2k <= 7; c2k++ {
				for c3i := 0; c3i <= 7; c3i++ {
					for c3j := 0; c3j <= 7; c3j++ {
						for c3k := -7; c3k <= 7; c3k++ {
							c2[0], c2[1], c2[2] = c2i, c2j, c2k
							c3[0], c3[1], c3[2] = c3i, c3j, c3k
							v1, v2, v3 := calcArea(c1, c2, c3)
							if v1 == V1 && v2 == V2 && v3 == V3 {
								fmt.Fprintln(wt, "Yes")
								fmt.Fprintln(wt, c1[0], c1[1], c1[2], c2[0], c2[1], c2[2], c3[0], c3[1], c3[2])
								return
							}
						}
					}
				}
			}
		}
	}
	fmt.Fprintln(wt, "No")
}
