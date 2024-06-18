package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	fmt.Println("Do not delete this file before adding any new main .go files!!!")
	test, _ := runtime.Bool("true")
	fmt.Println(test)
}
