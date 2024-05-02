// a program that takes a slice of strings and checks if it forms a palindrome (reads the same forwards and backwards).
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var str string
	str = "madam kayak radar test thesewords"
	str = strings.ReplaceAll(str," ","")
	st, _ := strconv.Atoi(os.Args[1])
	en, _ := strconv.Atoi(os.Args[2])
	
	poli := str[st:en+1]
	fmt.Println("pil",poli)
	leng := len(poli)
	for i := 0; i < leng/2; i++ {
		if poli[i] != poli[leng-i-1] {
			fmt.Println("Not a polindrom")
			return
		}
	}
	fmt.Println("Polindrom")
	return

}
