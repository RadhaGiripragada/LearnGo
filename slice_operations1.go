//A program to declare a slice. Append 12,345,678 to slice. 
//Display subset and modify slice [5] to 100
package main
import (
	"fmt"
)

func main() {
	intslice := []int{1,2,3,4,5,6}
	fmt.Println(intslice)

	intslice = append(intslice, 12,345,678)
	fmt.Println(intslice)

	intslice[5] = 100
	fmt.Println(intslice)
}
