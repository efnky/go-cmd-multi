package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("go-cmd-multi worker started (no HTTP port)")
	for {
		fmt.Println("worker tick")
		time.Sleep(30 * time.Second)
	}
}
