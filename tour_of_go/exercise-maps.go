package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var myarr []string
	
	myarr = strings.Fields(s)
	
	for _,v := range(myarr) {
		m[v]++
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}