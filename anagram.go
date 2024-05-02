// check if 2  strings are anagrams or not(contain the same characters in a different order) using a for range loop.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func anagramtest(str1 []string, str2 []string) bool {
	sort.Strings(str1)
	sort.Strings(str2)
	if (len(str1)!=len(str2)) {
		return false
	}
	for i:=0;i<len(str1);i++ {
		if str1[i]!=str2[i] {
			return false
		}
	}
	return true
}

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]

	
	sstr1 := strings.Split(strings.ToLower(strings.ReplaceAll(str1," ","")),"")
	sstr2 := strings.Split(strings.ToLower(strings.ReplaceAll(str2," ","")),"")
	fmt.Println(anagramtest(sstr1, sstr2))
}
