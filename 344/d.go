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

// readString は次の整数を読み込むヘルパー関数です。
func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wt.Flush()          // プログラムの最後でバッファの内容をフラッシュする
	sc.Split(bufio.ScanWords) // 単語単位で読み込むように設定
	T := readString()
	N := readInt()
	candidate := map[string]int{"": 0}
	for i := 0; i < N; i++ {
		ai := readInt()
		for j := 0; j < ai; j++ {
			s := readString()
			for c, score := range candidate {
				newC := c + s
				if len(T) >= len(newC) && T[:len(newC)] == newC {
					newScore := score + 1
					if _, ok := candidate[newC]; !ok {
						candidate[newC] = newScore
					} else if candidate[newC] > newScore {
						candidate[newC] = newScore
					}
				}
			}
		}
	}
	ans := -1
	for c, score := range candidate {
		if c == T {
			if ans == -1 || score < ans {
				ans = score
			}
		}
	}
	fmt.Fprintln(wt, ans)
}
