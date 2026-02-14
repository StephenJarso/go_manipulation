package main

import (
	"fmt"
	"manipulation/exercises"
)
type Key struct{
	Row, Col int
}
func main(){
m:=[]int{2,3,4,5}
v:=[]int{4,9,16,25}
	
fmt.Println(exercises.Comp(m,v))
}