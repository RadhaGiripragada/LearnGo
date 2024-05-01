//A program that takes a slice of integers and removes all occurrences of a specific value from the slice
package main 
import (
	"fmt"
)

func main() {
	vals1 := []int{12,34,45,56,67,78}
	vals2 :=[]int{11,33,44,55,66} 

	vals1 = append (vals1,vals2...)
	fmt.Println("vals ",vals1)
}
