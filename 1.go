package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	formattedTime := currentTime.Format("02-01-2006 15:04:05")

	fmt.Println("Date and time now:", formattedTime)
}
