package main

import (
	"fmt"
	"strconv"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	data := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{""}
	for i, _ := range digits {
		index, _ := strconv.Atoi(string(digits[i]))
		str := data[index]
		temp := []string{}
		for s := range res {
			for j, _ := range str {
				temp = append(temp, res[s]+string(str[j]))
			}
		}

		res = temp
	}

	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
