package main

import "fmt"

func square(n int, sqchan chan int) {
	defer close(sqchan)
	sqchan <- n * n
}
func Cube(n int, cubechan chan int) {
	defer close(cubechan)
	cubechan <- n * n * n
}
func main() {

	number := 5
	sqchan := make(chan int)
	cubechan := make(chan int)
	go square(number, sqchan)
	go Cube(number, cubechan)

	square := <-sqchan
	cube := <-cubechan

	fmt.Printf("%d  %d\n", square, cube)

}
