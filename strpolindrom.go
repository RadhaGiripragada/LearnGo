// check if the string is a polindrom or not
package main

import (
	"fmt"
	"os"
)

func main() {
	var str1 string = os.Args[1]
	//var str2 string = "testAthis"
	leng := len(str1)
	for i := 0; i < leng/2; i++ {
		if str1[i] != str1[leng-i-1] {
			fmt.Println("Not a polidrom")
			return
		}
	}
	fmt.Println("A polidrom")
}
