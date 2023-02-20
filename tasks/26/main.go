package main

import (
	"fmt"
	"strings"
)

func IsUniqueString(s string) bool {
	s = strings.ToLower(s) // перевеодим всю строку в нижний регистр
	chars := make(map[int32]bool) // делаем множество для проверки входит в него символ или нет
	for _, c := range s { // делаем проверку на соответствие 
		if chars[c] {
			return false
		}
		chars[c] = true
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(IsUniqueString(str1)) 
	fmt.Println(IsUniqueString(str2)) 
	fmt.Println(IsUniqueString(str3)) 
}
