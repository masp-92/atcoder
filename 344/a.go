package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	S := readString()

	sDiv := strings.Split(S, "|")
	fmt.Fprintf(wt, "%s%s\n", sDiv[0], sDiv[2])
}
