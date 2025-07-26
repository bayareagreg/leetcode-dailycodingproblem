package main

import "fmt"

func convert(s string, numRows int) string {
	// i-th char belongs to which line
	m := make(map[int]int)
	line, delta := 0, 1
	for i := range s {
		if numRows == 1 {
			m[i] = 0
		} else {
			m[i] = line
			line += delta
			if line == 0 || line == numRows-1 {
				delta = -delta
			}
		}
	}

	/*
		fmt.Printf("%v\n", m)
		for i := range s {
			fmt.Printf("%c: %v, ", s[i], m[i])
		}
		fmt.Println()
	*/

	res := ""
	for k := range numRows {
		for i := range s {
			if m[i] == k {
				res += string(s[i])
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("%v\n", convert("PAYPALISHIRING", 4))

}
