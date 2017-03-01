package main

//计算斐波那契数列第N位的数字
import (
	"time"
	"fmt"
	"flag"
)

var n = flag.Int("n", 0, "斐波那契数列位置")

func main() {
	flag.Parse()
	go spinner(100 * time.Microsecond)
	fibN := fib(*n)
	fmt.Printf("\rFibonacci(%d) = %d\n", *n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-|\/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}