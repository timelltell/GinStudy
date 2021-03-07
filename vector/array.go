package vector

import (
	"fmt"
)
var a [3]int = [3]int{1,2,3}


func Run(){
	fmt.Println(a)
	for i,v := range a{
		fmt.Printf("i %d   v   %d  ",i,v)
	}
	var b [3]int=[3]int{1,2}
	c:=[...]int{1,2,3}
	fmt.Println(b==c)
	fmt.Printf("%T  ",c)
	var array [3][2]int
	array=[3][2]int{{1,1},{2,2},{3,3}}
	fmt.Println(array)
	array=[3][2]int{1:{1,1},2:{1:3}}
	fmt.Println(array)
}