package main

import (
	"fmt"
	"github.com/tinthid/simple-orchestration-engine/server"
)

func main() {
	s := server.CreateServer("1", "2")
	fmt.Println(s)
)
