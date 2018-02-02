package main
import "fmt"

func main() {
	a := 1
	b := "hello, world"
	c := make(chan int)

	go func(c chan int) {
		fmt.Println("In goroutine")
		t := "xiayi"
		fmt.Println("name is %s", t)
		c <- 1
	}(c)

	fmt.Println("num is :%d", a)
	fmt.Println("string is %s", b)

	sig := <- c
	fmt.Println("sig is :%d", sig)
}
