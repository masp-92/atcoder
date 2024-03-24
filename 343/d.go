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

func main() {
	defer wt.Flush()          // プログラムの最後でバッファの内容をフラッシュする
	sc.Split(bufio.ScanWords) // 単語単位で読み込むように設定

	n := readInt()
	t := readInt()

	scores := make([]int, n)
	scoreMap := map[int]int{0: n}
	variation := 1
	for i := 0; i < t; i++ {
		a := readInt()
		b := readInt()
		a--
		scoreMap[scores[a]]--
		if scoreMap[scores[a]] == 0 {
			variation--
		}
		scores[a] += b
		scoreMap[scores[a]]++
		if scoreMap[scores[a]] == 1 {
			variation++
		}
		fmt.Fprintln(wt, variation)
	}

}
