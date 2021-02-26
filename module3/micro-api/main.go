package main

import (
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/micro/v2/cmd"
)

func main() {
	cmd.Init()
}
