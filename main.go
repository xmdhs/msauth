package main

import (
	"encoding/json"
	"fmt"

	"github.com/xmdhs/msauth/auth"
)

func main() {
	code, err := auth.Getcode("")
	ok := true
	var errmsg string
	if err != nil {
		ok = false
		errmsg = err.Error()
	}
	i := info{
		Ok:   ok,
		Err:  errmsg,
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
