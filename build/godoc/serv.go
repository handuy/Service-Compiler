package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Complier Go is ready !!!")
	interval := time.Tick(1 * time.Hour)
	for c := range interval {
		fmt.Println(c.Date())
	}
}
