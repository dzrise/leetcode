//12. Integer to Roman
//Medium
//Topics
//Companies
//Seven different symbols represent Roman numerals with the following values:
//
//Symbol	Value
//I	1
//V	5
//X	10
//L	50
//C	100
//D	500
//M	1000
//Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:
//
//If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
//If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
//Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
//Given an integer, convert it to a Roman numeral.
//
//
//
//Example 1:
//
//Input: num = 3749
//
//Output: "MMMDCCXLIX"
//
//Explanation:
//
//3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
//700 = DCC as 500 (D) + 100 (C) + 100 (C)
//40 = XL as 10 (X) less of 50 (L)
//9 = IX as 1 (I) less of 10 (X)
//Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
//Example 2:
//
//Input: num = 58
//
//Output: "LVIII"
//
//Explanation:
//
//50 = L
//8 = VIII
//Example 3:
//
//Input: num = 1994
//
//Output: "MCMXCIV"
//
//Explanation:
//
//1000 = M
//900 = CM
//90 = XC
//4 = IV
//
//
//Constraints:

package main

import "fmt"

func intToRoman(num int) string {
	data := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
	res := ""

	if num/1000 > 0 {
		for i := 0; i < num/1000; i++ {
			res += data[1000]
		}
		num = num % 1000
	}

	if num/900 > 0 {
		res += data[100] + data[1000]
		num = num % 900
	}

	if num/500 > 0 {
		for i := 0; i < num/500; i++ {
			res += data[500]
			num = num % 500
		}
	}
	if num/400 > 0 {
		res += data[100] + data[500]
		num = num % 400
	}

	if num/100 > 0 {
		for i := 0; i < num/100; i++ {
			res += data[100]
		}
		num = num % 100
	}

	if num/90 > 0 {
		res += data[10] + data[100]
		num = num % 90
	}

	if num/50 > 0 {
		for i := 0; i < num/50; i++ {
			res += data[50]
		}
		num = num % 50
	}

	if num/40 > 0 {
		res += data[10] + data[50]
		num = num % 40

	}

	if num/10 > 0 {
		for i := 0; i < num/10; i++ {
			res += data[10]
		}
		num = num % 10
	}

	if num/9 > 0 {
		res += data[1] + data[10]
		num = num % 9
	}

	if num/8 > 0 {
		res += data[5] + data[1] + data[1] + data[1]
		num = num % 8
	}

	if num/8 > 0 {
		res += data[5] + data[1] + data[1]
		num = num % 8
	}

	if num/8 > 0 {
		res += data[5] + data[1]
		num = num % 8
	}

	if num/5 > 0 {
		for i := 0; i < num/5; i++ {
			res += data[5]
		}
		num = num % 5
	}

	if num/4 > 0 {
		res += data[1] + data[5]
		num = num % 4
	}

	if num/1 > 0 {
		for i := 0; i < num/1; i++ {
			res += data[1]
		}
		num = num % 1
	}

	return res

}

func main() {
	fmt.Println(intToRoman(3749))
	//fmt.Println(intToRoman(58))
	//fmt.Println(intToRoman(1994))
	//fmt.Println(intToRoman(1000))
	//fmt.Println(intToRoman(500))
	//fmt.Println(intToRoman(100))
	//fmt.Println(intToRoman(50))
	//fmt.Println(intToRoman(10))
	//fmt.Println(intToRoman(5))
	//fmt.Println(intToRoman(1))
	//fmt.Println(intToRoman(4))
	//fmt.Println(intToRoman(9))
	//fmt.Println(intToRoman(40))
}
