package myTime

import "fmt"

//latest upgrade

func WhatTimeIsItNow() {
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Pintln("\n")
}
