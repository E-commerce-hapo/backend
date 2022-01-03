package main

import (
	"github.com/E-commerce-hapo/backend/pkg/authorize/auth"
	"github.com/k0kubun/pp"
)

func main() {
	e := auth.New()
	res := e.Check(auth.Roles{"shipper"}, "action:post")
	pp.Println(res)

}
