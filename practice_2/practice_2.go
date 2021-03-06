package practice_2

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	// set reader and writer
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	// 1. NとQを入力値として受け取る
	var N, Q int
	fmt.Fscan(r, &N)
	fmt.Fscan(r, &Q)

	// 2. Nの数だけアルファベットを順に文字列としてつなげる
	// 3. N個のうち２つを組み合わせて出力
	// 4. <もしくは>を入力値として受け取る
	// 5. <であれば２つを並び替える
	// 6. 軽い順に並ぶまで3から5を繰り返す
	// 7. 並べた文字列を出力する
	fmt.Fprintln(w, "! ans")
	// 8. プログラムを終了する
}
