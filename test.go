package main

import (
	"fmt"
	"time"
)

func main() {

	year := time.Now().Format("20060102")

	fmt.Println(year)

}
