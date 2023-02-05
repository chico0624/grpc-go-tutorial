package main

import "fmt"

func main() {
	res := hello("ryota")
	fmt.Println(res)
}

func hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
