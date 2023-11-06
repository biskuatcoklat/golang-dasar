package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

type userSlice []user

func (value userSlice) Len() int {
	return len(value)
}

func (value userSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value userSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users:= []user {
		{"wahyu", 20},
		{"zaki", 6},
	}
	sort.Sort(userSlice(users))
	fmt.Println(users)
}
