//count vowels in string
package main
import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	var cnt int 
	for _,val := range input {
		if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' {
			cnt++	
		}
	}
	fmt.Println("\n The no of vowels are ",cnt)

}
