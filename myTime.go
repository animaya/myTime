package myTime

import "fmt"

func WhatTimeIsItNow() {
	fmt.Println(time.Now().Format(time.RFC3339))
}
