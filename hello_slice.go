//A basic program to declare and print a slice using for range

package main
import (
	"fmt"
)

func main() {
	intvals := []int{1,2,3,4,5,6}
	strvals := [ ]string{"test","this","string","values"}

	for i,val := range intvals {
		fmt.Println(i,"--",val)
	}
	for i,val := range strvals {
		fmt.Println(i,"--",val)
	}
}
