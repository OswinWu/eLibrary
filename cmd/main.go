package cmd

import (
	"fmt"
)

func Main() {
	fmt.Println("Hello, World!")
}

type Config struct {
	Host string
	Port int
}
