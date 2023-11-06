package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2023,11,6,9,49,10,10, time.UTC)
	fmt.Println(utc)

	layout := "2004-10-06"
	parse, _ := time.Parse(layout, "2003-05-10")
	fmt.Println(parse)
}