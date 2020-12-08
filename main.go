package main

import (
	"encoding/json"
	"fmt"

	"github.com/xmdhs/msauth/auth"
)

func main() {
	code, err := auth.Getcode()
	ok := true
	if err != nil {
		ok = false
	}
	i := info{
		Ok:   ok,
		Err:  err.Error(),
		Code: code,
	}
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

type info struct {
	Ok   bool   `json:"ok"`
	Err  string `json:"err"`
	Code string `json:"code"`
}
