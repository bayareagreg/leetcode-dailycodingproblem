package main

import (
	"fmt"
)

type item struct {
	n     int
	bytes []byte
}

func decodeString(str string) string {
	num := 0
	st := []item{{1, []byte{}}}

	for i := range str {
		switch {
		case str[i] == '0':
			num *= 10
		case str[i] > '0' && str[i] <= '9':
			num = num*10 + int(str[i]-'0')
		case str[i] == '[':
			st = append(st, item{num, []byte{}})
			num = 0
		case str[i] == ']':
			tmp := st[len(st)-1]
			st = st[:len(st)-1]
			for range tmp.n {
				st[len(st)-1].bytes = append(st[len(st)-1].bytes, tmp.bytes...)
			}
		default:
			st[len(st)-1].bytes = append(st[len(st)-1].bytes, str[i])
		}
	}

	return string(st[0].bytes)
}

func main() {
	//fmt.Printf("%v\n", decodeString("3[a]2[bc]"))
	//fmt.Printf("%v\n", decodeString("2[abc]3[cd]ef"))
	//fmt.Printf("%v\n", decodeString("3[a2[c]]"))
	//fmt.Printf("%v\n", decodeString("a2[c]"))
	fmt.Printf("%v\n", decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))

	/*
		"3[a]2[bc]"
		"3[a2[c]]"
		"2[abc]3[cd]ef"
		"c3[2[a]1[b]]d"
		"3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
		"2[3[a2[c]]d]e"
		"ef3[a2[c]]xy"
		"11[2[a]]"
	*/

}
