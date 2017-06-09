package main

import (
	"fmt"

	"github.com/Devatoria/go-nsenter"
)

func main() {
	config := &nsenter.Config{
		Target: 1,
		Mount:  true,
	}

	stdout, stderr, err := config.Execute("ls", "-la")
	fmt.Println(stdout)
	fmt.Println(stderr)
	if err != nil {
		panic(err)
	}
}
