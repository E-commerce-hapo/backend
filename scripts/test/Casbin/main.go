package main

import (
	"github.com/k0kubun/pp"
	"github.com/kiem-toan/infrastructure/authorize/auth"
)

func main() {
	e := auth.New()
	res := e.Check(auth.Roles{"shipper"}, "action:post")
	pp.Println(res)

}
