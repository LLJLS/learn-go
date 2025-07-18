package main

import "fmt"

func testReturn(b string) string {
	a := b
	return a
}

func main() {
	a := testReturn("abc")
	fmt.Println(a)

}
