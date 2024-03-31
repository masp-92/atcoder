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

	var N, M, L, Q int
	N = readInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readInt()
	}
	M = readInt()
	B := make([]int, M)
	for i := 0; i < M; i++ {
		B[i] = readInt()
	}
	L = readInt()
	C := make([]int, L)
	for i := 0; i < L; i++ {
		C[i] = readInt()
	}
	Q = readInt()
	X := make([]int, Q)
	for i := 0; i < Q; i++ {
		X[i] = readInt()
	}

	convMap := map[int]bool{}
	for _, a := range A {
		for _, b := range B {
			for _, c := range C {
				convMap[a+b+c] = true
			}
		}
	}
	for _, x := range X {
		if convMap[x] {
			fmt.Fprintln(wt, "Yes")
		} else {
			fmt.Fprintln(wt, "No")
		}
	}
}
