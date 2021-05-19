package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &num)

	for i := 1; i < num; i++ {
		fmt.Fprintln(w, num)
		num--
	}
}
