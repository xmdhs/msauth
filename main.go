package main

import (
	"fmt"

	"github.com/xmdhs/msauth/auth"
)

func main() {
	code, err := auth.Getcode()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code)
}
