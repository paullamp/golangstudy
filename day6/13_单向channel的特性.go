package main

// "fmt"

func main() {
	var ch chan int
	ch <- 666
	// 双向channel 可以转换成单向
	var writech chan<- int = ch
	writech <- 888

	var readch <-chan int = ch
	<-readch
	//单向channel 无法转换成双向channel
	var ch2 chan int = writech //(type chan<- int) as type chan
}
