//A program that takes a slice of integers and reverses the order of its elements.

package main
import (
	"fmt"
)

func main() {
	vals := []int{1,2,3,4,5,10,11,12,13,14,15}
	length := len(vals)-1
	for i:=0;i<length/2;i++ {
		vals[i],vals[length-i] = vals[length-i], vals[i]

	}
	fmt.Println(vals)
}
