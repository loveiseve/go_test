package main
import "fmt"

func main() {
	a := 1
	b := "hello, world"

	go func() {
		fmt.Println("In goroutine")
		t := "xiayi"
		fmt.Println("name is %s", t)
	}()

	fmt.Println("num is :%d", a)
	fmt.Println("string is %s", b)
}
