package main

import (
	"fmt"
)
func main(){
	m:=make(map[string]int)
	//insert/update
	m["apple"] = 0
	//
v,exist := m["apple"]
	
fmt.Println(v, exist)
}