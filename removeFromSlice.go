// search for a value in slice and remove all the references of the value
package main

import (
	"fmt"
)

func main() {
	testslice := []int{1,2,3,4,5,7,6,67,7,7,9,8}
	remove:=7

	length := len(testslice)
	var cnt int

	for i,val := range testslice {
		if val == remove {
			testslice[i] = testslice[length-cnt-1]
			cnt++	
		}
	}
	if cnt==0 {
		cnt=1
	}
	testslice = testslice[:length-cnt]
	fmt.Println(testslice)
}
