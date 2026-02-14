package exercises


 func Comp(array1 []int,array2 []int)bool{
	if len(array1) != len(array2){
		return false
	}
	freq:=make(map[int]int,len(array1))
	for _,v:=range array1{
		square := v*v
		freq[square]++
	}
	for _,v:=range array2{
		if freq[v] == 0 {
			return false
		}
		freq[v]--
	}
	return true

 }