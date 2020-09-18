package bufioPractice

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, c int
	var s string

	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b, &c)
	fmt.Fscan(r, &s)

	fmt.Fprintln(w, "sum is " + string(a+b+c))
	fmt.Fprintln(w, "your name is " + s)

}