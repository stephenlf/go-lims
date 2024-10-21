package main

import (
	"github.com/kr/pretty"
	"github.com/stephenlf/go-lims/db"
)

func main() {
	state := db.Init()
	pretty.Print(state, nil)
}
